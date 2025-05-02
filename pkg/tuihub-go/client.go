package tuihub

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/tuihub/librarian/pkg/tuihub-go/internal"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"

	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/metadata"
)

type LibrarianClient struct {
	pb.LibrarianSephirahServiceClient

	accessToken       string
	refreshToken      string
	muToken           sync.RWMutex
	backgroundRefresh bool
	consulConfig      *capi.Config
}

type ClientOption func(*LibrarianClient)

func WithoutBackgroundRefresh() ClientOption {
	return func(c *LibrarianClient) {
		c.backgroundRefresh = false
	}
}

func WithClientConsulConfig(config *capi.Config) ClientOption {
	return func(c *LibrarianClient) {
		c.consulConfig = config
	}
}

func LoginByPassword(
	ctx context.Context,
	username string,
	password string,
	options ...ClientOption,
) (*LibrarianClient, error) {
	c := &LibrarianClient{
		LibrarianSephirahServiceClient: nil,
		accessToken:                    "",
		refreshToken:                   "",
		muToken:                        sync.RWMutex{},
		backgroundRefresh:              true,
		consulConfig:                   nil,
	}
	for _, o := range options {
		o(c)
	}
	client, err := internal.NewSephirahClient(ctx, c.consulConfig, os.Getenv(sephirahServiceName))
	if err != nil {
		return nil, err
	}
	c.LibrarianSephirahServiceClient = client
	resp, err := client.GetToken(ctx, &pb.GetTokenRequest{
		Username: username,
		Password: password,
		DeviceId: nil,
	})
	if err != nil {
		return nil, err
	}
	c.accessToken = resp.GetAccessToken()
	c.refreshToken = resp.GetRefreshToken()

	if c.backgroundRefresh {
		go c.RunBackgroundRefresh()
	}
	return c, nil
}

func LoginByRefreshToken(
	ctx context.Context,
	refreshToken string,
	options ...ClientOption,
) (*LibrarianClient, error) {
	c := &LibrarianClient{
		LibrarianSephirahServiceClient: nil,
		accessToken:                    "",
		refreshToken:                   "",
		muToken:                        sync.RWMutex{},
		backgroundRefresh:              true,
		consulConfig:                   nil,
	}
	for _, o := range options {
		o(c)
	}
	client, err := internal.NewSephirahClient(ctx, c.consulConfig, os.Getenv(sephirahServiceName))
	if err != nil {
		return nil, err
	}
	c.LibrarianSephirahServiceClient = client
	resp, err := client.RefreshToken(
		WithToken(ctx, refreshToken),
		new(pb.RefreshTokenRequest),
	)
	if err != nil {
		return nil, err
	}
	c.accessToken = resp.GetAccessToken()
	c.refreshToken = resp.GetRefreshToken()

	if c.backgroundRefresh {
		go c.RunBackgroundRefresh()
	}
	return c, nil
}

func WithToken(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
}

func (c *LibrarianClient) RunBackgroundRefresh() {
	for {
		c.muToken.RLock()
		resp, err := c.RefreshToken(
			WithToken(context.Background(), c.refreshToken),
			new(pb.RefreshTokenRequest),
		)
		if err == nil {
			return
		}
		c.muToken.RUnlock()

		c.muToken.Lock()
		c.accessToken = resp.GetAccessToken()
		c.refreshToken = resp.GetRefreshToken()
		c.muToken.Unlock()

		time.Sleep(time.Hour)
	}
}

func (c *LibrarianClient) WithToken(ctx context.Context) context.Context {
	c.muToken.RLock()
	defer c.muToken.RUnlock()
	return WithToken(ctx, c.accessToken)
}

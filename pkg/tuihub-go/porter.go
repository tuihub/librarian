package tuihub

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/tuihub/librarian/pkg/tuihub-go/internal"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	serviceID           = "PORTER_SERVICE_ID"
	serverNetwork       = "SERVER_NETWORK"
	serverAddr          = "SERVER_ADDRESS"
	serverTimeout       = "SERVER_TIMEOUT"
	sephirahServiceName = "SEPHIRAH_SERVICE_NAME"
)

type Porter struct {
	server            *grpc.Server
	requireAsUser     bool
	wrapper           *serviceWrapper
	logger            log.Logger
	app               *kratos.App
	consulConfig      *capi.Config
	serverConfig      *ServerConfig
	backgroundServers []transport.Server
}

type ServerConfig struct {
	Network string
	Addr    string
	Timeout *time.Duration
}

type PorterOption func(*Porter)

func WithLogger(logger log.Logger) PorterOption {
	return func(p *Porter) {
		p.logger = logger
	}
}

func WithPorterConsulConfig(config *capi.Config) PorterOption {
	return func(p *Porter) {
		p.consulConfig = config
	}
}

func WithAsUser() PorterOption {
	return func(p *Porter) {
		p.requireAsUser = true
	}
}

func WithBackgroundServer(server transport.Server) PorterOption {
	return func(p *Porter) {
		if server == nil {
			return
		}
		p.backgroundServers = append(p.backgroundServers, server)
	}
}

func NewPorter(
	ctx context.Context,
	info *pb.GetPorterInformationResponse,
	service pb.LibrarianPorterServiceServer,
	options ...PorterOption,
) (*Porter, error) {
	porterName := "unknown"
	if info == nil {
		return nil, fmt.Errorf("[tuihub-go][%s] info is nil", porterName)
	}
	if info.GetBinarySummary() != nil && len(info.GetBinarySummary().GetName()) > 0 {
		porterName = info.GetBinarySummary().GetName()
	}
	if service == nil {
		return nil, fmt.Errorf("[tuihub-go][%s]serviceServer is nil", porterName)
	}
	if info.GetBinarySummary() == nil {
		return nil, fmt.Errorf("[tuihub-go][%s]binary summary is nil", porterName)
	}
	if info.GetGlobalName() == "" {
		return nil, fmt.Errorf("[tuihub-go][%s]global name is empty", porterName)
	}
	if info.GetFeatureSummary() == nil {
		return nil, fmt.Errorf("[tuihub-go][%s]feature summary is nil", porterName)
	}
	p := new(Porter)
	p.logger = log.DefaultLogger
	for _, o := range options {
		o(p)
	}
	if p.serverConfig == nil {
		p.serverConfig = defaultServerConfig()
	}
	client, err := internal.NewPorterClient(ctx, p.consulConfig, os.Getenv(sephirahServiceName))
	if err != nil {
		return nil, fmt.Errorf("[tuihub-go][%s]failed to create consul client: %w", porterName, err)
	}
	r, err := internal.NewRegistrar(p.consulConfig)
	if err != nil {
		return nil, fmt.Errorf("[tuihub-go][%s]failed to create registrar: %w", porterName, err)
	}
	c := &serviceWrapper{
		LibrarianPorterServiceServer: service,
		Info:                         info,
		Logger:                       p.logger,
		Client:                       client,
		RequireToken:                 p.requireAsUser,
		Token:                        nil,
		tokenMu:                      sync.Mutex{},
		lastHeartbeat:                time.Time{},
		lastRefreshToken:             time.Time{},
	}
	p.wrapper = c
	p.server = NewServer(
		p.serverConfig,
		NewService(c),
		p.logger,
	)
	id, _ := os.Hostname()
	name := "porter"
	id = fmt.Sprintf("%s-%s-%s", id, name, info.GetBinarySummary().GetName())
	if customID, exist := os.LookupEnv(serviceID); exist {
		id = fmt.Sprintf("%s-%s", id, customID)
	}
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(name),
		kratos.Version(p.wrapper.Info.GetBinarySummary().GetBuildVersion()),
		kratos.Metadata(map[string]string{
			"PorterName": p.wrapper.Info.GetGlobalName(),
		}),
		kratos.Server(append([]transport.Server{p.server}, p.backgroundServers...)...),
		kratos.Registrar(r),
	)
	p.app = app
	return p, nil
}

func (p *Porter) SetSephirahPorterClient(client pb.LibrarianSephirahPorterServiceClient) {
	if p.wrapper == nil {
		p.wrapper = new(serviceWrapper)
	}
	p.wrapper.Client = client
}

func (p *Porter) Run() error {
	return p.app.Run()
}

func (p *Porter) Stop() error {
	return p.app.Stop()
}

func defaultServerConfig() *ServerConfig {
	config := ServerConfig{
		Network: "",
		Addr:    "",
		Timeout: nil,
	}
	if network, exist := os.LookupEnv(serverNetwork); exist {
		config.Network = network
	}
	if addr, exist := os.LookupEnv(serverAddr); exist {
		config.Addr = addr
	}
	if timeout, exist := os.LookupEnv(serverTimeout); exist {
		d, err := time.ParseDuration(timeout)
		if err == nil {
			config.Timeout = &d
		}
	}
	return &config
}

func WellKnownToString(e protoreflect.Enum) string {
	return fmt.Sprint(proto.GetExtension(
		e.
			Descriptor().
			Values().
			ByNumber(
				e.
					Number(),
			).
			Options(),
		librarian.E_ToString,
	))
}

type PorterClient struct {
	pb.LibrarianSephirahPorterServiceClient

	accessToken string
}

func (c *PorterClient) WithToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+c.accessToken)
}

func (p *Porter) GetPorterService() pb.LibrarianPorterServiceServer {
	return p.wrapper
}

func (p *Porter) GetBackgroundServices() []transport.Server {
	return p.backgroundServers
}

func (p *Porter) ReverseCall(ctx context.Context) (*PorterClient, error) {
	if !p.requireAsUser {
		return nil, errors.New("init porter with `WithAsUser` option to use this method")
	}
	if p.wrapper.Token == nil {
		return nil, errors.New("porter not enabled")
	}
	client, err := internal.NewPorterClient(ctx, p.consulConfig, os.Getenv(sephirahServiceName))
	if err != nil {
		return nil, err
	}
	return &PorterClient{
		LibrarianSephirahPorterServiceClient: client,
		accessToken:                          p.wrapper.Token.AccessToken,
	}, nil
}

func (p *Porter) AsUser(ctx context.Context, userID int64) (*LibrarianClient, error) {
	if !p.requireAsUser {
		return nil, errors.New("init porter with `WithAsUser` option to use this method")
	}
	if p.wrapper.Token == nil {
		return nil, errors.New("porter not enabled")
	}
	client, err := internal.NewPorterClient(ctx, p.consulConfig, os.Getenv(sephirahServiceName))
	if err != nil {
		return nil, err
	}
	resp, err := client.AcquireUserToken(
		WithToken(ctx, p.wrapper.Token.AccessToken),
		&pb.AcquireUserTokenRequest{
			UserId: &librarian.InternalID{Id: userID},
		},
	)
	if err != nil {
		return nil, err
	}
	client2, err := internal.NewSephirahClient(ctx, p.consulConfig, os.Getenv(sephirahServiceName))
	if err != nil {
		return nil, err
	}
	return &LibrarianClient{
		LibrarianSephirahServiceClient: client2,
		accessToken:                    resp.GetAccessToken(),
		refreshToken:                   "",
		muToken:                        sync.RWMutex{},
		backgroundRefresh:              false,
		consulConfig:                   p.consulConfig,
	}, nil
}

package client

import (
	"github.com/tuihub/librarian/internal/lib/libkratos"
	"github.com/tuihub/librarian/pkg/tuihub-go"
	tuihubbangumi "github.com/tuihub/librarian/pkg/tuihub-bangumi"
	tuihubrss "github.com/tuihub/librarian/pkg/tuihub-rss"
	tuihubsteam "github.com/tuihub/librarian/pkg/tuihub-steam"
	tuihubtelegram "github.com/tuihub/librarian/pkg/tuihub-telegram"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"

	"github.com/fullstorydev/grpchan/inprocgrpc"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/samber/lo"
)

type InprocPorter struct {
	porters   []*tuihub.Porter
	Instances map[string]porter.LibrarianPorterServiceServer
	Servers   []transport.Server
}

func NewInprocPorter() (*InprocPorter, error) {
	rss, err := tuihubrss.NewPorter("")
	if err != nil {
		return nil, err
	}
	steam, err := tuihubsteam.NewPorter("")
	if err != nil {
		return nil, err
	}
	telegram, err := tuihubtelegram.NewPorter("")
	if err != nil {
		return nil, err
	}
	bangumi, err := tuihubbangumi.NewPorter("")
	if err != nil {
		return nil, err
	}
	return &InprocPorter{
		porters: []*tuihub.Porter{
			rss, steam, telegram, bangumi,
		},
		Instances: map[string]porter.LibrarianPorterServiceServer{
			"inproc://tuihub-rss":      rss.GetPorterService(),
			"inproc://tuihub-steam":    steam.GetPorterService(),
			"inproc://tuihub-telegram": telegram.GetPorterService(),
			"inproc://tuihub-bangumi":  bangumi.GetPorterService(),
		},
		Servers: lo.Flatten([][]transport.Server{
			rss.GetBackgroundServices(),
			steam.GetBackgroundServices(),
			telegram.GetBackgroundServices(),
			bangumi.GetBackgroundServices(),
		}),
	}, nil
}

func (i *InprocPorter) SetSephirahServer(
	server pb.LibrarianSephirahPorterServiceServer,
	middlewares []middleware.Middleware,
) {
	channel := inprocgrpc.Channel{}
	channel.WithServerUnaryInterceptor(libkratos.MiddlewareToGRPCUnaryInterceptor(middlewares...))
	pb.RegisterLibrarianSephirahPorterServiceServer(&channel, server)
	client := pb.NewLibrarianSephirahPorterServiceClient(&channel)

	for _, p := range i.porters {
		p.SetSephirahPorterClient(client)
	}
}

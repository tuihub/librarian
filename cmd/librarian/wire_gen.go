// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	service2 "github.com/tuihub/librarian/app/miner/pkg/service"
	"github.com/tuihub/librarian/app/searcher/pkg/service"
	service3 "github.com/tuihub/librarian/app/sephirah/pkg/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(librarian_EnableServiceDiscovery *conf.Librarian_EnableServiceDiscovery, sephirahServer *conf.SephirahServer, database *conf.Database, s3 *conf.S3, porter *conf.Porter, mapper_Data *conf.Mapper_Data, searcher_Data *conf.Searcher_Data, miner_Data *conf.Miner_Data, auth *conf.Auth, mq *conf.MQ, cache *conf.Cache, consul *conf.Consul, settings *libapp.Settings) (*kratos.App, func(), error) {
	libauthAuth, err := libauth.NewAuth(auth)
	if err != nil {
		return nil, nil, err
	}
	builtInObserver, err := libobserve.NewBuiltInObserver()
	if err != nil {
		return nil, nil, err
	}
	libmqMQ, cleanup, err := libmq.NewMQ(mq, database, cache, settings, builtInObserver)
	if err != nil {
		return nil, nil, err
	}
	cron, err := libcron.NewCron()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	store, err := libcache.NewStore(cache)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianSearcherServiceServer, cleanup2, err := service.NewSearcherService(searcher_Data, settings)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianMinerServiceServer, cleanup3, err := service2.NewMinerService(miner_Data, settings)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	inprocClients := inprocgrpc.NewInprocClients(librarianSearcherServiceServer, librarianMinerServiceServer)
	librarianSearcherServiceClient, err := searcherClientSelector(librarian_EnableServiceDiscovery, consul, inprocClients, settings)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianMinerServiceClient, err := minerClientSelector(librarian_EnableServiceDiscovery, consul, inprocClients, settings)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianSephirahServiceServer, cleanup4, err := service3.NewSephirahService(sephirahServer, database, s3, porter, consul, libauthAuth, libmqMQ, cron, store, settings, librarianSearcherServiceClient, librarianMinerServiceClient)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcServer, err := server.NewGRPCServer(sephirahServer, libauthAuth, librarianSephirahServiceServer, settings, builtInObserver)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewGrpcWebServer(grpcServer, sephirahServer, libauthAuth, settings, builtInObserver)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, err := newApp(grpcServer, httpServer, libmqMQ, cron, builtInObserver, consul)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/mapper/pkg/service"
	service3 "github.com/tuihub/librarian/app/porter/pkg/service"
	service2 "github.com/tuihub/librarian/app/searcher/pkg/service"
	service4 "github.com/tuihub/librarian/app/sephirah/pkg/service"
	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(librarian_EnableServiceDiscovery *conf.Librarian_EnableServiceDiscovery, sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, mapper_Data *conf.Mapper_Data, searcher_Data *conf.Searcher_Data, porter_Data *conf.Porter_Data, auth *conf.Auth, mq *conf.MQ, settings *libapp.Settings) (*kratos.App, func(), error) {
	libauthAuth, err := libauth.NewAuth(auth)
	if err != nil {
		return nil, nil, err
	}
	libmqMQ, cleanup, err := libmq.NewMQ(mq, settings)
	if err != nil {
		return nil, nil, err
	}
	cron := libcron.NewCron()
	librarianMapperServiceServer, cleanup2, err := service.NewMapperService(mapper_Data, settings)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianSearcherServiceServer, cleanup3, err := service2.NewSearcherService(searcher_Data, settings)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianPorterServiceServer, cleanup4, err := service3.NewPorterService(porter_Data, settings)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	inprocClients := inprocgrpc.NewInprocClients(librarianMapperServiceServer, librarianSearcherServiceServer, librarianPorterServiceServer)
	discoverClients, err := client.NewDiscoverClients()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianMapperServiceClient := mapperClientSelector(librarian_EnableServiceDiscovery, inprocClients, discoverClients)
	librarianSearcherServiceClient := searcherClientSelector(librarian_EnableServiceDiscovery, inprocClients, discoverClients)
	librarianPorterServiceClient := porterClientSelector(librarian_EnableServiceDiscovery, inprocClients, discoverClients)
	librarianSephirahServiceServer, cleanup5, err := service4.NewSephirahService(sephirah_Data, libauthAuth, libmqMQ, cron, settings, librarianMapperServiceClient, librarianSearcherServiceClient, librarianPorterServiceClient)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcServer := server.NewGRPCServer(sephirah_Server, libauthAuth, librarianSephirahServiceServer, settings)
	httpServer := server.NewGrpcWebServer(grpcServer, sephirah_Server, libauthAuth, settings)
	app := newApp(grpcServer, httpServer, libmqMQ, cron)
	return app, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, auth *conf.Auth) (*kratos.App, func(), error) {
	libauthAuth, err := libauth.NewAuth(auth)
	if err != nil {
		return nil, nil, err
	}
	librarianMapperServiceClient, err := client.NewMapperClient()
	if err != nil {
		return nil, nil, err
	}
	librarianPorterServiceClient, err := client.NewPorterClient()
	if err != nil {
		return nil, nil, err
	}
	librarianSearcherServiceClient, err := client.NewSearcherClient()
	if err != nil {
		return nil, nil, err
	}
	angelaBase, err := bizangela.NewAngelaBase(librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	if err != nil {
		return nil, nil, err
	}
	entClient, cleanup, err := data.NewSQLClient(sephirah_Data)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(entClient)
	tipherethRepo := data.NewTipherethRepo(dataData)
	topicImpl := bizangela.NewPullSteamAccountAppRelationTopic(angelaBase)
	libmqTopicImpl := bizangela.NewPullAccountTopic(angelaBase, topicImpl)
	tipherethUseCase, err := biztiphereth.NewTipherethUseCase(tipherethRepo, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient, libmqTopicImpl)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	geburaRepo := data.NewGeburaRepo(dataData)
	callbackControlBlock := bizbinah.NewCallbackControl()
	geburaUseCase := bizgebura.NewGeburaUseCase(geburaRepo, libauthAuth, callbackControlBlock, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	binahUseCase := bizbinah.NewBinahUseCase(callbackControlBlock, libauthAuth, librarianMapperServiceClient, librarianPorterServiceClient, librarianSearcherServiceClient)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angelaBase, tipherethUseCase, geburaUseCase, binahUseCase)
	grpcServer := server.NewGRPCServer(sephirah_Server, libauthAuth, librarianSephirahServiceServer)
	app := newApp(grpcServer)
	return app, func() {
		cleanup()
	}, nil
}

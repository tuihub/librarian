// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/data"
	"github.com/tuihub/librarian/app/porter/internal/server"
	"github.com/tuihub/librarian/app/porter/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(porter_Server *conf.Porter_Server, porter_Data *conf.Porter_Data, settings *libapp.Settings) (*kratos.App, func(), error) {
	s3Repo, err := data.NewS3Repo(porter_Data)
	if err != nil {
		return nil, nil, err
	}
	s3 := bizs3.NewS3(s3Repo)
	librarianPorterServiceServer := service.NewLibrarianPorterServiceService(s3)
	grpcServer := server.NewGRPCServer(porter_Server, librarianPorterServiceServer, settings)
	registrar, err := libapp.NewRegistrar()
	if err != nil {
		return nil, nil, err
	}
	mainMetadata := newMetadata(s3)
	app := newApp(grpcServer, registrar, mainMetadata)
	return app, func() {
	}, nil
}

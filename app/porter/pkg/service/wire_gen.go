// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	"github.com/tuihub/librarian/app/porter/internal/client"
	"github.com/tuihub/librarian/app/porter/internal/client/feed"
	"github.com/tuihub/librarian/app/porter/internal/client/steam"
	"github.com/tuihub/librarian/app/porter/internal/data"
	"github.com/tuihub/librarian/app/porter/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/protos/pkg/librarian/porter/v1"
)

// Injectors from wire.go:

func NewPorterService(porter_Data *conf.Porter_Data, settings *libapp.Settings) (v1.LibrarianPorterServiceServer, func(), error) {
	collector := client.NewColly()
	rssRepo, err := feed.NewRSSRepo(collector)
	if err != nil {
		return nil, nil, err
	}
	feedUseCase := bizfeed.NewFeed(rssRepo)
	steamSteam, err := steam.NewSteam(collector, porter_Data)
	if err != nil {
		return nil, nil, err
	}
	steamUseCase := bizsteam.NewSteamUseCase(steamSteam)
	s3Repo, err := data.NewS3Repo(porter_Data)
	if err != nil {
		return nil, nil, err
	}
	s3 := bizs3.NewS3(s3Repo)
	librarianPorterServiceServer := service.NewLibrarianPorterServiceService(feedUseCase, steamUseCase, s3)
	return librarianPorterServiceServer, func() {
	}, nil
}

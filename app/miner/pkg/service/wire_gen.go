// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/tuihub/librarian/app/miner/internal/biz"
	"github.com/tuihub/librarian/app/miner/internal/data"
	"github.com/tuihub/librarian/app/miner/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/protos/pkg/librarian/miner/v1"
)

// Injectors from wire.go:

func NewMinerService(miner_Data *conf.Miner_Data, settings *libapp.Settings) (v1.LibrarianMinerServiceServer, func(), error) {
	minerRepo := data.NewMinerRepo(miner_Data)
	miner := biz.NewMiner(minerRepo)
	librarianMinerServiceServer := service.NewLibrarianMinerServiceService(miner)
	return librarianMinerServiceServer, func() {
	}, nil
}
// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	client2 "github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
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
func wireApp(sephirahServer *conf.SephirahServer, database *conf.Database, s3 *conf.S3, porter *conf.Porter, auth *conf.Auth, mq *conf.MQ, cache *conf.Cache, consul *conf.Consul, settings *libapp.Settings) (*kratos.App, func(), error) {
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
	entClient, cleanup2, err := data.NewSQLClient(database, settings)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dataData := data.NewData(entClient)
	angelaRepo := data.NewAngelaRepo(dataData)
	librarianPorterServiceClient, err := client.NewPorterClient(consul)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientPorter, err := client.NewPorter(librarianPorterServiceClient, consul)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	netzachRepo := data.NewNetzachRepo(dataData)
	librarianSearcherServiceClient, err := client2.NewSearcherClient(consul)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	searcher := client.NewSearcher(librarianSearcherServiceClient)
	topic := biznetzach.NewSystemNotificationTopic(netzachRepo, searcher)
	supervisorSupervisor, err := supervisor.NewSupervisor(porter, libauthAuth, clientPorter, topic)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	geburaRepo := data.NewGeburaRepo(dataData)
	angelaBase, err := bizangela.NewAngelaBase(angelaRepo, supervisorSupervisor, geburaRepo, librarianPorterServiceClient, searcher)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	store, err := libcache.NewStore(cache)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	libcacheMap := bizangela.NewAppInfoCache(geburaRepo, store)
	libmqTopic := bizangela.NewUpdateAppInfoIndexTopic(angelaBase)
	topic2 := bizangela.NewPullAppInfoTopic(angelaBase, libcacheMap, libmqTopic)
	topic3 := bizangela.NewPullAccountAppInfoRelationTopic(angelaBase, topic2)
	topic4 := bizangela.NewPullAccountTopic(angelaBase, topic3)
	map2 := bizangela.NewNotifyFlowCache(netzachRepo, store)
	map3 := bizangela.NewFeedToNotifyFlowCache(netzachRepo, store)
	map4 := bizangela.NewNotifyTargetCache(netzachRepo, store)
	topic5 := bizangela.NewNotifyPushTopic(angelaBase, map4)
	topic6 := bizangela.NewNotifyRouterTopic(angelaBase, map2, map3, topic5)
	topic7 := bizangela.NewFeedItemPostprocessTopic(angelaBase, topic6, topic)
	topic8 := bizangela.NewPullFeedTopic(angelaBase, topic7, topic)
	angela, err := bizangela.NewAngela(libmqMQ, topic4, topic3, topic2, topic8, topic6, topic5, topic7, libmqTopic)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tipherethRepo := data.NewTipherethRepo(dataData)
	cron, err := libcron.NewCron()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	key := biztiphereth.NewUserCountCache(tipherethRepo, store)
	tiphereth, err := biztiphereth.NewTiphereth(settings, tipherethRepo, libauthAuth, supervisorSupervisor, searcher, topic4, cron, key)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, libauthAuth, searcher, librarianPorterServiceClient, supervisorSupervisor, libmqTopic, topic2, libcacheMap)
	binahRepo, err := data.NewBinahRepo(s3)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	controlBlock := bizbinah.NewControlBlock(libauthAuth)
	binah := bizbinah.NewBinah(binahRepo, controlBlock, libauthAuth, librarianSearcherServiceClient)
	yesodRepo := data.NewYesodRepo(dataData)
	map5 := bizyesod.NewFeedOwnerCache(yesodRepo, store)
	yesod, err := bizyesod.NewYesod(yesodRepo, supervisorSupervisor, cron, searcher, topic8, topic, map5)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	netzach, err := biznetzach.NewNetzach(netzachRepo, supervisorSupervisor, searcher, libmqMQ, map3, map2, map4, topic)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	chesedRepo := data.NewChesedRepo(dataData)
	librarianMinerServiceClient, err := client2.NewMinerClient(consul)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	map6 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, binahRepo, cron, librarianPorterServiceClient, searcher, librarianMinerServiceClient, controlBlock, map6)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v := server.NewAuthMiddleware(libauthAuth)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(angela, tiphereth, gebura, binah, yesod, netzach, chesed, supervisorSupervisor, settings, libauthAuth, v, sephirahServer)
	grpcServer, err := server.NewGRPCServer(sephirahServer, libauthAuth, librarianSephirahServiceServer, settings, builtInObserver)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewGrpcWebServer(grpcServer, sephirahServer, libauthAuth, settings, builtInObserver)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	registrar, err := libapp.NewRegistrar(consul)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := newApp(grpcServer, httpServer, libmqMQ, cron, registrar)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}

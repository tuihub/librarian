// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/miner/pkg/service"
	"github.com/tuihub/librarian/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/biz/bizchesed"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/bizkether"
	"github.com/tuihub/librarian/internal/biz/biznetzach"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/biz/bizyesod"
	"github.com/tuihub/librarian/internal/client/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/server"
	"github.com/tuihub/librarian/internal/service/angelaweb"
	"github.com/tuihub/librarian/internal/service/sephirah"
	"github.com/tuihub/librarian/internal/service/supervisor"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(librarian_EnableServiceDiscovery *conf.Librarian_EnableServiceDiscovery, sephirahServer *conf.SephirahServer, database *conf.Database, s3 *conf.S3, porter *conf.Porter, miner_Data *conf.Miner_Data, auth *conf.Auth, mq *conf.MQ, cache *conf.Cache, consul *conf.Consul, search *conf.Search, settings *libapp.Settings) (*kratos.App, func(), error) {
	libauthAuth, err := libauth.NewAuth(auth)
	if err != nil {
		return nil, nil, err
	}
	entClient, cleanup, err := data.NewSQLClient(database, settings)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(entClient)
	ketherRepo := data.NewKetherRepo(dataData)
	builtInObserver, err := libobserve.NewBuiltInObserver()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	libmqMQ, cleanup2, err := libmq.NewMQ(mq, database, cache, settings, builtInObserver)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianPorterServiceClient, err := client.NewPorterClient(consul, porter, settings)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientPorter, err := client.NewPorter(librarianPorterServiceClient, consul, porter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	netzachRepo := data.NewNetzachRepo(dataData)
	idGenerator := libidgenerator.NewIDGenerator()
	topic := biznetzach.NewSystemNotificationTopic(netzachRepo, idGenerator)
	tipherethRepo := data.NewTipherethRepo(dataData)
	store, err := libcache.NewStore(cache)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	libcacheMap := biztiphereth.NewPorterInstanceCache(tipherethRepo, store)
	map2 := biztiphereth.NewPorterContextCache(tipherethRepo, store)
	supervisorSupervisor, err := supervisor.NewSupervisor(porter, libmqMQ, libauthAuth, clientPorter, topic, libcacheMap, map2)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	geburaRepo := data.NewGeburaRepo(dataData)
	libsearchSearch, err := libsearch.NewSearch(search, settings)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	ketherBase, err := bizkether.NewKetherBase(ketherRepo, supervisorSupervisor, geburaRepo, librarianPorterServiceClient, libsearchSearch, idGenerator)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	map3 := bizkether.NewAppInfoCache(geburaRepo, store)
	libmqTopic := bizkether.NewUpdateAppInfoIndexTopic(ketherBase)
	topic2 := bizkether.NewPullAppInfoTopic(ketherBase, map3, libmqTopic)
	topic3 := bizkether.NewPullAccountAppInfoRelationTopic(ketherBase, topic2)
	topic4 := bizkether.NewPullAccountTopic(ketherBase, topic3)
	map4 := bizkether.NewNotifyFlowCache(netzachRepo, store)
	map5 := bizkether.NewFeedToNotifyFlowCache(netzachRepo, store)
	map6 := bizkether.NewNotifyTargetCache(netzachRepo, store)
	topic5 := bizkether.NewNotifyPushTopic(ketherBase, map6)
	topic6 := bizkether.NewNotifyRouterTopic(ketherBase, map4, map5, topic5)
	topic7 := bizkether.NewFeedItemPostprocessTopic(ketherBase, topic6, topic)
	topic8 := bizkether.NewPullFeedTopic(ketherBase, topic7, topic)
	kether, err := bizkether.NewKether(ketherBase, libmqMQ, topic4, topic3, topic2, topic8, topic6, topic5, topic7, libmqTopic)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cron, err := libcron.NewCron()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	key := biztiphereth.NewUserCountCache(tipherethRepo, store)
	tiphereth, err := biztiphereth.NewTiphereth(settings, tipherethRepo, libauthAuth, supervisorSupervisor, idGenerator, libsearchSearch, topic4, cron, key, libcacheMap)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gebura := bizgebura.NewGebura(geburaRepo, libauthAuth, idGenerator, libsearchSearch, librarianPorterServiceClient, supervisorSupervisor, libmqTopic, topic2, map3)
	binahRepo, err := data.NewBinahRepo(s3)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	controlBlock := bizbinah.NewControlBlock(libauthAuth)
	binah := bizbinah.NewBinah(binahRepo, controlBlock, libauthAuth)
	yesodRepo := data.NewYesodRepo(dataData)
	map7 := bizyesod.NewFeedOwnerCache(yesodRepo, store)
	yesod, err := bizyesod.NewYesod(yesodRepo, supervisorSupervisor, cron, idGenerator, libsearchSearch, topic8, topic, map7)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	netzach, err := biznetzach.NewNetzach(netzachRepo, supervisorSupervisor, idGenerator, libsearchSearch, libmqMQ, map5, map4, map6, topic)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	chesedRepo := data.NewChesedRepo(dataData)
	librarianMinerServiceServer, cleanup3, err := service.NewMinerService(miner_Data, settings)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	inprocClients := inprocgrpc.NewInprocClients(librarianMinerServiceServer)
	librarianMinerServiceClient, err := minerClientSelector(librarian_EnableServiceDiscovery, consul, inprocClients, settings)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	map8 := bizchesed.NewImageCache(store)
	chesed, err := bizchesed.NewChesed(chesedRepo, binahRepo, idGenerator, libsearchSearch, cron, librarianPorterServiceClient, librarianMinerServiceClient, controlBlock, map8)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	librarianSephirahServiceServer := sephirah.NewLibrarianSephirahServiceService(kether, tiphereth, gebura, binah, yesod, netzach, chesed, supervisorSupervisor, settings, libauthAuth, sephirahServer)
	grpcServer, err := server.NewGRPCServer(sephirahServer, libauthAuth, librarianSephirahServiceServer, settings, builtInObserver)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewGrpcWebServer(grpcServer, sephirahServer, libauthAuth, settings, builtInObserver)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	angelaWeb := angelaweb.NewAngelaWeb(libauthAuth, tiphereth, key)
	app, err := newApp(grpcServer, httpServer, angelaWeb, libmqMQ, cron, builtInObserver, consul)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

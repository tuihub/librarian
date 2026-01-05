package models

import (
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"
)

func GetModels() []interface{} {
	return []interface{}{
		&model.User{},
		&model.Account{},
		&model.Session{},
		&model.Device{},
		&model.Tag{},
		&model.KV{},

		&modelgebura.App{},
		&modelgebura.AppCategory{},
		&modelgebura.AppAppCategory{},
		&modelgebura.AppInfo{},
		&modelgebura.AppRunTime{},
		&modelgebura.Sentinel{},
		&modelgebura.SentinelAppBinary{},
		&modelgebura.SentinelAppBinaryFile{},
		&modelgebura.SentinelLibrary{},
		&modelgebura.SentinelSession{},
		&modelgebura.StoreApp{},
		&modelgebura.StoreAppBinary{},

		&modelyesod.FeedConfig{},
		&modelyesod.FeedActionSet{},
		&modelyesod.FeedConfigAction{},
		&modelyesod.FeedItemCollection{},
		&modelyesod.FeedItemCollectionFeedItem{},

		&modelfeed.Feed{},
		&modelfeed.Item{},

		&modelnetzach.NotifyFlow{},
		&modelnetzach.NotifyFlowSource{},
		&modelnetzach.NotifyFlowTarget{},
		&modelnetzach.NotifySource{},
		&modelnetzach.NotifyTarget{},
		&modelnetzach.SystemNotification{},

		&modelsupervisor.PorterInstance{},
		&modelsupervisor.PorterContext{},

		&modelchesed.Image{},
		&modelchesed.File{},
	}
}

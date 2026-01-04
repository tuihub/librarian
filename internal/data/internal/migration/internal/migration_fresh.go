package internal

import (
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"

	"gorm.io/gorm"
)

func MigrationFresh(db *gorm.DB) error {
	return db.AutoMigrate([]interface{}{
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

		&modelfeed.Feed{},
		&modelfeed.Item{},

		&modelyesod.FeedConfig{},
		&modelyesod.FeedActionSet{},
		&modelyesod.FeedConfigAction{},
		&modelyesod.FeedItemCollection{},
		&modelyesod.FeedItemCollectionFeedItem{},

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
	}...)
}

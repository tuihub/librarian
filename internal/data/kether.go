package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelyesod"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type KetherRepo struct {
	data *Data
}

// NewKetherRepo .
func NewKetherRepo(data *Data) *KetherRepo {
	return &KetherRepo{
		data: data,
	}
}

func (k *KetherRepo) UpsertAccount(ctx context.Context, acc model.Account) error {
	account := gormschema.Account{
		ID:                acc.ID,
		Platform:          acc.Platform,
		PlatformAccountID: acc.PlatformAccountID,
		Name:              acc.Name,
		ProfileURL:        acc.ProfileURL,
		AvatarURL:         acc.AvatarURL,
	}
	return k.data.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "platform"}, {Name: "platform_account_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "profile_url", "avatar_url", "updated_at"}),
	}).Create(&account).Error
}

func (k *KetherRepo) UpsertAppInfo(
	ctx context.Context, ap *modelgebura.AppInfo, internal *modelgebura.AppInfo,
) error {
	return k.data.WithTx(ctx, func(tx *gorm.DB) error {
		appInfo := gormschema.AppInfo{
			ID:                 ap.ID,
			Source:             ap.Source,
			SourceAppID:        ap.SourceAppID,
			SourceURL:          ap.SourceURL,
			Name:               ap.Name,
			Type:               gormschema.ToSchemaAppType(ap.Type),
			ShortDescription:   ap.ShortDescription,
			Description:        ap.Description,
			IconImageURL:       ap.IconImageURL,
			BackgroundImageURL: ap.BackgroundImageURL,
			CoverImageURL:      ap.CoverImageURL,
			ReleaseDate:        ap.ReleaseDate,
			Developer:          ap.Developer,
			Publisher:          ap.Publisher,
		}

		// Check if exists
		var count int64
		if err := tx.Model(&gormschema.AppInfo{}).
			Where("source = ? AND source_app_id = ?", ap.Source, ap.SourceAppID).
			Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			// Create internal app info first
			internalInfo := gormschema.AppInfo{
				ID:          internal.ID,
				Source:      internal.Source,
				SourceAppID: internal.SourceAppID,
				Name:        internal.Name,
				Type:        gormschema.ToSchemaAppType(internal.Type),
			}
			if err := tx.Create(&internalInfo).Error; err != nil {
				return err
			}
		}

		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "source"}, {Name: "source_app_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"source_url", "name", "type", "short_description", "description", "icon_image_url", "background_image_url", "cover_image_url", "release_date", "developer", "publisher", "updated_at"}),
		}).Create(&appInfo).Error
	})
}

func (k *KetherRepo) UpsertAppInfos(ctx context.Context, al []*modelgebura.AppInfo) error {
	apps := make([]gormschema.AppInfo, len(al))
	for i, ap := range al {
		apps[i] = gormschema.AppInfo{
			ID:                 ap.ID,
			Source:             ap.Source,
			SourceAppID:        ap.SourceAppID,
			SourceURL:          ap.SourceURL,
			Name:               ap.Name,
			Type:               gormschema.ToSchemaAppType(ap.Type),
			ShortDescription:   ap.ShortDescription,
			Description:        ap.Description,
			IconImageURL:       ap.IconImageURL,
			BackgroundImageURL: ap.BackgroundImageURL,
			CoverImageURL:      ap.CoverImageURL,
			ReleaseDate:        ap.ReleaseDate,
			Developer:          ap.Developer,
			Publisher:          ap.Publisher,
		}
	}
	return k.data.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "source"}, {Name: "source_app_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"source_url", "name", "type", "short_description", "description", "icon_image_url", "background_image_url", "cover_image_url", "release_date", "developer", "publisher", "updated_at"}),
	}).Create(&apps).Error
}

func (k *KetherRepo) UpsertFeed(ctx context.Context, f *modelfeed.Feed) error {
	return k.data.WithTx(ctx, func(tx *gorm.DB) error {
		// Check config exists
		var config gormschema.FeedConfig
		if err := tx.Where("id = ?", f.ID).First(&config).Error; err != nil {
			return err
		}

		feed := gormschema.Feed{
			ID:          f.ID,
			Title:       f.Title,
			Description: f.Description,
			Link:        f.Link,
			Authors:     gormschema.FeedPersonArrayVal(f.Authors),
			Language:    f.Language,
		}
		if f.Image != nil {
			img := gormschema.FeedImageVal(*f.Image)
			feed.Image = &img
		}

		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"title", "description", "link", "authors", "language", "image", "updated_at"}),
		}).Create(&feed).Error
	})
}

func (k *KetherRepo) CheckNewFeedItems(
	ctx context.Context,
	items []*modelfeed.Item,
	feedID model.InternalID,
) ([]string, error) {
	guids := make([]string, 0, len(items))
	for _, item := range items {
		guids = append(guids, item.GUID)
	}

	var existItems []gormschema.FeedItem
	if err := k.data.db.WithContext(ctx).
		Select("guid").
		Where("feed_id = ? AND guid IN ?", feedID, guids).
		Find(&existItems).Error; err != nil {
		return nil, err
	}

	existItemMap := make(map[string]bool)
	for _, item := range existItems {
		existItemMap[item.GUID] = true
	}

	res := make([]string, 0, len(items)-len(existItems))
	for _, item := range items {
		if _, exist := existItemMap[item.GUID]; !exist {
			res = append(res, item.GUID)
		}
	}
	return res, nil
}

func (k *KetherRepo) UpsertFeedItems(
	ctx context.Context,
	items []*modelfeed.Item,
	feedID model.InternalID,
) error {
	feedItems := make([]gormschema.FeedItem, len(items))
	for i, item := range items {
		feedItems[i] = gormschema.FeedItem{
			ID:                item.ID,
			FeedID:            feedID,
			Title:             item.Title,
			Description:       item.Description,
			Content:           item.Content,
			Link:              item.Link,
			Updated:           item.Updated,
			UpdatedParsed:     item.UpdatedParsed,
			Published:         item.Published,
			Authors:           gormschema.FeedPersonArrayVal(item.Authors),
			GUID:              item.GUID,
			Enclosures:        gormschema.FeedEnclosureArrayVal(item.Enclosures),
			PublishPlatform:   item.PublishPlatform,
			DigestDescription: item.DigestDescription,
			DigestImages:      gormschema.FeedImageArrayVal(item.DigestImages),
		}
		if item.Image != nil {
			img := gormschema.FeedImageVal(*item.Image)
			feedItems[i].Image = &img
		}
		if item.PublishedParsed != nil {
			feedItems[i].PublishedParsed = *item.PublishedParsed
		} else {
			feedItems[i].PublishedParsed = time.Now()
		}
	}

	return k.data.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "feed_id"}, {Name: "guid"}},
		DoNothing: true,
	}).Create(&feedItems).Error
}

func (k *KetherRepo) UpdateFeedPullStatus(ctx context.Context, conf *modelyesod.FeedConfig) error {
	var c gormschema.FeedConfig
	if err := k.data.db.WithContext(ctx).Where("id = ?", conf.ID).First(&c).Error; err != nil {
		return err
	}
	return k.data.db.WithContext(ctx).Model(&gormschema.FeedConfig{}).
		Where("id = ?", conf.ID).
		Updates(map[string]any{
			"latest_pull_at":      conf.LatestPullTime,
			"latest_pull_status":  gormschema.ToSchemaFeedConfigPullStatus(conf.LatestPullStatus),
			"latest_pull_message": conf.LatestPullMessage,
			"next_pull_begin_at":  conf.LatestPullTime.Add(c.PullInterval),
		}).Error
}

func (k *KetherRepo) GetFeedItem(ctx context.Context, id model.InternalID) (*modelfeed.Item, error) {
	var item gormschema.FeedItem
	if err := k.data.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizFeedItem(&item), nil
}

func (k *KetherRepo) GetFeedActions(ctx context.Context, id model.InternalID) ([]*modelyesod.FeedActionSet, error) {
	var actions []gormschema.FeedActionSet
	if err := k.data.db.WithContext(ctx).
		Joins("JOIN feed_config_actions ON feed_action_sets.id = feed_config_actions.feed_action_set_id").
		Where("feed_config_actions.feed_config_id = ?", id).
		Find(&actions).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizFeedActionSetList(ptrSlice(actions)), nil
}

func (k *KetherRepo) GetNotifyTargetItems(
	ctx context.Context,
	id model.InternalID,
	paging model.Paging,
) (*model.FeatureRequest, []*modelfeed.Item, error) {
	var fr *model.FeatureRequest
	var it []*modelfeed.Item

	err := k.data.WithTx(ctx, func(tx *gorm.DB) error {
		var target gormschema.NotifyTarget
		if err := tx.First(&target, id).Error; err != nil {
			return err
		}
		if target.Destination != nil {
			dest := model.FeatureRequest(*target.Destination)
			fr = &dest
		}

		// Get flow IDs for this target
		var flowTargets []gormschema.NotifyFlowTarget
		if err := tx.Where("notify_target_id = ?", id).Find(&flowTargets).Error; err != nil {
			return err
		}
		flowIDs := make([]model.InternalID, len(flowTargets))
		for i, ft := range flowTargets {
			flowIDs[i] = ft.NotifyFlowID
		}

		if len(flowIDs) == 0 {
			return nil
		}

		var items []gormschema.FeedItem
		if err := tx.
			Joins("JOIN feed_item_collection_feed_items ON feed_items.id = feed_item_collection_feed_items.feed_item_id").
			Where("feed_item_collection_feed_items.feed_item_collection_id IN ?", flowIDs).
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			Find(&items).Error; err != nil {
			return err
		}
		it = gormschema.ToBizFeedItemList(ptrSlice(items))
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return fr, it, nil
}

func (k *KetherRepo) AddFeedItemsToCollection(
	ctx context.Context,
	collectionID model.InternalID,
	itemIDs []model.InternalID,
) error {
	items := make([]gormschema.FeedItemCollectionItem, len(itemIDs))
	for i, id := range itemIDs {
		items[i] = gormschema.FeedItemCollectionItem{
			FeedItemCollectionID: collectionID,
			FeedItemID:           id,
		}
	}
	return k.data.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&items).Error
}

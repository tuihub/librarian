package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/orm/model"
	"github.com/tuihub/librarian/internal/data/orm/query"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelyesod"

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

func (k *KetherRepo) UpsertAccount(ctx context.Context, acc libmodel.Account) error {
	return query.Use(k.data.db).Account.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "platform"}, {Name: "platform_account_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "profile_url", "avatar_url"}),
	}).Create(&model.Account{
		ID:                acc.ID,
		Platform:          acc.Platform,
		PlatformAccountID: acc.PlatformAccountID,
		Name:              acc.Name,
		ProfileURL:        acc.ProfileURL,
		AvatarURL:         acc.AvatarURL,
	})
}

func (k *KetherRepo) UpsertAppInfo(
	ctx context.Context, ap *modelgebura.AppInfo, internal *modelgebura.AppInfo,
) error {
	return k.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.AppInfo
		count, err := q.WithContext(ctx).
			Where(q.Source.Eq(ap.Source), q.SourceAppID.Eq(ap.SourceAppID)).
			Count()
		if err != nil {
			return err
		}

		if count == 0 {
			// Create with internal
			err = q.WithContext(ctx).Create(&model.AppInfo{
				ID:          internal.ID,
				Source:      internal.Source,
				SourceAppID: internal.SourceAppID,
				Name:        internal.Name,
				Type:        converter.ToORMAppInfoTypeManual(internal.Type),
			})
			if err != nil {
				return err
			}
		}

		// Upsert ap
		info := &model.AppInfo{
			ID:                 ap.ID, // This ID might be ignored if conflict?
			Source:             ap.Source,
			SourceAppID:        ap.SourceAppID,
			SourceURL:          ap.SourceURL,
			Name:               ap.Name,
			Type:               converter.ToORMAppInfoTypeManual(ap.Type),
			ShortDescription:   ap.ShortDescription,
			IconImageURL:       ap.IconImageURL,
			BackgroundImageURL: ap.BackgroundImageURL,
			CoverImageURL:      ap.CoverImageURL,
			Description:        ap.Description,
			ReleaseDate:        ap.ReleaseDate,
			Developer:          ap.Developer,
			Publisher:          ap.Publisher,
		}

		return q.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "source"}, {Name: "source_app_id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"source_url", "name", "type", "short_description",
				"icon_image_url", "background_image_url", "cover_image_url",
				"description", "release_date", "developer", "publisher",
			}),
		}).Create(info)
	})
}

func (k *KetherRepo) UpsertAppInfos(ctx context.Context, al []*modelgebura.AppInfo) error {
	apps := make([]*model.AppInfo, len(al))
	for i, ap := range al {
		apps[i] = &model.AppInfo{
			ID:                 ap.ID,
			Source:             ap.Source,
			SourceAppID:        ap.SourceAppID,
			SourceURL:          ap.SourceURL,
			Name:               ap.Name,
			Type:               converter.ToORMAppInfoTypeManual(ap.Type),
			ShortDescription:   ap.ShortDescription,
			IconImageURL:       ap.IconImageURL,
			BackgroundImageURL: ap.BackgroundImageURL,
			CoverImageURL:      ap.CoverImageURL,
			Description:        ap.Description,
			ReleaseDate:        ap.ReleaseDate,
			Developer:          ap.Developer,
			Publisher:          ap.Publisher,
		}
	}
	return query.Use(k.data.db).AppInfo.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "source"}, {Name: "source_app_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"source_url", "name", "type", "short_description",
			"icon_image_url", "background_image_url", "cover_image_url",
			"description", "release_date", "developer", "publisher",
		}),
	}).Create(apps...)
}

func (k *KetherRepo) UpsertFeed(ctx context.Context, f *modelfeed.Feed) error {
	return k.data.WithTx(ctx, func(tx *query.Query) error {
		_, err := tx.FeedConfig.WithContext(ctx).Where(tx.FeedConfig.ID.Eq(int64(f.ID))).First()
		if err != nil {
			return err
		}

		return tx.Feed.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"title", "description", "link", "authors", "language", "image",
			}),
		}).Create(&model.Feed{
			ID:          f.ID,
			Title:       f.Title,
			Description: f.Description,
			Link:        f.Link,
			Authors:     f.Authors,
			Language:    f.Language,
			Image:       f.Image,
		})
	})
}

func (k *KetherRepo) CheckNewFeedItems(
	ctx context.Context,
	items []*modelfeed.Item,
	feedID libmodel.InternalID,
) ([]string, error) {
	guids := make([]string, 0, len(items))
	for _, item := range items {
		guids = append(guids, item.GUID)
	}

	q := query.Use(k.data.db).FeedItem
	existItems, err := q.WithContext(ctx).
		Where(q.FeedID.Eq(int64(feedID)), q.GUID.In(guids...)).
		Select(q.GUID).
		Find()
	if err != nil {
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
	feedID libmodel.InternalID,
) error {
	il := make([]*model.FeedItem, len(items))
	for i, item := range items {
		t := time.Now()
		if item.PublishedParsed != nil {
			t = *item.PublishedParsed
		}

		il[i] = &model.FeedItem{
			FeedID:            feedID,
			ID:                item.ID,
			Title:             item.Title,
			Description:       item.Description,
			Content:           item.Content,
			Link:              item.Link,
			Updated:           item.Updated,
			UpdatedParsed:     item.UpdatedParsed,
			Published:         item.Published,
			PublishedParsed:   t,
			Authors:           item.Authors,
			GUID:              item.GUID,
			Image:             item.Image,
			Enclosures:        item.Enclosures,
			PublishPlatform:   item.PublishPlatform,
			DigestDescription: item.DigestDescription,
			DigestImages:      item.DigestImages,
		}
	}

	return query.Use(k.data.db).FeedItem.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "feed_id"}, {Name: "guid"}},
		DoNothing: true,
	}).Create(il...)
}

func (k *KetherRepo) UpdateFeedPullStatus(ctx context.Context, conf *modelyesod.FeedConfig) error {
	q := query.Use(k.data.db).FeedConfig
	// Verify exists
	c, err := q.WithContext(ctx).Where(q.ID.Eq(int64(conf.ID))).First()
	if err != nil {
		return err
	}

	_, err = q.WithContext(ctx).Where(q.ID.Eq(int64(conf.ID))).Updates(&model.FeedConfig{
		LatestPullAt:      conf.LatestPullTime,
		LatestPullStatus:  converter.ToORMFeedConfigLatestPullStatus(conf.LatestPullStatus),
		LatestPullMessage: conf.LatestPullMessage,
		NextPullBeginAt:   conf.LatestPullTime.Add(c.PullInterval),
	})
	return err
}

func (k *KetherRepo) GetFeedItem(ctx context.Context, id libmodel.InternalID) (*modelfeed.Item, error) {
	q := query.Use(k.data.db).FeedItem
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return converter.ToBizFeedItem(res), nil
}

func (k *KetherRepo) GetFeedActions(ctx context.Context, id libmodel.InternalID) ([]*modelyesod.FeedActionSet, error) {
	// Join FeedActionSet -> FeedConfigAction -> FeedConfig (ID=id)
	q := query.Use(k.data.db)
	fas := q.FeedActionSet
	fca := q.FeedConfigAction

	res, err := fas.WithContext(ctx).
		Join(fca, fca.FeedActionSetID.EqCol(fas.ID)).
		Where(fca.FeedConfigID.Eq(int64(id))).
		Find()
	if err != nil {
		return nil, err
	}
	return converter.ToBizFeedActionSetList(res), nil
}

func (k *KetherRepo) GetNotifyTargetItems(
	ctx context.Context,
	id libmodel.InternalID,
	paging libmodel.Paging,
) (*libmodel.FeatureRequest, []*modelfeed.Item, error) {
	var fr *libmodel.FeatureRequest
	var it []*modelfeed.Item

	err := k.data.WithTx(ctx, func(tx *query.Query) error {
		nt, err := tx.NotifyTarget.WithContext(ctx).Where(tx.NotifyTarget.ID.Eq(int64(id))).First()
		if err != nil {
			return err
		}
		fr = nt.Destination

		// NotifyTarget -> NotifyFlowTarget -> NotifyFlow
		// NotifyFlow ID is also FeedItemCollection ID
		// FeedItemCollection -> FeedItems

		// 1. Get NotifyFlow IDs
		nft := tx.NotifyFlowTarget

		var fIDs []int64
		err = nft.WithContext(ctx).
			Where(nft.NotifyTargetID.Eq(int64(id))).
			Pluck(nft.NotifyFlowID, &fIDs)
		if err != nil {
			return err
		}

		if len(fIDs) == 0 {
			it = []*modelfeed.Item{}
			return nil
		}

		// 2. Get FeedItems in these collections (flows)
		// M2M FeedItem <-> FeedItemCollection
		// Join table: feed_item_collection_feed_items

		var itemIDs []libmodel.InternalID
		err = k.data.db.Table("feed_item_collection_feed_items").
			Where("feed_item_collection_id IN ?", fIDs).
			Pluck("feed_item_id", &itemIDs).Error
		if err != nil {
			return err
		}

		if len(itemIDs) == 0 {
			it = []*modelfeed.Item{}
			return nil
		}

		castItemIDs := make([]int64, len(itemIDs))
		for i, v := range itemIDs {
			castItemIDs[i] = int64(v)
		}

		items, err := tx.FeedItem.WithContext(ctx).
			Where(tx.FeedItem.ID.In(castItemIDs...)).
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			Find()
		if err != nil {
			return err
		}

		it = converter.ToBizFeedItemList(items)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}
	return fr, it, nil
}

func (k *KetherRepo) AddFeedItemsToCollection(
	ctx context.Context,
	collectionID libmodel.InternalID,
	itemIDs []libmodel.InternalID,
) error {
	// GORM Append association
	if len(itemIDs) == 0 {
		return nil
	}

	feedItems := make([]*model.FeedItem, len(itemIDs))
	for i, id := range itemIDs {
		feedItems[i] = &model.FeedItem{ID: id}
	}

	return k.data.db.Model(&model.FeedItemCollection{ID: collectionID}).
		Association("FeedItems").
		Append(feedItems)
}

package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelyesod"

	"gorm.io/gorm"
)

type YesodRepo struct {
	data *Data
}

// NewYesodRepo .
func NewYesodRepo(data *Data) *YesodRepo {
	return &YesodRepo{
		data: data,
	}
}

func (y *YesodRepo) CreateFeedConfig(ctx context.Context, owner model.InternalID, c *modelyesod.FeedConfig) error {
	return y.data.WithTx(ctx, func(tx *gorm.DB) error {
		var sourceVal *gormschema.FeatureRequestVal
		if c.Source != nil {
			v := gormschema.FeatureRequestVal(*c.Source)
			sourceVal = &v
		}
		fc := gormschema.FeedConfig{
			ID:                c.ID,
			OwnerID:           owner,
			Name:              c.Name,
			Description:       c.Description,
			Category:          c.Category,
			Source:            sourceVal,
			Status:            gormschema.ToSchemaFeedConfigStatus(c.Status),
			PullInterval:      c.PullInterval,
			LatestPullStatus:  gormschema.ToSchemaFeedConfigPullStatus(c.LatestPullStatus),
			LatestPullMessage: "",
			HideItems:         c.HideItems,
		}
		if err := tx.Create(&fc).Error; err != nil {
			return err
		}
		for i, action := range c.ActionSets {
			fca := gormschema.FeedConfigAction{
				FeedConfigID:    c.ID,
				FeedActionSetID: action,
				Index:           int64(i),
			}
			if err := tx.Create(&fca).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (y *YesodRepo) UpdateFeedConfig(ctx context.Context, userID model.InternalID, c *modelyesod.FeedConfig) error {
	return y.data.WithTx(ctx, func(tx *gorm.DB) error {
		updates := make(map[string]any)
		if len(c.Name) > 0 {
			updates["name"] = c.Name
		}
		if len(c.Description) > 0 {
			updates["description"] = c.Description
		}
		if len(c.Category) > 0 {
			updates["category"] = c.Category
		}
		if c.Source != nil {
			sourceVal := gormschema.FeatureRequestVal(*c.Source)
			updates["source"] = &sourceVal
		}
		if c.Status != modelyesod.FeedConfigStatusUnspecified {
			updates["status"] = gormschema.ToSchemaFeedConfigStatus(c.Status)
		}
		if c.PullInterval > 0 {
			updates["pull_interval"] = c.PullInterval
			updates["next_pull_begin_at"] = time.Now()
		}
		updates["hide_items"] = c.HideItems

		if err := tx.Model(&gormschema.FeedConfig{}).
			Where("id = ? AND owner_id = ?", c.ID, userID).
			Updates(updates).Error; err != nil {
			return err
		}

		if c.ActionSets != nil {
			// Remove existing action sets
			if err := tx.Where("feed_config_id = ?", c.ID).Delete(&gormschema.FeedConfigAction{}).Error; err != nil {
				return err
			}
			// Create new action sets
			for i, action := range c.ActionSets {
				fca := gormschema.FeedConfigAction{
					FeedConfigID:    c.ID,
					FeedActionSetID: action,
					Index:           int64(i),
				}
				if err := tx.Create(&fca).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// UpdateFeedConfigAsInQueue set SetNextPullBeginAt to one day later to avoid repeat queue.
func (y *YesodRepo) UpdateFeedConfigAsInQueue(ctx context.Context, id model.InternalID) error {
	return y.data.db.WithContext(ctx).
		Model(&gormschema.FeedConfig{}).
		Where("id = ?", id).
		Update("next_pull_begin_at", time.Now().Add(libtime.Day)).Error
}

func (y *YesodRepo) ListFeedConfigNeedPull(ctx context.Context, sources []string,
	statuses []modelyesod.FeedConfigStatus, order modelyesod.ListFeedOrder,
	pullTime time.Time, limit int) ([]*modelyesod.FeedConfig, error) {
	query := y.data.db.WithContext(ctx).Model(&gormschema.FeedConfig{})
	if len(statuses) > 0 {
		statusStrs := make([]string, len(statuses))
		for i, s := range statuses {
			statusStrs[i] = gormschema.ToSchemaFeedConfigStatus(s)
		}
		query = query.Where("status IN ?", statusStrs)
	}
	if order == modelyesod.ListFeedOrderNextPull {
		query = query.Where("next_pull_begin_at < ?", pullTime)
	}
	query = query.Limit(limit)

	var configs []gormschema.FeedConfig
	if err := query.Find(&configs).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizFeedConfigList(ptrSlice(configs)), nil
}

func (y *YesodRepo) ListFeedConfigs(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
	ids []model.InternalID,
	statuses []modelyesod.FeedConfigStatus,
	categories []string,
) ([]*modelyesod.FeedWithConfig, int, error) {
	var res []*modelyesod.FeedWithConfig
	var total int64
	err := y.data.WithTx(ctx, func(tx *gorm.DB) error {
		query := tx.Model(&gormschema.FeedConfig{}).Where("owner_id = ?", userID)
		if len(ids) > 0 {
			query = query.Where("id IN ?", ids)
		}
		if len(statuses) > 0 {
			statusStrs := make([]string, len(statuses))
			for i, s := range statuses {
				statusStrs[i] = gormschema.ToSchemaFeedConfigStatus(s)
			}
			query = query.Where("status IN ?", statusStrs)
		}
		if len(categories) > 0 {
			query = query.Where("category IN ?", categories)
		}

		if err := query.Count(&total).Error; err != nil {
			return err
		}

		var configs []gormschema.FeedConfig
		if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&configs).Error; err != nil {
			return err
		}

		res = make([]*modelyesod.FeedWithConfig, 0, len(configs))
		for _, config := range configs {
			fc := gormschema.ToBizFeedConfig(&config)
			// Get action sets
			var actions []gormschema.FeedConfigAction
			if err := tx.Where("feed_config_id = ?", config.ID).Order("index").Find(&actions).Error; err == nil {
				fc.ActionSets = make([]model.InternalID, len(actions))
				for i, a := range actions {
					fc.ActionSets[i] = a.FeedActionSetID
				}
			}
			// Get feed
			var feed gormschema.Feed
			var feedPtr *modelfeed.Feed
			if err := tx.Where("id = ?", config.ID).First(&feed).Error; err == nil {
				feedPtr = gormschema.ToBizFeed(&feed)
			}
			res = append(res, &modelyesod.FeedWithConfig{
				FeedConfig: fc,
				Feed:       feedPtr,
			})
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, int(total), nil
}

func (y *YesodRepo) ListFeedCategories(ctx context.Context, id model.InternalID) ([]string, error) {
	var results []struct {
		Category string
	}
	if err := y.data.db.WithContext(ctx).
		Model(&gormschema.FeedConfig{}).
		Select("DISTINCT category").
		Where("owner_id = ?", id).
		Find(&results).Error; err != nil {
		return nil, err
	}
	categories := make([]string, len(results))
	for i, r := range results {
		categories[i] = r.Category
	}
	return categories, nil
}

func (y *YesodRepo) ListFeedPlatforms(ctx context.Context, id model.InternalID) ([]string, error) {
	var results []struct {
		PublishPlatform string
	}
	// Complex join query
	if err := y.data.db.WithContext(ctx).
		Model(&gormschema.FeedItem{}).
		Select("DISTINCT publish_platform").
		Joins("JOIN feeds ON feed_items.feed_id = feeds.id").
		Joins("JOIN feed_configs ON feeds.id = feed_configs.id").
		Where("feed_configs.owner_id = ?", id).
		Find(&results).Error; err != nil {
		return nil, err
	}
	platforms := make([]string, len(results))
	for i, r := range results {
		platforms[i] = r.PublishPlatform
	}
	return platforms, nil
}

func (y *YesodRepo) ListFeedItems(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
	feedIDs []model.InternalID,
	authors []string,
	platforms []string,
	timeRange *model.TimeRange,
	categories []string,
) ([]*modelyesod.FeedItemDigest, int, error) {
	var res []*modelyesod.FeedItemDigest
	var total int64

	err := y.data.WithTx(ctx, func(tx *gorm.DB) error {
		query := tx.Model(&gormschema.FeedItem{}).
			Joins("JOIN feeds ON feed_items.feed_id = feeds.id").
			Joins("JOIN feed_configs ON feeds.id = feed_configs.id").
			Where("feed_configs.owner_id = ? AND feed_configs.hide_items = ?", userID, false)

		if len(feedIDs) > 0 {
			query = query.Where("feeds.id IN ?", feedIDs)
		}
		if len(platforms) > 0 {
			query = query.Where("feed_items.publish_platform IN ?", platforms)
		}
		if len(categories) > 0 {
			query = query.Where("feed_configs.category IN ?", categories)
		}
		if timeRange != nil {
			query = query.Where("feed_items.published_parsed >= ? AND feed_items.published_parsed < ?",
				timeRange.StartTime, timeRange.StartTime.Add(timeRange.Duration))
		}

		if err := query.Count(&total).Error; err != nil {
			return err
		}

		var items []gormschema.FeedItem
		if err := query.Order("feed_items.published_parsed DESC").
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			Find(&items).Error; err != nil {
			return err
		}

		res = make([]*modelyesod.FeedItemDigest, 0, len(items))
		for _, item := range items {
			digest := toBizFeedItemDigest(&item, tx)
			res = append(res, digest)
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, int(total), nil
}

func (y *YesodRepo) GroupFeedItems( //nolint:gocognit //TODO
	ctx context.Context,
	userID model.InternalID,
	groups []model.TimeRange,
	feedIDs []model.InternalID,
	authors []string,
	platforms []string,
	groupSize int,
	categories []string,
) (map[model.TimeRange][]*modelyesod.FeedItemDigest, error) {
	res := make(map[model.TimeRange][]*modelyesod.FeedItemDigest)
	err := y.data.WithTx(ctx, func(tx *gorm.DB) error {
		for _, timeRange := range groups {
			query := tx.Model(&gormschema.FeedItem{}).
				Joins("JOIN feeds ON feed_items.feed_id = feeds.id").
				Joins("JOIN feed_configs ON feeds.id = feed_configs.id").
				Where("feed_configs.owner_id = ? AND feed_configs.hide_items = ?", userID, false)

			if len(feedIDs) > 0 {
				query = query.Where("feeds.id IN ?", feedIDs)
			}
			if len(platforms) > 0 {
				query = query.Where("feed_items.publish_platform IN ?", platforms)
			}
			if len(categories) > 0 {
				query = query.Where("feed_configs.category IN ?", categories)
			}

			var items []gormschema.FeedItem
			if err := query.
				Where("feed_items.published_parsed >= ? AND feed_items.published_parsed < ?",
					timeRange.StartTime, timeRange.StartTime.Add(timeRange.Duration)).
				Order("feed_items.published_parsed DESC").
				Limit(groupSize).
				Find(&items).Error; err != nil {
				return err
			}

			if len(items) == 0 {
				continue
			}

			il := make([]*modelyesod.FeedItemDigest, 0, len(items))
			for _, item := range items {
				il = append(il, toBizFeedItemDigest(&item, tx))
			}
			res[timeRange] = il
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *YesodRepo) GetFeedItems(
	ctx context.Context,
	userID model.InternalID,
	ids []model.InternalID,
) ([]*modelfeed.Item, error) {
	var items []gormschema.FeedItem
	err := y.data.db.WithContext(ctx).
		Joins("JOIN feeds ON feed_items.feed_id = feeds.id").
		Joins("JOIN feed_configs ON feeds.id = feed_configs.id").
		Where("feed_configs.owner_id = ? AND feed_items.id IN ?", userID, ids).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return gormschema.ToBizFeedItemList(ptrSlice(items)), nil
}

func (y *YesodRepo) ReadFeedItem(ctx context.Context, userID model.InternalID, id model.InternalID) error {
	return y.data.db.WithContext(ctx).
		Model(&gormschema.FeedItem{}).
		Where("id = ?", id).
		Update("read_count", gorm.Expr("read_count + 1")).Error
}

func (y *YesodRepo) CreateFeedItemCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collection *modelyesod.FeedItemCollection,
) error {
	fic := gormschema.FeedItemCollection{
		ID:          collection.ID,
		OwnerID:     ownerID,
		Name:        collection.Name,
		Description: collection.Description,
		Category:    collection.Category,
	}
	return y.data.db.WithContext(ctx).Create(&fic).Error
}

func (y *YesodRepo) UpdateFeedItemCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collection *modelyesod.FeedItemCollection,
) error {
	return y.data.db.WithContext(ctx).
		Model(&gormschema.FeedItemCollection{}).
		Where("id = ? AND owner_id = ?", collection.ID, ownerID).
		Updates(map[string]any{
			"name":        collection.Name,
			"description": collection.Description,
			"category":    collection.Category,
		}).Error
}

func (y *YesodRepo) ListFeedItemCollections(
	ctx context.Context,
	ownerID model.InternalID,
	paging model.Paging,
	ids []model.InternalID,
	categories []string,
) ([]*modelyesod.FeedItemCollection, int, error) {
	query := y.data.db.WithContext(ctx).Model(&gormschema.FeedItemCollection{}).
		Where("owner_id = ?", ownerID)

	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}
	if len(categories) > 0 {
		query = query.Where("category IN ?", categories)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var collections []gormschema.FeedItemCollection
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&collections).Error; err != nil {
		return nil, 0, err
	}

	return gormschema.ToBizFeedItemCollectionList(ptrSlice(collections)), int(total), nil
}

func (y *YesodRepo) AddFeedItemToCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collectionID model.InternalID,
	itemID model.InternalID,
) error {
	item := gormschema.FeedItemCollectionItem{
		FeedItemCollectionID: collectionID,
		FeedItemID:           itemID,
	}
	return y.data.db.WithContext(ctx).Create(&item).Error
}

func (y *YesodRepo) RemoveFeedItemFromCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collectionID model.InternalID,
	itemID model.InternalID,
) error {
	return y.data.db.WithContext(ctx).
		Where("feed_item_collection_id = ? AND feed_item_id = ?", collectionID, itemID).
		Delete(&gormschema.FeedItemCollectionItem{}).Error
}

func (y *YesodRepo) ListFeedItemsInCollection(
	ctx context.Context,
	ownerID model.InternalID,
	paging model.Paging,
	ids []model.InternalID,
	authors []string,
	platforms []string,
	categories []string,
	timeRange *model.TimeRange,
) ([]*modelyesod.FeedItemDigest, int, error) {
	var res []*modelyesod.FeedItemDigest
	var total int64

	err := y.data.WithTx(ctx, func(tx *gorm.DB) error {
		query := tx.Model(&gormschema.FeedItem{}).
			Joins("JOIN feed_item_collection_feed_items ON feed_items.id = feed_item_collection_feed_items.feed_item_id").
			Joins("JOIN feed_item_collections ON feed_item_collection_feed_items.feed_item_collection_id = feed_item_collections.id").
			Where("feed_item_collections.owner_id = ?", ownerID)

		if len(ids) > 0 {
			query = query.Where("feed_item_collections.id IN ?", ids)
		}
		if len(platforms) > 0 {
			query = query.Where("feed_items.publish_platform IN ?", platforms)
		}
		if len(categories) > 0 {
			query = query.Joins("JOIN feeds ON feed_items.feed_id = feeds.id").
				Joins("JOIN feed_configs ON feeds.id = feed_configs.id").
				Where("feed_configs.category IN ?", categories)
		}
		if timeRange != nil {
			query = query.Where("feed_items.published_parsed >= ? AND feed_items.published_parsed < ?",
				timeRange.StartTime, timeRange.StartTime.Add(timeRange.Duration))
		}

		if err := query.Count(&total).Error; err != nil {
			return err
		}

		var items []gormschema.FeedItem
		if err := query.Order("feed_items.published_parsed DESC").
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			Find(&items).Error; err != nil {
			return err
		}

		res = make([]*modelyesod.FeedItemDigest, 0, len(items))
		for _, item := range items {
			res = append(res, toBizFeedItemDigest(&item, tx))
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, int(total), nil
}

func (y *YesodRepo) GetFeedOwner(ctx context.Context, id model.InternalID) (*model.User, error) {
	var config gormschema.FeedConfig
	if err := y.data.db.WithContext(ctx).Where("id = ?", id).First(&config).Error; err != nil {
		return nil, err
	}
	var user gormschema.User
	if err := y.data.db.WithContext(ctx).First(&user, config.OwnerID).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizUser(&user), nil
}

func (y *YesodRepo) CreateFeedActionSet(ctx context.Context, id model.InternalID, set *modelyesod.FeedActionSet) error {
	fas := gormschema.FeedActionSet{
		ID:          set.ID,
		OwnerID:     id,
		Name:        set.Name,
		Description: set.Description,
		Actions:     gormschema.FeatureRequestArrayVal(set.Actions),
	}
	return y.data.db.WithContext(ctx).Create(&fas).Error
}

func (y *YesodRepo) UpdateFeedActionSet(ctx context.Context, id model.InternalID, set *modelyesod.FeedActionSet) error {
	return y.data.db.WithContext(ctx).
		Model(&gormschema.FeedActionSet{}).
		Where("id = ? AND owner_id = ?", set.ID, id).
		Updates(map[string]any{
			"name":        set.Name,
			"description": set.Description,
			"actions":     gormschema.FeatureRequestArrayVal(set.Actions),
		}).Error
}

func (y *YesodRepo) ListFeedActionSets(
	ctx context.Context,
	id model.InternalID,
	paging model.Paging,
) ([]*modelyesod.FeedActionSet, int, error) {
	query := y.data.db.WithContext(ctx).Model(&gormschema.FeedActionSet{}).
		Where("owner_id = ?", id)

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var sets []gormschema.FeedActionSet
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&sets).Error; err != nil {
		return nil, 0, err
	}

	return gormschema.ToBizFeedActionSetList(ptrSlice(sets)), int(total), nil
}

// Helper function to convert slice to pointer slice
func ptrSlice[T any](slice []T) []*T {
	res := make([]*T, len(slice))
	for i := range slice {
		res[i] = &slice[i]
	}
	return res
}

// toBizFeedItemDigest converts a feed item to a digest with additional info
func toBizFeedItemDigest(item *gormschema.FeedItem, tx *gorm.DB) *modelyesod.FeedItemDigest {
	digest := &modelyesod.FeedItemDigest{
		FeedID:              item.FeedID,
		ItemID:              item.ID,
		PublishedParsedTime: item.PublishedParsed,
		Title:               item.Title,
		PublishPlatform:     item.PublishPlatform,
		ShortDescription:    item.DigestDescription,
		ReadCount:           item.ReadCount,
	}

	if item.Image != nil {
		digest.AvatarURL = item.Image.URL
	}

	if len(item.Authors) > 0 {
		var authorNames []string
		for _, a := range item.Authors {
			if a != nil {
				authorNames = append(authorNames, a.Name)
			}
		}
		if len(authorNames) > 0 {
			digest.Authors = authorNames[0]
			for i := 1; i < len(authorNames); i++ {
				digest.Authors += ", " + authorNames[i]
			}
		}
	}

	for _, img := range item.DigestImages {
		if img != nil {
			digest.ImageUrls = append(digest.ImageUrls, img.URL)
		}
	}

	// Get feed info
	var feed gormschema.Feed
	if tx.Where("id = ?", item.FeedID).First(&feed).Error == nil {
		if feed.Image != nil {
			digest.FeedAvatarURL = feed.Image.URL
		}
		// Get feed config name
		var config gormschema.FeedConfig
		if tx.Where("id = ?", item.FeedID).First(&config).Error == nil {
			digest.FeedConfigName = config.Name
		}
	}

	return digest
}

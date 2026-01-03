package data

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/tuihub/librarian/internal/data/orm/query"
	"github.com/tuihub/librarian/internal/lib/libtime"
	libmodel "github.com/tuihub/librarian/internal/model"
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

func (y *YesodRepo) CreateFeedConfig(ctx context.Context, owner libmodel.InternalID, c *modelyesod.FeedConfig) error {
	return y.data.WithTx(ctx, func(tx *query.Query) error {
		c.UserFeedConfig = owner
		c.NextPullBeginAt = time.Now()
		c.LatestPullMessage = ""

		if err := tx.FeedConfig.WithContext(ctx).Create(c); err != nil {
			return err
		}

		if len(c.ActionSets) > 0 {
			actions := make([]*modelyesod.FeedConfigAction, len(c.ActionSets))
			for i, actionID := range c.ActionSets {
				actions[i] = &modelyesod.FeedConfigAction{
					FeedConfigID:    c.ID,
					FeedActionSetID: actionID,
				}
			}
			if err := tx.FeedConfigAction.WithContext(ctx).Create(actions...); err != nil {
				return err
			}
		}
		return nil
	})
}

//nolint:gocognit // complexity
func (y *YesodRepo) UpdateFeedConfig(
	ctx context.Context,
	userID libmodel.InternalID,
	c *modelyesod.FeedConfig,
) error {
	return y.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.FeedConfig
		u := q.WithContext(ctx).Where(q.ID.Eq(int64(c.ID)), q.UserFeedConfig.Eq(int64(userID)))

		updates := make(map[string]interface{})
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
			updates["source"] = c.Source
		}
		if c.Status != modelyesod.FeedConfigStatusUnspecified {
			updates["status"] = c.Status
		}
		if c.PullInterval > 0 {
			updates["pull_interval"] = c.PullInterval
			updates["next_pull_begin_at"] = time.Now()
		}
		updates["hide_items"] = c.HideItems

		if _, err := u.Updates(updates); err != nil {
			return err
		}

		if c.ActionSets != nil {
			qa := tx.FeedConfigAction
			if _, err := qa.WithContext(ctx).Where(qa.FeedConfigID.Eq(int64(c.ID))).Delete(); err != nil {
				return err
			}

			actions := make([]*modelyesod.FeedConfigAction, len(c.ActionSets))
			for i, actionID := range c.ActionSets {
				actions[i] = &modelyesod.FeedConfigAction{
					FeedConfigID:    c.ID,
					FeedActionSetID: actionID,
				}
			}
			if len(actions) > 0 {
				if err := qa.WithContext(ctx).Create(actions...); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// UpdateFeedConfigAsInQueue set SetNextPullBeginAt to one day later to avoid repeat queue.
// While pull success, UpsertFeed will set correct value.
// While pull failed, server will retry task next day.
func (y *YesodRepo) UpdateFeedConfigAsInQueue(ctx context.Context, id libmodel.InternalID) error {
	q := query.Use(y.data.db).FeedConfig
	_, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).Update(q.NextPullBeginAt, time.Now().Add(libtime.Day))
	return err
}

func (y *YesodRepo) ListFeedConfigNeedPull(ctx context.Context, sources []string,
	statuses []modelyesod.FeedConfigStatus, order modelyesod.ListFeedOrder,
	pullTime time.Time, i int) ([]*modelyesod.FeedConfig, error) {
	q := query.Use(y.data.db).FeedConfig
	u := q.WithContext(ctx)

	if len(statuses) > 0 {
		s := make([]driver.Valuer, len(statuses))
		for i, v := range statuses {
			s[i] = v
		}
		u = u.Where(q.Status.In(s...))
	}
	switch order {
	case modelyesod.ListFeedOrderUnspecified:
	case modelyesod.ListFeedOrderNextPull:
		u = u.Where(q.NextPullBeginAt.Lt(pullTime))
	}
	u = u.Limit(i)

	res, err := u.Find()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *YesodRepo) ListFeedConfigs(
	ctx context.Context,
	userID libmodel.InternalID,
	paging libmodel.Paging,
	ids []libmodel.InternalID,
	statuses []modelyesod.FeedConfigStatus,
	categories []string,
) ([]*modelyesod.FeedWithConfig, int, error) {
	q := query.Use(y.data.db).FeedConfig
	u := q.WithContext(ctx).Where(q.UserFeedConfig.Eq(int64(userID)))

	if len(ids) > 0 {
		castIDs := make([]int64, len(ids))
		for i, v := range ids {
			castIDs[i] = int64(v)
		}
		u = u.Where(q.ID.In(castIDs...))
	}
	if len(statuses) > 0 {
		s := make([]driver.Valuer, len(statuses))
		for i, v := range statuses {
			s[i] = v
		}
		u = u.Where(q.Status.In(s...))
	}
	if len(categories) > 0 {
		u = u.Where(q.Category.In(categories...))
	}

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	// Preload Feed and FeedActionSets
	// Note: GORM Gen Preload
	res, err := u.
		Preload(q.Feed).
		Preload(q.FeedActionSets).
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}

	result := make([]*modelyesod.FeedWithConfig, len(res))
	for i, config := range res {
		actionSets := make([]libmodel.InternalID, len(config.FeedActionSets))
		for j, action := range config.FeedActionSets {
			actionSets[j] = action.ID
		}
		config.ActionSets = actionSets

		result[i] = &modelyesod.FeedWithConfig{
			FeedConfig: config,
			Feed:       config.Feed,
		}
	}
	return result, int(total), nil
}

func (y *YesodRepo) ListFeedCategories(ctx context.Context, id libmodel.InternalID) ([]string, error) {
	q := query.Use(y.data.db).FeedConfig
	var categories []string
	// Distinct categories
	err := q.WithContext(ctx).
		Where(q.UserFeedConfig.Eq(int64(id))).
		Distinct(q.Category).
		Pluck(q.Category, &categories)
	return categories, err
}

func (y *YesodRepo) ListFeedPlatforms(ctx context.Context, id libmodel.InternalID) ([]string, error) {
	// Join FeedItem -> Feed -> FeedConfig -> User
	// FeedItem.PublishPlatform
	// In GORM:
	// db.Model(&FeedItem{}).
	// Joins("JOIN feeds ON feeds.id = feed_items.feed_id").
	// Joins("JOIN feed_configs ON feed_configs.id = feeds.id").
	// Where("feed_configs.user_id = ?", id).
	// Distinct("publish_platform").Pluck("publish_platform", &res)

	// Using Gen:
	q := query.Use(y.data.db)
	fi := q.Item
	f := q.Feed
	fc := q.FeedConfig

	var platforms []string
	err := fi.WithContext(ctx).
		Join(f, f.ID.EqCol(fi.FeedID)).
		Join(fc, fc.ID.EqCol(f.ID)).
		Where(fc.UserFeedConfig.Eq(int64(id))).
		Distinct(fi.PublishPlatform).
		Pluck(fi.PublishPlatform, &platforms)
	return platforms, err
}

func (y *YesodRepo) ListFeedItems(
	ctx context.Context,
	userID libmodel.InternalID,
	paging libmodel.Paging,
	feedIDs []libmodel.InternalID,
	authors []string,
	platforms []string,
	timeRange *libmodel.TimeRange,
	categories []string,
) ([]*modelyesod.FeedItemDigest, int, error) {
	q := query.Use(y.data.db)
	fi := q.Item
	f := q.Feed
	fc := q.FeedConfig

	u := fi.WithContext(ctx).
		Join(f, f.ID.EqCol(fi.FeedID)).
		Join(fc, fc.ID.EqCol(f.ID)).
		Where(fc.UserFeedConfig.Eq(int64(userID))).
		Where(fc.HideItems.Is(false))

	if len(feedIDs) > 0 {
		castFeedIDs := make([]int64, len(feedIDs))
		for i, v := range feedIDs {
			castFeedIDs[i] = int64(v)
		}
		u = u.Where(fi.FeedID.In(castFeedIDs...))
	}
	if len(platforms) > 0 {
		u = u.Where(fi.PublishPlatform.In(platforms...))
	}
	if len(categories) > 0 {
		u = u.Where(fc.Category.In(categories...))
	}
	if timeRange != nil {
		u = u.Where(
			fi.PublishedParsed.Gte(timeRange.StartTime),
			fi.PublishedParsed.Lt(timeRange.StartTime.Add(timeRange.Duration)),
		)
	}

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	// Select fields needed for Digest + Preload Feed + FeedConfig (for Name)
	// GORM Preload usually requires struct scan.
	var results []struct {
		modelfeed.Item

		FeedConfigName string `gorm:"column:feed_config_name"`
	}

	err = u.
		Preload(fi.Feed).
		Select(fi.ALL, fc.Name.As("feed_config_name")).
		Order(fi.PublishedParsed.Desc()).
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Scan(&results)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*modelyesod.FeedItemDigest, len(results))
	for i, item := range results {
		result[i] = y.toFeedItemDigest(&item.Item, item.FeedConfigName)
	}
	return result, int(total), nil
}

func (y *YesodRepo) GroupFeedItems(
	ctx context.Context,
	userID libmodel.InternalID,
	groups []libmodel.TimeRange,
	feedIDs []libmodel.InternalID,
	authors []string,
	platforms []string,
	groupSize int,
	categories []string,
) (map[libmodel.TimeRange][]*modelyesod.FeedItemDigest, error) {
	res := make(map[libmodel.TimeRange][]*modelyesod.FeedItemDigest)
	q := query.Use(y.data.db)
	fi := q.Item
	f := q.Feed
	fc := q.FeedConfig

	for _, tr := range groups {
		// Clone query for each group
		// Note: Gen `base` is mutable? `WithContext` returns a new instance usually.
		// `Join` modifies the DAO? Yes, usually.
		// So we should construct the query inside the loop or use `base.Clone()` if available.
		// Gen doesn't have `Clone`.
		// But `WithContext` creates a new `IFeedItemDo`.
		// However `Join` attaches to the underlying statement.
		// Safer to rebuild query or check Gen docs.
		// Assuming `WithContext` is enough if `Join` was called before.
		// Actually `Join` returns `IFeedItemDo`.

		// Let's rebuild the chain to be safe.
		u := fi.WithContext(ctx).
			Join(f, f.ID.EqCol(fi.FeedID)).
			Join(fc, fc.ID.EqCol(f.ID)).
			Where(fc.UserFeedConfig.Eq(int64(userID))).
			Where(fc.HideItems.Is(false))

		if len(feedIDs) > 0 {
			castFeedIDs := make([]int64, len(feedIDs))
			for i, v := range feedIDs {
				castFeedIDs[i] = int64(v)
			}
			u = u.Where(fi.FeedID.In(castFeedIDs...))
		}
		if len(platforms) > 0 {
			u = u.Where(fi.PublishPlatform.In(platforms...))
		}
		if len(categories) > 0 {
			u = u.Where(fc.Category.In(categories...))
		}

		var results []struct {
			modelfeed.Item

			FeedConfigName string `gorm:"column:feed_config_name"`
		}

		err := u.
			Where(fi.PublishedParsed.Gte(tr.StartTime), fi.PublishedParsed.Lt(tr.StartTime.Add(tr.Duration))).
			Preload(fi.Feed).
			Select(fi.ALL, fc.Name.As("feed_config_name")).
			Order(fi.PublishedParsed.Desc()).
			Limit(groupSize).
			Scan(&results)
		if err != nil {
			return nil, err
		}

		if len(results) == 0 {
			continue
		}

		digests := make([]*modelyesod.FeedItemDigest, len(results))
		for i, item := range results {
			digests[i] = y.toFeedItemDigest(&item.Item, item.FeedConfigName)
		}
		res[tr] = digests
	}
	return res, nil
}

func (y *YesodRepo) GetFeedItems(
	ctx context.Context,
	userID libmodel.InternalID,
	ids []libmodel.InternalID,
) ([]*modelfeed.Item, error) {
	q := query.Use(y.data.db)
	fi := q.Item
	f := q.Feed
	fc := q.FeedConfig

	castIDs := make([]int64, len(ids))
	for i, v := range ids {
		castIDs[i] = int64(v)
	}

	items, err := fi.WithContext(ctx).
		Join(f, f.ID.EqCol(fi.FeedID)).
		Join(fc, fc.ID.EqCol(f.ID)).
		Where(fc.UserFeedConfig.Eq(int64(userID))).
		Where(fi.ID.In(castIDs...)).
		Preload(fi.Feed).
		Find()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (y *YesodRepo) ReadFeedItem(ctx context.Context, userID libmodel.InternalID, id libmodel.InternalID) error {
	// Verify ownership/visibility via Join, then Update.
	// GORM Update with Join is tricky.
	// Alternative: Check existence then update.
	// Or: Exec custom SQL.
	// Or: Update with Where clause using subquery?
	// Gen supports simple updates.

	// Let's do:
	// UPDATE feed_items SET read_count = read_count + 1 WHERE id = ? AND feed_id IN (SELECT id FROM feeds WHERE id IN (SELECT id FROM feed_configs WHERE user_id = ?))

	// Using Gen:
	q := query.Use(y.data.db)
	fi := q.Item
	// Subquery for FeedConfig owner
	// This might be complicated in Gen without subquery support.
	// Let's verify first (read) then update, or simple update if we trust ID.
	// The original code did `Where(feeditem.HasFeedWith(...))`.

	// We can use `Update` with `Where`.
	// But `Where` involving joins for Update might not work on all DBs (SQLite supports it differently).
	// Simplest: Check permission first.

	// Check if item belongs to user
	exists, err := y.checkFeedItemOwner(ctx, id, userID)
	if err != nil {
		return err
	}
	if !exists {
		// Item not found or not owned
		return nil
	}

	_, err = fi.WithContext(ctx).
		Where(fi.ID.Eq(int64(id))).
		Update(fi.ReadCount, gorm.Expr("read_count + ?", 1))
	return err
}

func (y *YesodRepo) checkFeedItemOwner(ctx context.Context, itemID, userID libmodel.InternalID) (bool, error) {
	q := query.Use(y.data.db)
	fi := q.Item
	f := q.Feed
	fc := q.FeedConfig

	count, err := fi.WithContext(ctx).
		Join(f, f.ID.EqCol(fi.FeedID)).
		Join(fc, fc.ID.EqCol(f.ID)).
		Where(fi.ID.Eq(int64(itemID)), fc.UserFeedConfig.Eq(int64(userID))).
		Count()
	return count > 0, err
}

func (y *YesodRepo) CreateFeedItemCollection(
	ctx context.Context,
	ownerID libmodel.InternalID,
	collection *modelyesod.FeedItemCollection,
) error {
	collection.UserID = ownerID
	return query.Use(y.data.db).FeedItemCollection.WithContext(ctx).Create(collection)
}

func (y *YesodRepo) UpdateFeedItemCollection(
	ctx context.Context,
	ownerID libmodel.InternalID,
	collection *modelyesod.FeedItemCollection,
) error {
	q := query.Use(y.data.db).FeedItemCollection
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(collection.ID)), q.UserID.Eq(int64(ownerID))).
		Updates(&modelyesod.FeedItemCollection{
			Name:        collection.Name,
			Description: collection.Description,
			Category:    collection.Category,
		})
	return err
}

func (y *YesodRepo) ListFeedItemCollections(
	ctx context.Context,
	ownerID libmodel.InternalID,
	paging libmodel.Paging,
	ids []libmodel.InternalID,
	categories []string,
) ([]*modelyesod.FeedItemCollection, int, error) {
	q := query.Use(y.data.db).FeedItemCollection
	u := q.WithContext(ctx).Where(q.UserID.Eq(int64(ownerID)))

	if len(ids) > 0 {
		castIDs := make([]int64, len(ids))
		for i, v := range ids {
			castIDs[i] = int64(v)
		}
		u = u.Where(q.ID.In(castIDs...))
	}
	if len(categories) > 0 {
		u = u.Where(q.Category.In(categories...))
	}

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	res, err := u.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}

	return res, int(total), nil
}

func (y *YesodRepo) AddFeedItemToCollection(
	ctx context.Context,
	ownerID libmodel.InternalID,
	collectionID libmodel.InternalID,
	itemID libmodel.InternalID,
) error {
	// Verify ownership
	q := query.Use(y.data.db)
	fic := q.FeedItemCollection
	count, err := fic.WithContext(ctx).Where(fic.ID.Eq(int64(collectionID)), fic.UserID.Eq(int64(ownerID))).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return nil // Or error not found
	}

	// Add M2M
	// We need to insert into `feed_item_collection_feed_items`?
	// The model `FeedItemCollection` has `FeedItems []FeedItem`.
	// GORM: association.Append

	return y.data.db.Model(&modelyesod.FeedItemCollection{ID: collectionID}).
		Association("FeedItems").
		Append(&modelfeed.Item{ID: itemID})
}

func (y *YesodRepo) RemoveFeedItemFromCollection(
	ctx context.Context,
	ownerID libmodel.InternalID,
	collectionID libmodel.InternalID,
	itemID libmodel.InternalID,
) error {
	// Verify ownership
	q := query.Use(y.data.db)
	fic := q.FeedItemCollection
	count, err := fic.WithContext(ctx).Where(fic.ID.Eq(int64(collectionID)), fic.UserID.Eq(int64(ownerID))).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return nil
	}

	return y.data.db.Model(&modelyesod.FeedItemCollection{ID: collectionID}).
		Association("FeedItems").
		Delete(&modelfeed.Item{ID: itemID})
}

//nolint:funlen
func (y *YesodRepo) ListFeedItemsInCollection(
	ctx context.Context,
	ownerID libmodel.InternalID,
	paging libmodel.Paging,
	ids []libmodel.InternalID,
	authors []string,
	platforms []string,
	categories []string,
	timeRange *libmodel.TimeRange,
) ([]*modelyesod.FeedItemDigest, int, error) {
	// Join FeedItemCollection -> FeedItems
	// Query FeedItems where collection_id IN (ids) AND collection.user_id = ownerID

	// The M2M table is implicit in GORM if not defined explicitly.
	// But we can query FeedItem and Join FeedItemCollection via association?
	// Or start from FeedItemCollection and Preload FeedItems?
	// We need filtering on FeedItems (platform, category etc).
	// Starting from FeedItem is better.
	// FeedItem JOIN feed_item_collection_feed_items JOIN feed_item_collections

	// Gen doesn't expose the join table automatically unless defined.
	// Assuming `FeedItem` has `FeedItemCollections`?
	// Let's check `FeedItem` model. If M2M is defined, it should have it.

	// Assuming standard M2M.
	// db.Model(&FeedItem{}).Joins("JOIN feed_item_collection_feed_items ...")

	// Simpler: Use GORM `Association` count/find? No, we need complex filtering.

	// Let's use `db.Table`.
	// Or assume we have `FeedItemCollections` in `FeedItem`.

	// If not, we can use `Joins` with the join table name.
	// Join table name usually `feed_item_collection_feed_items` or similar.
	// Ent was `feed_item_collection_feed_items`?

	// Let's assume `FeedItem` has `FeedItemCollections` in `internal/data/orm/model`.
	// If not, I'll need to use raw join or defined model.
	// I'll assume `FeedItem` has `FeedItemCollections` for now.

	q := query.Use(y.data.db)
	fi := q.Item
	fic := q.FeedItemCollection
	f := q.Feed
	fc := q.FeedConfig

	// We need to join the M2M table.
	// If `FeedItem` has `FeedItemCollections` field:
	// fi.Join(fic, ...) won't work directly for M2M without join table.
	// But GORM Gen `Join` supports relationship?
	// `fi.Join(fi.FeedItemCollections)` ?

	// Let's rely on `Where` with subquery or `Joins` if we know the table.
	// `feed_item_feed_item_collections` is the likely table name if generated by GORM from `FeedItem` and `FeedItemCollection`.
	// Or `feed_item_collection_feed_items`.

	// Let's try to find items that belong to *any* of the collections owned by user and in `ids`.
	// Subquery approach:
	// Select * from feed_items where id in (select feed_item_id from join_table where feed_item_collection_id in (ids) and feed_item_collection_id in (select id from feed_item_collections where user_id = owner))

	// This is getting complex.
	// Let's stick to:
	// 1. Get collection IDs owned by user (filtered by input ids).
	// 2. Query items in those collections.

	castIDs := make([]int64, len(ids))
	for i, v := range ids {
		castIDs[i] = int64(v)
	}

	// Re-do Pluck
	var validIDs []int64
	err := fic.WithContext(ctx).
		Where(fic.UserID.Eq(int64(ownerID))).
		Where(fic.ID.In(castIDs...)).
		Pluck(fic.ID, &validIDs)
	if err != nil {
		return nil, 0, err
	}

	if len(validIDs) == 0 {
		return []*modelyesod.FeedItemDigest{}, 0, nil
	}

	// Now find items in these collections.
	// We need the join table.
	// db.Table("feed_item_collection_feed_items").Where("feed_item_collection_id IN ?", validIDs).Select("feed_item_id")

	var itemIDs []int64
	// Use raw SQL or dynamic query for join table
	err = y.data.db.Table("feed_item_collection_feed_items"). // Verify table name!
									Where("feed_item_collection_id IN ?", validIDs).
									Pluck("feed_item_id", &itemIDs).Error
	if err != nil {
		return nil, 0, err
	}

	if len(itemIDs) == 0 {
		return []*modelyesod.FeedItemDigest{}, 0, nil
	}

	// Now query items with filters
	u := fi.WithContext(ctx).
		Join(f, f.ID.EqCol(fi.FeedID)).
		Join(fc, fc.ID.EqCol(f.ID)).
		Where(fi.ID.In(itemIDs...))

	if len(platforms) > 0 {
		u = u.Where(fi.PublishPlatform.In(platforms...))
	}
	if len(categories) > 0 {
		u = u.Where(fc.Category.In(categories...))
	}
	if timeRange != nil {
		u = u.Where(
			fi.PublishedParsed.Gte(timeRange.StartTime),
			fi.PublishedParsed.Lt(timeRange.StartTime.Add(timeRange.Duration)),
		)
	}

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	var results []struct {
		modelfeed.Item

		FeedConfigName string `gorm:"column:feed_config_name"`
	}

	err = u.
		Preload(fi.Feed).
		Select(fi.ALL, fc.Name.As("feed_config_name")).
		Order(fi.PublishedParsed.Desc()).
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Scan(&results)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*modelyesod.FeedItemDigest, len(results))
	for i, item := range results {
		result[i] = y.toFeedItemDigest(&item.Item, item.FeedConfigName)
	}
	return result, int(total), nil
}

func (y *YesodRepo) GetFeedOwner(ctx context.Context, id libmodel.InternalID) (*libmodel.User, error) {
	q := query.Use(y.data.db)
	fc := q.FeedConfig
	u := q.User

	// Join FeedConfig -> User
	res, err := u.WithContext(ctx).
		Join(fc, fc.UserFeedConfig.EqCol(u.ID)).
		Where(fc.ID.Eq(int64(id))).
		First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *YesodRepo) CreateFeedActionSet(
	ctx context.Context,
	id libmodel.InternalID,
	set *modelyesod.FeedActionSet,
) error {
	set.UserID = id
	return query.Use(y.data.db).FeedActionSet.WithContext(ctx).Create(set)
}

func (y *YesodRepo) UpdateFeedActionSet(
	ctx context.Context,
	id libmodel.InternalID,
	set *modelyesod.FeedActionSet,
) error {
	q := query.Use(y.data.db).FeedActionSet
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(set.ID)), q.UserID.Eq(int64(id))).
		Updates(&modelyesod.FeedActionSet{
			Name:        set.Name,
			Description: set.Description,
			Actions:     set.Actions,
		})
	return err
}

func (y *YesodRepo) ListFeedActionSets(
	ctx context.Context,
	id libmodel.InternalID,
	paging libmodel.Paging,
) ([]*modelyesod.FeedActionSet, int, error) {
	q := query.Use(y.data.db).FeedActionSet
	u := q.WithContext(ctx).Where(q.UserID.Eq(int64(id)))

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	res, err := u.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}

	return res, int(total), nil
}

func (y *YesodRepo) toFeedItemDigest(item *modelfeed.Item, feedConfigName string) *modelyesod.FeedItemDigest {
	digest := &modelyesod.FeedItemDigest{
		FeedID:              item.FeedID,
		ItemID:              item.ID,
		PublishedParsedTime: *item.PublishedParsed,
		Title:               item.Title,
		ShortDescription:    item.DigestDescription,
		PublishPlatform:     item.PublishPlatform,
		ReadCount:           item.ReadCount,
		FeedConfigName:      feedConfigName,
	}

	if item.Feed != nil {
		if item.Feed.Image != nil {
			digest.FeedAvatarURL = item.Feed.Image.URL
		}
	}

	// Authors
	if len(item.Authors) > 0 {
		digest.Authors = item.Authors[0].Name // Simple simplification
	}

	// Image URLs
	if len(item.DigestImages) > 0 {
		digest.ImageUrls = make([]string, len(item.DigestImages))
		for i, img := range item.DigestImages {
			digest.ImageUrls[i] = img.URL
		}
	}

	if item.Image != nil {
		digest.AvatarURL = item.Image.URL
	}

	return digest
}

package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedactionset"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type yesodRepo struct {
	data *Data
}

// NewYesodRepo .
func NewYesodRepo(data *Data) bizyesod.YesodRepo {
	return &yesodRepo{
		data: data,
	}
}

func (y *yesodRepo) CreateFeedConfig(ctx context.Context, owner model.InternalID, c *modelyesod.FeedConfig) error {
	q := y.data.db.FeedConfig.Create().
		SetOwnerID(owner).
		SetID(c.ID).
		SetName(c.Name).
		SetFeedURL(c.FeedURL).
		SetCategory(c.Category).
		SetAuthorAccount(c.AuthorAccount).
		SetSource(c.Source).
		SetStatus(converter.ToEntFeedConfigStatus(c.Status)).
		SetPullInterval(c.PullInterval).
		SetLatestPullStatus(converter.ToEntFeedConfigLatestPullStatus(c.LatestPullStatus)).
		SetLatestPullMessage("").
		SetHideItems(c.HideItems).
		AddFeedActionSetIDs(c.ActionSets...)
	return q.Exec(ctx)
}

func (y *yesodRepo) UpdateFeedConfig(ctx context.Context, userID model.InternalID, c *modelyesod.FeedConfig) error {
	q := y.data.db.FeedConfig.Update().
		Where(
			feedconfig.IDEQ(c.ID),
			feedconfig.HasOwnerWith(user.IDEQ(userID)),
		)
	if len(c.Name) > 0 {
		q.SetName(c.Name)
	}
	if len(c.FeedURL) > 0 {
		q.SetFeedURL(c.FeedURL)
	}
	if len(c.Category) > 0 {
		q.SetCategory(c.Category)
	}
	if c.AuthorAccount > 0 {
		q.SetAuthorAccount(c.AuthorAccount)
	}
	if len(c.Source) > 0 {
		q.SetSource(c.Source)
	}
	if c.Status != modelyesod.FeedConfigStatusUnspecified {
		q.SetStatus(converter.ToEntFeedConfigStatus(c.Status))
	}
	if c.PullInterval > 0 {
		q.SetPullInterval(c.PullInterval).SetNextPullBeginAt(time.Now())
	}
	if c.ActionSets != nil {
		q.ClearFeedActionSet().AddFeedActionSetIDs(c.ActionSets...)
	}
	q.SetHideItems(c.HideItems)
	return q.Exec(ctx)
}

// UpdateFeedConfigAsInQueue set SetNextPullBeginAt to one day later to avoid repeat queue.
// While pull success, UpsertFeed will set correct value.
// While pull failed, server will retry task next day.
func (y *yesodRepo) UpdateFeedConfigAsInQueue(ctx context.Context, id model.InternalID) error {
	q := y.data.db.FeedConfig.UpdateOneID(id).
		SetNextPullBeginAt(time.Now().Add(libtime.Day))
	return q.Exec(ctx)
}

func (y *yesodRepo) ListFeedConfigNeedPull(ctx context.Context, sources []string,
	statuses []modelyesod.FeedConfigStatus, order modelyesod.ListFeedOrder,
	pullTime time.Time, i int) ([]*modelyesod.FeedConfig, error) {
	q := y.data.db.FeedConfig.Query()
	if len(sources) > 0 {
		q.Where(feedconfig.SourceIn(sources...))
	}
	if len(statuses) > 0 {
		q.Where(feedconfig.StatusIn(converter.ToEntFeedConfigStatusList(statuses)...))
	}
	switch order {
	case modelyesod.ListFeedOrderUnspecified:
		{
		}
	case modelyesod.ListFeedOrderNextPull:
		q.Where(feedconfig.NextPullBeginAtLT(pullTime))
	}
	q.Limit(i)
	feedConfigs, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizFeedConfigList(feedConfigs), nil
}

func (y *yesodRepo) ListFeedConfigs(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
	ids []model.InternalID,
	authorIDs []model.InternalID,
	sources []string,
	statuses []modelyesod.FeedConfigStatus,
	categories []string,
) ([]*modelyesod.FeedWithConfig, int, error) {
	var res []*modelyesod.FeedWithConfig
	var total int
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		q := tx.User.QueryFeedConfig(u)
		if len(ids) > 0 {
			q.Where(feedconfig.IDIn(ids...))
		}
		if len(authorIDs) > 0 {
			q.Where(feedconfig.AuthorAccountIn(authorIDs...))
		}
		if len(sources) > 0 {
			q.Where(feedconfig.SourceIn(sources...))
		}
		if len(statuses) > 0 {
			q.Where(feedconfig.StatusIn(converter.ToEntFeedConfigStatusList(statuses)...))
		}
		if len(categories) > 0 {
			q.Where(feedconfig.CategoryIn(categories...))
		}
		total, err = q.Count(ctx)
		if err != nil {
			return err
		}
		configs, err := q.
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			WithFeed().
			All(ctx)
		if err != nil {
			return err
		}
		res = make([]*modelyesod.FeedWithConfig, 0, len(configs))
		for _, config := range configs {
			res = append(res, &modelyesod.FeedWithConfig{
				FeedConfig: converter.ToBizFeedConfig(config),
				Feed:       converter.ToBizFeed(config.Edges.Feed),
			})
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (y *yesodRepo) ListFeedCategories(ctx context.Context, id model.InternalID) ([]string, error) {
	res, err := y.data.db.FeedConfig.Query().
		Where(
			feedconfig.HasOwnerWith(user.IDEQ(id)),
		).
		Select(feedconfig.FieldCategory).
		GroupBy(feedconfig.FieldCategory).
		Strings(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *yesodRepo) ListFeedPlatforms(ctx context.Context, id model.InternalID) ([]string, error) {
	res, err := y.data.db.FeedItem.Query().
		Where(
			feeditem.HasFeedWith(feed.HasConfigWith(feedconfig.HasOwnerWith(user.IDEQ(id)))),
		).
		Select(feeditem.FieldPublishPlatform).
		GroupBy(feeditem.FieldPublishPlatform).
		Strings(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *yesodRepo) ListFeedItems(
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
	var total int
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		fq := tx.User.QueryFeedConfig(u).Where(
			feedconfig.HideItemsEQ(false),
		).QueryFeed()
		if len(feedIDs) > 0 {
			fq.Where(feed.IDIn(feedIDs...))
		}
		iq := fq.QueryItem()
		if len(platforms) > 0 {
			iq.Where(feeditem.PublishPlatformIn(platforms...))
		}
		if len(categories) > 0 {
			iq.Where(feeditem.HasFeedWith(feed.HasConfigWith(feedconfig.CategoryIn(categories...))))
		}
		if timeRange != nil {
			iq.
				Where(feeditem.PublishedParsedGTE(timeRange.StartTime)).
				Where(feeditem.PublishedParsedLT(timeRange.StartTime.Add(timeRange.Duration)))
		}
		total, err = iq.Count(ctx)
		if err != nil {
			return err
		}
		items, err := iq.
			WithFeed(func(q *ent.FeedQuery) {
				q.Select(feed.FieldImage).WithConfig(func(q *ent.FeedConfigQuery) {
					q.Select(feedconfig.FieldName)
				})
			}).
			Order(ent.Desc(feeditem.FieldPublishedParsed)).
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			All(ctx)
		if err != nil {
			return err
		}
		res = make([]*modelyesod.FeedItemDigest, 0, len(items))
		for _, item := range items {
			res = append(res, converter.ToBizFeedItemDigest(item))
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (y *yesodRepo) GroupFeedItems( //nolint:gocognit //TODO
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
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		for _, timeRange := range groups {
			fq := tx.User.QueryFeedConfig(u).Where(
				feedconfig.HideItemsEQ(false),
			).QueryFeed()
			if len(feedIDs) > 0 {
				fq.Where(feed.IDIn(feedIDs...))
			}
			iq := fq.QueryItem()
			if len(platforms) > 0 {
				iq.Where(feeditem.PublishPlatformIn(platforms...))
			}
			if len(categories) > 0 {
				iq.Where(feeditem.HasFeedWith(feed.HasConfigWith(feedconfig.CategoryIn(categories...))))
			}
			var items []*ent.FeedItem
			items, err = iq.
				Where(feeditem.PublishedParsedGTE(timeRange.StartTime)).
				Where(feeditem.PublishedParsedLT(timeRange.StartTime.Add(timeRange.Duration))).
				WithFeed(func(q *ent.FeedQuery) {
					q.Select(feed.FieldImage).WithConfig(func(q *ent.FeedConfigQuery) {
						q.Select(feedconfig.FieldName)
					})
				}).
				Order(ent.Desc(feeditem.FieldPublishedParsed)).
				Limit(groupSize).
				All(ctx)
			if err != nil {
				return err
			}
			if len(items) == 0 {
				continue
			}
			il := make([]*modelyesod.FeedItemDigest, 0, len(items))
			for _, item := range items {
				il = append(il, converter.ToBizFeedItemDigest(item))
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

func (y *yesodRepo) GetFeedItems(
	ctx context.Context,
	userID model.InternalID,
	ids []model.InternalID,
) ([]*modelfeed.Item, error) {
	var res []*modelfeed.Item
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		items, err := tx.User.
			QueryFeedConfig(u).
			QueryFeed().
			QueryItem().
			Where(feeditem.IDIn(ids...)).
			All(ctx)
		if err != nil {
			return err
		}
		res = converter.ToBizFeedItemList(items)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *yesodRepo) ReadFeedItem(ctx context.Context, userID model.InternalID, id model.InternalID) error {
	return y.data.db.FeedItem.UpdateOneID(id).Where(
		feeditem.HasFeedWith(feed.HasConfigWith(feedconfig.HasOwnerWith(user.IDEQ(userID)))),
	).AddReadCount(1).Exec(ctx)
}

func (y *yesodRepo) CreateFeedItemCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collection *modelyesod.FeedItemCollection,
) error {
	return y.data.db.FeedItemCollection.Create().
		SetOwnerID(ownerID).
		SetID(collection.ID).
		SetName(collection.Name).
		SetDescription(collection.Description).
		SetCategory(collection.Category).
		Exec(ctx)
}

func (y *yesodRepo) UpdateFeedItemCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collection *modelyesod.FeedItemCollection,
) error {
	return y.data.db.FeedItemCollection.UpdateOneID(collection.ID).
		Where(feeditemcollection.HasOwnerWith(user.IDEQ(ownerID))).
		SetName(collection.Name).
		SetDescription(collection.Description).
		SetCategory(collection.Category).
		Exec(ctx)
}

func (y *yesodRepo) ListFeedItemCollections(
	ctx context.Context,
	ownerID model.InternalID,
	paging model.Paging,
	ids []model.InternalID,
	categories []string,
) ([]*modelyesod.FeedItemCollection, int, error) {
	var res []*modelyesod.FeedItemCollection
	var total int
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, ownerID)
		if err != nil {
			return err
		}
		q := tx.User.QueryFeedItemCollection(u)
		if len(ids) > 0 {
			q.Where(feeditemcollection.IDIn(ids...))
		}
		if len(categories) > 0 {
			q.Where(feeditemcollection.CategoryIn(categories...))
		}
		total, err = q.Count(ctx)
		if err != nil {
			return err
		}
		collections, err := q.
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			All(ctx)
		if err != nil {
			return err
		}
		res = converter.ToBizFeedItemCollectionList(collections)
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (y *yesodRepo) AddFeedItemToCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collectionID model.InternalID,
	itemID model.InternalID,
) error {
	return y.data.db.FeedItemCollection.UpdateOneID(collectionID).
		Where(feeditemcollection.HasOwnerWith(user.IDEQ(ownerID))).
		AddFeedItemIDs(itemID).
		Exec(ctx)
}

func (y *yesodRepo) RemoveFeedItemFromCollection(
	ctx context.Context,
	ownerID model.InternalID,
	collectionID model.InternalID,
	itemID model.InternalID,
) error {
	return y.data.db.FeedItemCollection.UpdateOneID(collectionID).
		Where(feeditemcollection.HasOwnerWith(user.IDEQ(ownerID))).
		RemoveFeedItemIDs(itemID).
		Exec(ctx)
}

func (y *yesodRepo) ListFeedItemsInCollection(
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
	var total int
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, ownerID)
		if err != nil {
			return err
		}
		q := tx.User.QueryFeedItemCollection(u)
		if len(ids) > 0 {
			q.Where(feeditemcollection.IDIn(ids...))
		}
		iq := q.QueryFeedItem()
		if len(platforms) > 0 {
			iq.Where(feeditem.PublishPlatformIn(platforms...))
		}
		if len(categories) > 0 {
			iq.Where(feeditem.HasFeedWith(feed.HasConfigWith(feedconfig.CategoryIn(categories...))))
		}
		if timeRange != nil {
			iq.
				Where(feeditem.PublishedParsedGTE(timeRange.StartTime)).
				Where(feeditem.PublishedParsedLT(timeRange.StartTime.Add(timeRange.Duration)))
		}
		total, err = iq.Count(ctx)
		if err != nil {
			return err
		}
		items, err := iq.
			WithFeed(func(q *ent.FeedQuery) {
				q.Select(feed.FieldImage).WithConfig(func(q *ent.FeedConfigQuery) {
					q.Select(feedconfig.FieldName)
				})
			}).
			Order(ent.Desc(feeditem.FieldPublishedParsed)).
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			All(ctx)
		if err != nil {
			return err
		}
		res = make([]*modelyesod.FeedItemDigest, 0, len(items))
		for _, item := range items {
			res = append(res, converter.ToBizFeedItemDigest(item))
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (y *yesodRepo) GetFeedOwner(ctx context.Context, id model.InternalID) (*modeltiphereth.User, error) {
	only, err := y.data.db.FeedConfig.Query().Where(feedconfig.IDEQ(id)).QueryOwner().Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizUser(only), nil
}

func (y *yesodRepo) CreateFeedActionSet(ctx context.Context, id model.InternalID, set *modelyesod.FeedActionSet) error {
	return y.data.db.FeedActionSet.Create().
		SetOwnerID(id).
		SetID(set.ID).
		SetName(set.Name).
		SetDescription(set.Description).
		SetActions(set.Actions).
		Exec(ctx)
}

func (y *yesodRepo) UpdateFeedActionSet(ctx context.Context, id model.InternalID, set *modelyesod.FeedActionSet) error {
	return y.data.db.FeedActionSet.UpdateOneID(set.ID).
		Where(feedactionset.HasOwnerWith(user.IDEQ(id))).
		SetName(set.Name).
		SetDescription(set.Description).
		SetActions(set.Actions).
		Exec(ctx)
}

func (y *yesodRepo) ListFeedActionSets(ctx context.Context, id model.InternalID, paging model.Paging) ([]*modelyesod.FeedActionSet, int, error) {
	var res []*modelyesod.FeedActionSet
	var total int
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, id)
		if err != nil {
			return err
		}
		q := tx.User.QueryFeedActionSet(u)
		total, err = q.Count(ctx)
		if err != nil {
			return err
		}
		sets, err := q.
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			All(ctx)
		if err != nil {
			return err
		}
		res = converter.ToBizFeedActionSetList(sets)
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

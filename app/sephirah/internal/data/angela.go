package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"entgo.io/ent/dialect/sql"
)

type angelaRepo struct {
	data *Data
}

// NewAngelaRepo .
func NewAngelaRepo(data *Data) bizangela.AngelaRepo {
	return &angelaRepo{
		data: data,
	}
}

func (a *angelaRepo) UpdateAccount(ctx context.Context, acc modeltiphereth.Account) error {
	return a.data.db.Account.Update().Where(
		account.IDEQ(acc.ID),
		account.PlatformEQ(acc.Platform),
		account.PlatformAccountIDEQ(acc.PlatformAccountID),
	).
		SetName(acc.Name).
		SetProfileURL(acc.ProfileURL).
		SetAvatarURL(acc.AvatarURL).
		Exec(ctx)
}

func (a *angelaRepo) UpdateApp( //nolint:gocognit //TODO
	ctx context.Context, ap *modelgebura.App, internal *modelgebura.App,
) error {
	return a.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.App.Update().
			Where(
				app.IDEQ(ap.ID),
				app.SourceEQ(ap.Source),
				app.SourceAppIDEQ(ap.SourceAppID),
			)
		if len(ap.SourceURL) > 0 {
			q.SetSourceURL(ap.SourceURL)
		}
		if len(ap.Name) > 0 {
			q.SetName(ap.Name)
		}
		if ap.Type != modelgebura.AppTypeUnspecified {
			q.SetType(converter.ToEntAppType(ap.Type))
		}
		if len(ap.ShortDescription) > 0 {
			q.SetShortDescription(ap.ShortDescription)
		}
		if len(ap.IconImageURL) > 0 {
			q.SetIconImageURL(ap.IconImageURL)
		}
		if len(ap.HeroImageURL) > 0 {
			q.SetHeroImageURL(ap.HeroImageURL)
		}
		if ap.Details != nil { //nolint:nestif // TODO
			if len(ap.Details.Description) > 0 {
				q.SetDescription(ap.Details.Description)
			}
			if len(ap.Details.ReleaseDate) > 0 {
				q.SetReleaseDate(ap.Details.ReleaseDate)
			}
			if len(ap.Details.Developer) > 0 {
				q.SetDeveloper(ap.Details.Developer)
			}
			if len(ap.Details.Publisher) > 0 {
				q.SetPublisher(ap.Details.Publisher)
			}
			if len(ap.Details.Version) > 0 {
				q.SetVersion(ap.Details.Version)
			}
		}
		count, err := tx.App.Query().Where(
			app.SourceEQ(ap.Source),
			app.SourceAppIDEQ(ap.SourceAppID),
			app.HasBindInternalWith(app.IDNEQ(0)),
		).Count(ctx)
		if err != nil {
			return err
		}
		if count == 0 {
			err = tx.App.Create().
				SetID(internal.ID).
				SetSource(internal.Source).
				SetSourceAppID(internal.SourceAppID).
				SetName(internal.Name).
				SetType(converter.ToEntAppType(internal.Type)).
				SetBindInternalID(internal.BoundInternal).
				Exec(ctx)
			if err != nil {
				return err
			}
			q.SetBindInternalID(internal.ID)
		}
		return q.Exec(ctx)
	})
}

func (a *angelaRepo) UpsertApps(ctx context.Context, al []*modelgebura.App) error {
	apps := make([]*ent.AppCreate, len(al))
	for i, ap := range al {
		if ap.Details == nil {
			ap.Details = new(modelgebura.AppDetails)
		}
		apps[i] = a.data.db.App.Create().
			SetID(ap.ID).
			SetInternal(ap.Internal).
			SetSource(ap.Source).
			SetSourceAppID(ap.SourceAppID).
			SetSourceURL(ap.SourceURL).
			SetName(ap.Name).
			SetType(converter.ToEntAppType(ap.Type)).
			SetShortDescription(ap.ShortDescription).
			SetIconImageURL(ap.IconImageURL).
			SetHeroImageURL(ap.HeroImageURL)
		if ap.Details != nil {
			apps[i].
				SetDescription(ap.Details.Description).
				SetReleaseDate(ap.Details.ReleaseDate).
				SetDeveloper(ap.Details.Developer).
				SetPublisher(ap.Details.Publisher).
				SetVersion(ap.Details.Version)
		}
	}
	return a.data.db.App.
		CreateBulk(apps...).
		OnConflict(
			sql.ConflictColumns(app.FieldSource, app.FieldSourceAppID),
			resolveWithIgnores([]string{
				app.FieldID,
			}),
		).
		Exec(ctx)
}

func (a *angelaRepo) AccountPurchaseApps(
	ctx context.Context, id model.InternalID, ids []model.InternalID,
) error {
	return a.data.WithTx(ctx, func(tx *ent.Tx) error {
		appIDs, err := tx.App.Query().Where(
			app.IDIn(ids...),
		).
			IDs(ctx)
		if err != nil {
			return err
		}
		return a.data.db.Account.
			UpdateOneID(id).
			AddPurchasedAppIDs(appIDs...).
			Exec(ctx)
	})
}

func (a *angelaRepo) UpsertFeed(ctx context.Context, f *modelfeed.Feed) error {
	return a.data.WithTx(ctx, func(tx *ent.Tx) error {
		conf, err := tx.FeedConfig.Query().
			Where(feedconfig.IDEQ(f.ID)).
			Only(ctx)
		if err != nil {
			return err
		}
		err = tx.Feed.Create().
			SetConfig(conf).
			SetID(f.ID).
			SetTitle(f.Title).
			SetDescription(f.Description).
			SetLink(f.Link).
			SetAuthors(f.Authors).
			SetLanguage(f.Language).
			SetImage(f.Image).
			OnConflict(
				sql.ConflictColumns(feed.FieldID),
				sql.ResolveWithNewValues(),
			).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.FeedConfig.Update().
			Where(feedconfig.IDEQ(f.ID)).
			SetLatestPullAt(time.Now()).
			SetNextPullBeginAt(time.Now().Add(conf.PullInterval)).
			Exec(ctx)
		return err
	})
}

func (a *angelaRepo) UpsertFeedItems(
	ctx context.Context,
	items []*modelfeed.Item,
	feedID model.InternalID,
) ([]string, error) {
	guids := make([]string, 0, len(items))
	for _, item := range items {
		guids = append(guids, item.GUID)
	}
	existItems, err := a.data.db.FeedItem.Query().Where(
		feeditem.FeedID(feedID),
		feeditem.GUIDIn(guids...),
	).Select(feeditem.FieldGUID).All(ctx)
	if err != nil {
		return nil, err
	}
	il := make([]*ent.FeedItemCreate, len(items))
	for i, item := range items {
		il[i] = a.data.db.FeedItem.Create().
			SetFeedID(feedID).
			SetID(item.ID).
			SetTitle(item.Title).
			SetDescription(item.Description).
			SetContent(item.Content).
			SetLink(item.Link).
			SetUpdated(item.Updated).
			SetNillableUpdatedParsed(item.UpdatedParsed).
			SetPublished(item.Published).
			SetAuthors(item.Authors).
			SetGUID(item.GUID).
			SetImage(item.Image).
			SetEnclosures(item.Enclosures).
			SetPublishPlatform(item.PublishPlatform)
		if item.PublishedParsed != nil {
			il[i].SetPublishedParsed(*item.PublishedParsed)
		} else {
			il[i].SetPublishedParsed(time.Now())
		}
	}
	err = a.data.db.FeedItem.CreateBulk(il...).
		OnConflict(
			sql.ConflictColumns(feeditem.FieldFeedID, feeditem.FieldGUID),
			//
			// Update feed item every time result in large disk writes
			//
			// resolveWithIgnores([]string{
			//	feeditem.FieldID,
			// }),
			sql.DoNothing(),
		).Exec(ctx)
	if err != nil {
		return nil, err
	}
	existItemMap := make(map[string]bool)
	res := make([]string, 0, len(items)-len(existItems))
	for _, item := range existItems {
		existItemMap[item.GUID] = true
	}
	for _, item := range items {
		if _, exist := existItemMap[item.GUID]; !exist {
			res = append(res, item.GUID)
		}
	}
	return res, nil
}

func (a *angelaRepo) GetFeedItem(ctx context.Context, id model.InternalID) (*modelfeed.Item, error) {
	item, err := a.data.db.FeedItem.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToBizFeedItem(item), nil
}

func (a *angelaRepo) UpdateFeedItemDigest(ctx context.Context, item *modelfeed.Item) error {
	err := a.data.db.FeedItem.UpdateOneID(item.ID).
		SetDigestDescription(item.DigestDescription).
		SetDigestImages(item.DigestImages).
		Exec(ctx)
	return err
}

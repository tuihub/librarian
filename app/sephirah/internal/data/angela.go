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
		account.PlatformEQ(converter.ToEntAccountPlatform(acc.Platform)),
		account.PlatformAccountIDEQ(acc.PlatformAccountID),
	).
		SetName(acc.Name).
		SetProfileURL(acc.ProfileURL).
		SetAvatarURL(acc.AvatarURL).
		Exec(ctx)
}

func (a *angelaRepo) UpsertApps(ctx context.Context, al []*modelgebura.App) error {
	apps := make([]*ent.AppCreate, len(al))
	for i, ap := range al {
		if ap.Details == nil {
			ap.Details = new(modelgebura.AppDetails)
		}
		apps[i] = a.data.db.App.Create().
			SetID(ap.ID).
			SetSource(converter.ToEntAppSource(ap.Source)).
			SetSourceAppID(ap.SourceAppID).
			SetSourceURL(ap.SourceURL).
			SetName(ap.Name).
			SetType(converter.ToEntAppType(ap.Type)).
			SetShortDescription(ap.ShortDescription).
			SetImageURL(ap.ImageURL).
			SetBindInternalID(ap.BoundInternal)
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
			resolveWithIgnores([]string{
				feeditem.FieldID,
			}),
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

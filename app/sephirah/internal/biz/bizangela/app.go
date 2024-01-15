package bizangela

import (
	"context"
	"strconv"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPullAppTopic(
	a *AngelaBase,
	appCache *libcache.Map[modelgebura.AppID, modelgebura.App],
	updateAppIndex *libmq.Topic[modelangela.UpdateAppIndex],
) *libmq.Topic[modelangela.PullApp] {
	return libmq.NewTopic[modelangela.PullApp](
		"PullApp",
		func(ctx context.Context, r *modelangela.PullApp) error {
			if !a.supv.CheckAppSource(r.AppID.Source) {
				return nil
			}
			if app, err := appCache.GetWithFallBack(ctx, r.AppID, nil); err == nil &&
				app.LatestUpdateTime.Add(libtime.Day).After(time.Now()) {
				return nil
			}
			id, err := a.searcher.NewID(ctx)
			if err != nil {
				return err
			}
			resp, err := a.porter.PullApp(
				a.supv.CallAppSource(ctx, r.AppID.Source),
				&porter.PullAppRequest{AppId: &librarian.AppID{
					Internal:    false,
					Source:      r.AppID.Source,
					SourceAppId: r.AppID.SourceAppID,
				}},
			)
			if err != nil {
				return err
			}
			app := converter.ToBizApp(resp.GetApp())
			app.ID = r.ID
			app.Internal = false
			app.Source = r.AppID.Source
			internalApp := new(modelgebura.App)
			internalApp.ID = id
			internalApp.Internal = true
			internalApp.SourceAppID = strconv.FormatInt(int64(internalApp.ID), 10)
			internalApp.BoundInternal = id
			err = a.repo.UpsertApp(ctx, app, internalApp)
			if err != nil {
				return err
			}
			_ = appCache.Delete(ctx, r.AppID)
			_ = updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{id}})
			return nil
		},
	)
}

func NewAppCache(
	g bizgebura.GeburaRepo,
	store libcache.Store,
) *libcache.Map[modelgebura.AppID, modelgebura.App] {
	return libcache.NewMap[modelgebura.AppID, modelgebura.App](
		store,
		"App",
		func(k modelgebura.AppID) string {
			return k.Source + ":" + k.SourceAppID
		},
		func(ctx context.Context, id modelgebura.AppID) (*modelgebura.App, error) {
			res, err := g.GetApp(ctx, id)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

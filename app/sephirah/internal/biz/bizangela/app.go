package bizangela

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPullSteamAppTopic(
	a *AngelaBase,
	updateAppIndex *libmq.Topic[modelangela.UpdateAppIndex],
) *libmq.Topic[modelangela.PullApp] {
	return libmq.NewTopic[modelangela.PullApp](
		"PullApp",
		func(ctx context.Context, r *modelangela.PullApp) error {
			if !a.supv.CheckAppSource(r.Source) {
				return nil
			}
			id, err := a.searcher.NewID(ctx)
			if err != nil {
				return err
			}
			resp, err := a.porter.PullApp(
				a.supv.CallAppSource(ctx, r.Source),
				&porter.PullAppRequest{AppId: &librarian.AppID{
					Source:      r.Source,
					SourceAppId: r.AppID,
				}},
			)
			if err != nil {
				return err
			}
			app := converter.ToBizApp(resp.GetApp())
			app.ID = r.ID
			app.Source = r.Source
			internalApp := new(modelgebura.App)
			internalApp.ID = id
			internalApp.Internal = true
			internalApp.SourceAppID = strconv.FormatInt(int64(internalApp.ID), 10)
			internalApp.BoundInternal = id
			err = a.repo.UpdateApp(ctx, app, internalApp)
			if err != nil {
				return err
			}
			_ = updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{id}})
			return nil
		},
	)
}

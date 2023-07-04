package bizangela

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libmq"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPullSteamAppTopic(
	a *AngelaBase,
) *libmq.Topic[modelangela.PullSteamApp] {
	return libmq.NewTopic[modelangela.PullSteamApp](
		"PullSteamApp",
		func(ctx context.Context, r *modelangela.PullSteamApp) error {
			ctx = libapp.NewContext(ctx, string(porter.FeatureFlag_FEATURE_FLAG_SOURCE_STEAM))
			id, err := a.searcher.NewID(ctx, &pb.NewIDRequest{})
			if err != nil {
				return err
			}
			resp, err := a.porter.PullApp(ctx, &porter.PullAppRequest{AppId: &librarian.AppID{
				Source:      librarian.AppSource_APP_SOURCE_STEAM,
				SourceAppId: r.AppID,
			}})
			if err != nil {
				return err
			}
			app := converter.ToBizApp(resp.GetApp())
			app.ID = r.ID
			app.Source = modelgebura.AppSourceSteam
			internalApp := new(modelgebura.App)
			internalApp.ID = converter.ToBizInternalID(id.GetId())
			internalApp.Source = modelgebura.AppSourceInternal
			internalApp.SourceAppID = strconv.FormatInt(int64(internalApp.ID), 10)
			err = a.repo.UpdateApp(ctx, app, internalApp)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

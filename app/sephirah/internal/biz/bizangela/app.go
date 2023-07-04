package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libmq"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPullSteamAppTopic(
	a *AngelaBase,
) *libmq.Topic[modelangela.PullSteamApp] {
	return libmq.NewTopic[modelangela.PullSteamApp](
		"PullSteamApp",
		func(ctx context.Context, r *modelangela.PullSteamApp) error {
			ctx = libapp.NewContext(ctx, string(porter.FeatureFlag_FEATURE_FLAG_SOURCE_STEAM))
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
			err = a.repo.UpdateApp(ctx, app)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

package bizkether

import (
	"context"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelkether"
)

func NewPullAppInfoTopic(
	a *KetherBase,
	infoCache *libcache.Map[modelgebura.AppInfoID, modelgebura.AppInfo],
	updateAppInfoIndex *libmq.Topic[modelkether.UpdateAppInfoIndex],
) *libmq.Topic[modelkether.PullAppInfo] {
	return libmq.NewTopic[modelkether.PullAppInfo](
		"PullAppInfo",
		func(ctx context.Context, r *modelkether.PullAppInfo) error {
			// if !a.supv.HasAppInfoSource(r.AppInfoID.Source) {
			//	return nil
			//}
			// if !r.IgnoreRateLimit {
			//	if info, err := infoCache.Get(ctx, r.AppInfoID); err == nil &&
			//		info.UpdatedAt.Add(libtime.Day).After(time.Now()) {
			//		return nil
			//	}
			//}
			// id, err := a.id.New()
			// if err != nil {
			//	return err
			//}
			// resp, err := a.porter.GetAppInfo(
			//	a.supv.WithAppInfoSource(ctx, r.AppInfoID.Source),
			//	&porter.GetAppInfoRequest{
			//		Source:      r.AppInfoID.Source,
			//		SourceAppId: r.AppInfoID.SourceAppID,
			//	},
			//)
			// if err != nil {
			//	return err
			//}
			// info := converter.ToBizAppInfo(resp.GetAppInfo())
			// info.ID = r.ID
			// info.Source = r.AppInfoID.Source
			// if info.Type == modelgebura.AppTypeUnspecified {
			//	info.Type = modelgebura.AppTypeGame
			//}
			// internalInfo := new(modelgebura.AppInfo)
			// internalInfo.ID = id
			// internalInfo.SourceAppID = strconv.FormatInt(int64(internalInfo.ID), 10)
			// internalInfo.Type = info.Type
			// err = a.repo.UpsertAppInfo(ctx, info, internalInfo)
			// if err != nil {
			//	return err
			//}
			// _ = infoCache.Delete(ctx, r.AppInfoID)
			// _ = updateAppInfoIndex.Publish(ctx, modelkether.UpdateAppInfoIndex{IDs: []model.InternalID{id}})
			return nil
		},
	)
}

func NewAppInfoCache(
	g *data.GeburaRepo,
	store libcache.Store,
) *libcache.Map[modelgebura.AppInfoID, modelgebura.AppInfo] {
	return libcache.NewMap[modelgebura.AppInfoID, modelgebura.AppInfo](
		store,
		"AppInfo",
		func(k modelgebura.AppInfoID) string {
			return k.Source + ":" + k.SourceAppID
		},
		func(ctx context.Context, id modelgebura.AppInfoID) (*modelgebura.AppInfo, error) {
			res, err := g.GetAppInfo(ctx, id)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

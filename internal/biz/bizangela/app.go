package bizangela

import (
	"context"
	"strconv"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/converter"
	"github.com/tuihub/librarian/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPullAppInfoTopic(
	a *AngelaBase,
	infoCache *libcache.Map[modelgebura.AppInfoID, modelgebura.AppInfo],
	updateAppInfoIndex *libmq.Topic[modelangela.UpdateAppInfoIndex],
) *libmq.Topic[modelangela.PullAppInfo] {
	return libmq.NewTopic[modelangela.PullAppInfo](
		"PullAppInfo",
		func(ctx context.Context, r *modelangela.PullAppInfo) error {
			if !a.supv.HasAppInfoSource(r.AppInfoID.Source) {
				return nil
			}
			if r.AppInfoID.Internal {
				return nil
			}
			if !r.IgnoreRateLimit {
				if info, err := infoCache.Get(ctx, r.AppInfoID); err == nil &&
					info.LatestUpdateTime.Add(libtime.Day).After(time.Now()) {
					return nil
				}
			}
			id, err := a.id.New()
			if err != nil {
				return err
			}
			resp, err := a.porter.PullAppInfo(
				a.supv.WithAppInfoSource(ctx, r.AppInfoID.Source),
				&porter.PullAppInfoRequest{AppInfoId: &librarian.AppInfoID{
					Internal:    false,
					Source:      r.AppInfoID.Source,
					SourceAppId: r.AppInfoID.SourceAppID,
				}},
			)
			if err != nil {
				return err
			}
			info := converter.ToBizAppInfo(resp.GetAppInfo())
			info.ID = r.ID
			info.Internal = false
			info.Source = r.AppInfoID.Source
			if info.Type == modelgebura.AppTypeUnspecified {
				info.Type = modelgebura.AppTypeGame
			}
			internalInfo := new(modelgebura.AppInfo)
			internalInfo.ID = id
			internalInfo.Internal = true
			internalInfo.SourceAppID = strconv.FormatInt(int64(internalInfo.ID), 10)
			internalInfo.BoundInternal = id
			internalInfo.Type = info.Type
			err = a.repo.UpsertAppInfo(ctx, info, internalInfo)
			if err != nil {
				return err
			}
			_ = infoCache.Delete(ctx, r.AppInfoID)
			_ = updateAppInfoIndex.Publish(ctx, modelangela.UpdateAppInfoIndex{IDs: []model.InternalID{id}})
			return nil
		},
	)
}

func NewAppInfoCache(
	g bizgebura.GeburaRepo,
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

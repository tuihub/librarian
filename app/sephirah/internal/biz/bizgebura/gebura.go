package bizgebura

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type ReportAppPackageHandler interface {
	Handle(context.Context, []*modelgebura.AppBinary) *errors.Error
}

type GeburaRepo interface {
	CreateAppInfo(context.Context, *modelgebura.AppInfo) error
	UpdateAppInfo(context.Context, *modelgebura.AppInfo) error
	ListAppInfos(context.Context, model.Paging, []string, []modelgebura.AppType,
		[]model.InternalID, bool) ([]*modelgebura.AppInfo, int64, error)
	MergeAppInfos(context.Context, modelgebura.AppInfo, model.InternalID) error
	GetAppInfo(context.Context, modelgebura.AppInfoID) (*modelgebura.AppInfo, error)
	GetBoundAppInfos(context.Context, model.InternalID) ([]*modelgebura.AppInfo, error)
	GetBatchBoundAppInfos(context.Context, []model.InternalID) ([]*modelgebura.BoundAppInfos, error)
	PurchaseAppInfo(context.Context, model.InternalID,
		*modelgebura.AppInfoID, func(ctx2 context.Context) error) (model.InternalID, error)
	GetPurchasedAppInfos(context.Context, model.InternalID, string) ([]*modelgebura.BoundAppInfos, error)

	CreateApp(context.Context, model.InternalID, *modelgebura.App) error
	UpdateApp(context.Context, model.InternalID, *modelgebura.App) error
	ListApps(context.Context, model.Paging, []model.InternalID, []model.InternalID,
		[]model.InternalID, bool) ([]*modelgebura.App, int, error)
	AssignApp(context.Context, model.InternalID, model.InternalID, model.InternalID) error
	// ListAppPackageBinaryChecksumOfOneSource(context.Context, modelgebura.AppPackageSource,
	//	model.InternalID) ([]string, error)
	UnAssignApp(context.Context, model.InternalID, model.InternalID) error
	AddAppInstRunTime(context.Context, model.InternalID, model.InternalID, *model.TimeRange) error
	SumAppInstRunTime(context.Context, model.InternalID, model.InternalID, *model.TimeRange) (time.Duration, error)
	CreateAppInst(context.Context, model.InternalID, *modelgebura.AppInst) error
	UpdateAppInst(context.Context, model.InternalID, *modelgebura.AppInst) error
	ListAppInsts(context.Context, model.InternalID, model.Paging, []model.InternalID,
		[]model.InternalID, []model.InternalID) ([]*modelgebura.AppInst, int, error)
}

type Gebura struct {
	auth *libauth.Auth
	repo GeburaRepo
	// mapper         mapper.LibrarianMapperServiceClient
	searcher           *client.Searcher
	porter             porter.LibrarianPorterServiceClient
	supv               *supervisor.Supervisor
	updateAppInfoIndex *libmq.Topic[modelangela.UpdateAppInfoIndex]
	pullAppInfo        *libmq.Topic[modelangela.PullAppInfo]
	appInfoCache       *libcache.Map[modelgebura.AppInfoID, modelgebura.AppInfo]
}

func NewGebura(
	repo GeburaRepo,
	auth *libauth.Auth,
	// mClient mapper.LibrarianMapperServiceClient,
	sClient *client.Searcher,
	pClient porter.LibrarianPorterServiceClient,
	supv *supervisor.Supervisor,
	updateAppIndex *libmq.Topic[modelangela.UpdateAppInfoIndex],
	pullAppInfo *libmq.Topic[modelangela.PullAppInfo],
	appInfoCache *libcache.Map[modelgebura.AppInfoID, modelgebura.AppInfo],
) *Gebura {
	return &Gebura{
		auth: auth,
		repo: repo,
		//mapper:         mClient,
		searcher:           sClient,
		porter:             pClient,
		supv:               supv,
		updateAppInfoIndex: updateAppIndex,
		pullAppInfo:        pullAppInfo,
		appInfoCache:       appInfoCache,
	}
}

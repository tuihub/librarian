package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelkether"
	"github.com/tuihub/librarian/internal/service/supervisor"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type ReportAppPackageHandler interface {
	Handle(context.Context, []*modelgebura.AppBinary) *errors.Error
}

type Gebura struct {
	auth               *libauth.Auth
	repo               *data.GeburaRepo
	id                 *libidgenerator.IDGenerator
	search             libsearch.Search
	porter             porter.LibrarianPorterServiceClient
	supv               *supervisor.Supervisor
	updateAppInfoIndex *libmq.Topic[modelkether.UpdateAppInfoIndex]
	pullAppInfo        *libmq.Topic[modelkether.PullAppInfo]
	appInfoCache       *libcache.Map[modelgebura.AppInfoID, modelgebura.AppInfo]
}

func NewGebura(
	repo *data.GeburaRepo,
	auth *libauth.Auth,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	pClient porter.LibrarianPorterServiceClient,
	supv *supervisor.Supervisor,
	updateAppIndex *libmq.Topic[modelkether.UpdateAppInfoIndex],
	pullAppInfo *libmq.Topic[modelkether.PullAppInfo],
	appInfoCache *libcache.Map[modelgebura.AppInfoID, modelgebura.AppInfo],
) *Gebura {
	return &Gebura{
		auth:               auth,
		repo:               repo,
		id:                 id,
		search:             search,
		porter:             pClient,
		supv:               supv,
		updateAppInfoIndex: updateAppIndex,
		pullAppInfo:        pullAppInfo,
		appInfoCache:       appInfoCache,
	}
}

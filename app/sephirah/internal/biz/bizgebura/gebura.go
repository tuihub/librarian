package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type ReportAppPackageHandler interface {
	Handle(context.Context, []*modelgebura.AppPackage) *errors.Error
}

type GeburaRepo interface {
	CreateApp(context.Context, *modelgebura.App) error
	UpdateApp(context.Context, *modelgebura.App) error
	UpsertApps(context.Context, []*modelgebura.App) error
	ListApps(context.Context, model.Paging, []modelgebura.AppSource, []modelgebura.AppType,
		[]model.InternalID, bool) ([]*modelgebura.App, int64, error)
	MergeApps(context.Context, modelgebura.App, model.InternalID) error
	SearchApps(context.Context, model.Paging, string) ([]*modelgebura.App, int, error)
	GetBindApps(context.Context, model.InternalID) ([]*modelgebura.App, error)
	PurchaseApp(context.Context, model.InternalID, model.InternalID) error
	GetPurchasedApps(context.Context, model.InternalID) ([]model.InternalID, error)

	CreateAppPackage(context.Context, *modelgebura.AppPackage) error
	UpdateAppPackage(context.Context, *modelgebura.AppPackage) error
	UpsertAppPackages(context.Context, model.InternalID, []*modelgebura.AppPackage) error
	ListAppPackages(context.Context, model.Paging, []modelgebura.AppPackageSource,
		[]model.InternalID) ([]*modelgebura.AppPackage, int, error)
	AssignAppPackage(context.Context, model.InternalID, model.InternalID, model.InternalID) error
	ListAllAppPackageIDOfOneSource(context.Context, modelgebura.AppPackageSource,
		model.InternalID) ([]string, error)
	UnAssignAppPackage(context.Context, model.InternalID, model.InternalID) error
}

type Gebura struct {
	auth     *libauth.Auth
	repo     GeburaRepo
	mapper   mapper.LibrarianMapperServiceClient
	porter   porter.LibrarianPorterServiceClient
	searcher searcher.LibrarianSearcherServiceClient
}

func NewGebura(
	repo GeburaRepo,
	auth *libauth.Auth,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) *Gebura {
	return &Gebura{auth: auth, repo: repo, mapper: mClient, porter: pClient, searcher: sClient}
}

package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/lib/libauth"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type App struct {
	InternalID       int64
	Source           AppSource
	SourceAppID      string
	SourceURL        string
	Name             string
	Type             AppType
	ShortDescription string
	ImageURL         string
	Details          *AppDetails
}

type AppDetails struct {
	Description string
	ReleaseDate string
	Developer   string
	Publisher   string
	Version     string
}

type AppSource int

const (
	AppSourceUnspecified AppSource = iota
	AppSourceInternal
	AppSourceSteam
)

type AppType int

const (
	AppTypeUnspecified AppType = iota
	AppTypeGame
)

type AppPackage struct {
	InternalID      int64
	Source          AppPackageSource
	SourceID        int64
	SourcePackageID string
	Name            string
	Description     string
	Binary          *AppPackageBinary
}

type AppPackageBinary struct {
	Name      string
	Size      int64
	PublicURL string
}

type AppPackageSource int

const (
	AppPackageSourceUnspecified AppPackageSource = iota
	AppPackageSourceManual
	AppPackageSourceSentinel
)

type ReportAppPackageHandler interface {
	Handle(context.Context, []*AppPackage) *errors.Error
}

type Paging struct {
	PageSize int
	PageNum  int
}

type GeburaRepo interface {
	IsApp(context.Context, int64) error
	CreateApp(context.Context, *App) error
	UpdateApp(context.Context, *App) error
	UpsertApp(context.Context, []*App) error
	ListApp(context.Context, Paging, []AppSource, []AppType, []int64, bool) ([]*App, error)
	IsAppPackage(context.Context, int64) error
	CreateAppPackage(context.Context, *AppPackage) error
	UpdateAppPackage(context.Context, *AppPackage) error
	UpsertAppPackage(context.Context, []*AppPackage) error
	ListAppPackage(context.Context, Paging, []AppPackageSource, []int64) ([]*AppPackage, error)
	ListAllAppPackageIDOfOneSource(context.Context, AppPackageSource, int64) ([]string, error)
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
	block bizbinah.CallbackControlBlock,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) *Gebura {
	block.RegisterUploadCallback(bizbinah.UploadCallback{
		ID:   bizbinah.UploadArtifacts,
		Func: UploadArtifactsCallback,
	})
	return &Gebura{auth: auth, repo: repo, mapper: mClient, porter: pClient, searcher: sClient}
}

func UploadArtifactsCallback(file *bizbinah.UploadFile) error {
	panic("not impl")
}

package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type App struct {
	InternalID      int64
	Source          AppSource
	SourceAppID     string
	SourceURL       string
	Name            string
	Type            AppType
	ShorDescription string
	ImageURL        string
	Details         *AppDetails
}

type AppDetails struct {
	Description string
	ReleaseDate string
	Developer   string
	Publisher   string
}

type AppSource int

const (
	AppSourceUnspecified AppSource = iota
	AppSourceInternal
	AppSourceSteam
)

type AppType int

const (
	AppTypeGeneral AppType = iota
	AppTypeGame
)

type Paging struct {
	PageSize int
	PageNum  int
}

// GeburaRepo is an App repo.
type GeburaRepo interface {
	CreateApp(context.Context, *App) error
	UpdateApp(context.Context, *App) error
	ListApp(context.Context, Paging, []AppSource, []AppType, []int64, bool) ([]*App, error)
}

// GeburaUseCase is an App use case.
type GeburaUseCase struct {
	auth     *libauth.Auth
	repo     GeburaRepo
	mapper   mapper.LibrarianMapperServiceClient
	porter   porter.LibrarianPorterServiceClient
	searcher searcher.LibrarianSearcherServiceClient
}

// NewGeburaUseCase new an App use case.
func NewGeburaUseCase(
	repo GeburaRepo,
	auth *libauth.Auth,
	block bizbinah.CallbackControlBlock,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) *GeburaUseCase {
	block.RegisterUploadCallback(bizbinah.UploadCallback{
		ID:   bizbinah.UploadArtifacts,
		Func: UploadArtifactsCallback,
	})
	return &GeburaUseCase{auth: auth, repo: repo, mapper: mClient, porter: pClient, searcher: sClient}
}

func (g *GeburaUseCase) CreateApp(ctx context.Context, app *App) (*App, *errors.Error) {
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	app.InternalID = resp.Id
	app.Source = AppSourceInternal
	app.SourceAppID = ""
	app.SourceURL = ""
	err = g.repo.CreateApp(ctx, app)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return app, nil
}

func (g *GeburaUseCase) UpdateApp(ctx context.Context, app *App) *errors.Error {
	app.Source = AppSourceInternal
	app.SourceAppID = ""
	app.SourceURL = ""
	err := g.repo.UpdateApp(ctx, app)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *GeburaUseCase) ListApp(
	ctx context.Context,
	paging Paging,
	sources []AppSource,
	types []AppType,
	ids []int64,
	containDetails bool,
	withBind bool,
) ([]*App, *errors.Error) {
	if withBind {
		return nil, pb.ErrorErrorReasonNotImplemented("not support")
	}
	apps, err := g.repo.ListApp(ctx, paging, sources, types, ids, containDetails)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}

func (g *GeburaUseCase) BindApp(ctx context.Context, internal App, bind App) (*App, *errors.Error) {
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	bind.InternalID = resp.Id
	if err = g.repo.CreateApp(ctx, &bind); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &bind, nil
}

func UploadArtifactsCallback(file *bizbinah.UploadFile) error {
	panic("not impl")
}

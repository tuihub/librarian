package biz

import (
	"context"

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
	AppSourceInternal AppSource = iota
	AppSourceSteam
)

type AppType int

const (
	AppTypeGeneral AppType = iota
	AppTypeGame
)

// GeburaRepo is an App repo.
type GeburaRepo interface {
	CreateApp(context.Context, *App) error
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
func NewGeburaUseCase(repo GeburaRepo, auth *libauth.Auth, mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient, sClient searcher.LibrarianSearcherServiceClient) *GeburaUseCase {
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
		return nil, pb.ErrorErrorReasonUnspecified("create failed")
	}
	return app, nil
}

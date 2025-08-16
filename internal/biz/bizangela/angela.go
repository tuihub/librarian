package bizangela

import (
	"context"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/model/modelangela"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

type Angela struct {
	auth   *libauth.Auth
	repo   *data.AngelaRepo
	id     *libidgenerator.IDGenerator
	search libsearch.Search
	porter porter.LibrarianPorterServiceClient
	supv   *data.SupervisorRepo
}

func NewAngela(
	repo *data.AngelaRepo,
	auth *libauth.Auth,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	pClient porter.LibrarianPorterServiceClient,
	supv *data.SupervisorRepo,
) *Angela {
	return &Angela{
		auth:   auth,
		repo:   repo,
		id:     id,
		search: search,
		porter: pClient,
		supv:   supv,
	}
}

func (a *Angela) GetServerInstanceSummary(ctx context.Context) (*modelangela.ServerInstanceSummary, error) {
	summary, err := a.repo.GetServerInstanceSummary(ctx)
	if err != nil {
		return nil, err
	}
	if summary == nil {
		summary = new(modelangela.ServerInstanceSummary)
	}
	return summary, nil
}

func (a *Angela) SetServerInstanceSummary(
	ctx context.Context,
	summary *modelangela.ServerInstanceSummary,
) error {
	return a.repo.SetServerInstanceSummary(ctx, summary)
}

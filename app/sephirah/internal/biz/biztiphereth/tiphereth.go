package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
)

type TipherethRepo interface {
	FetchUserByPassword(context.Context, *modeltiphereth.User) (*modeltiphereth.User, error)
	CreateUser(context.Context, *modeltiphereth.User, model.InternalID) error
	UpdateUser(context.Context, *modeltiphereth.User, string) error
	ListUsers(context.Context, model.Paging, []model.InternalID,
		[]libauth.UserType, []modeltiphereth.UserStatus, []model.InternalID,
		model.InternalID) ([]*modeltiphereth.User, int64, error)
	LinkAccount(context.Context, modeltiphereth.Account, model.InternalID) error
	UnLinkAccount(context.Context, modeltiphereth.Account, model.InternalID) error
	ListLinkAccounts(context.Context, model.InternalID) ([]*modeltiphereth.Account, error)
	GetUser(context.Context, model.InternalID) (*modeltiphereth.User, error)
	UpsertPorters(context.Context, []*modeltiphereth.PorterInstance) error
	ListPorters(context.Context, model.Paging) ([]*modeltiphereth.PorterInstance, int64, error)
	UpdatePorterStatus(context.Context, model.InternalID, modeltiphereth.PorterInstanceStatus) error
	UpdatePorterPrivilege(context.Context, model.InternalID, model.InternalID,
		*modeltiphereth.PorterInstancePrivilege) error
	FetchPorterPrivilege(context.Context, model.InternalID, model.InternalID) (
		*modeltiphereth.PorterInstancePrivilege, error)
}

type Tiphereth struct {
	auth        *libauth.Auth
	repo        TipherethRepo
	supv        *supervisor.Supervisor
	mapper      mapper.LibrarianMapperServiceClient
	searcher    *client.Searcher
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo]
}

func NewTiphereth(
	repo TipherethRepo,
	auth *libauth.Auth,
	supv *supervisor.Supervisor,
	mClient mapper.LibrarianMapperServiceClient,
	sClient *client.Searcher,
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo],
	cron *libcron.Cron,
) (*Tiphereth, error) {
	t := &Tiphereth{
		auth:        auth,
		repo:        repo,
		supv:        supv,
		mapper:      mClient,
		searcher:    sClient,
		pullAccount: pullAccount,
	}
	err := cron.BySeconds(60, t.updatePorters, context.Background()) //nolint:gomnd // hard code min interval
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Tiphereth) CreateDefaultAdmin(ctx context.Context, user *modeltiphereth.User) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return
	}
	user.PassWord = password
	id, err := t.searcher.NewID(ctx)
	if err != nil {
		return
	}
	user.ID = id
	user.Status = modeltiphereth.UserStatusActive
	user.Type = libauth.UserTypeAdmin
	if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  int64(user.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
			Prop: nil,
		},
	}}); err != nil {
		return
	}
	if err = t.repo.CreateUser(ctx, user, user.ID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return
	}
}

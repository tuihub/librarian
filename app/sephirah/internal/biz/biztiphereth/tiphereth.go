package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTiphereth,
	NewUserCountCache,
)

type TipherethRepo interface {
	FetchUserByPassword(context.Context, string, string) (*modeltiphereth.User, error)
	CreateUser(context.Context, *modeltiphereth.User, model.InternalID) error
	UpdateUser(context.Context, *modeltiphereth.User, string) error
	ListUsers(context.Context, model.Paging, []model.InternalID,
		[]libauth.UserType, []modeltiphereth.UserStatus, []model.InternalID,
		model.InternalID) ([]*modeltiphereth.User, int64, error)
	GetUserCount(context.Context) (int, error)
	LinkAccount(context.Context, modeltiphereth.Account, model.InternalID) (model.InternalID, error)
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
	CreateDevice(context.Context, *modeltiphereth.DeviceInfo) error
	ListUserSessions(context.Context, model.InternalID) ([]*modeltiphereth.UserSession, error)
	DeleteUserSession(context.Context, model.InternalID, model.InternalID) error
	FetchDeviceInfo(context.Context, model.InternalID) (*modeltiphereth.DeviceInfo, error)
	CreateUserSession(context.Context, *modeltiphereth.UserSession) error
	FetchUserSession(context.Context, model.InternalID, string) (*modeltiphereth.UserSession, error)
	UpdateUserSession(context.Context, *modeltiphereth.UserSession) error
	ListDevices(context.Context, model.InternalID) ([]*modeltiphereth.DeviceInfo, error)
}

type Tiphereth struct {
	app  *libapp.Settings
	auth *libauth.Auth
	repo TipherethRepo
	supv *supervisor.Supervisor
	// mapper      mapper.LibrarianMapperServiceClient
	searcher       *client.Searcher
	pullAccount    *libmq.Topic[modeltiphereth.PullAccountInfo]
	userCountCache *libcache.Key[modeltiphereth.UserCount]
}

func NewTiphereth(
	app *libapp.Settings,
	repo TipherethRepo,
	auth *libauth.Auth,
	supv *supervisor.Supervisor,
	// mClient mapper.LibrarianMapperServiceClient,
	sClient *client.Searcher,
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo],
	cron *libcron.Cron,
	userCountCache *libcache.Key[modeltiphereth.UserCount],
) (*Tiphereth, error) {
	t := &Tiphereth{
		app:  app,
		auth: auth,
		repo: repo,
		supv: supv,
		//mapper:      mClient,
		searcher:       sClient,
		pullAccount:    pullAccount,
		userCountCache: userCountCache,
	}
	err := cron.BySeconds(
		"TipherethUpdatePorter",
		updatePorterInterval,
		t.updatePorters, context.Background(),
	)
	if err != nil {
		return nil, err
	}
	return t, nil
}

const (
	updatePorterInterval = 10
	demoAdminUserName    = "admin"
	demoAdminPassword    = "admin"
)

func (t *Tiphereth) CreateConfiguredAdmin() {
	ctx := context.Background()
	if !(t.app.EnvExist(libapp.EnvDemoMode) || t.app.EnvExist(libapp.EnvCreateAdminUserName)) {
		return
	}
	user := &modeltiphereth.User{
		ID:       0,
		UserName: demoAdminUserName,
		PassWord: demoAdminPassword,
		Type:     libauth.UserTypeAdmin,
		Status:   modeltiphereth.UserStatusActive,
	}
	if username, err := t.app.Env(libapp.EnvCreateAdminUserName); err == nil && username != "" {
		user.UserName = username
	}
	if password, err := t.app.Env(libapp.EnvCreateAdminPassword); err == nil && password != "" {
		user.PassWord = password
	}
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
	// if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
	//	{
	//		Vid:  int64(user.ID),
	//		Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
	//		Prop: nil,
	//	},
	// }}); err != nil {
	//	return
	//}
	if err = t.repo.CreateUser(ctx, user, user.ID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return
	}
}

func NewUserCountCache(
	t TipherethRepo,
	store libcache.Store,
) *libcache.Key[modeltiphereth.UserCount] {
	return libcache.NewKey[modeltiphereth.UserCount](
		store,
		"UserCount",
		func(ctx context.Context) (*modeltiphereth.UserCount, error) {
			res, err := t.GetUserCount(ctx)
			if err != nil {
				return nil, err
			}
			return &modeltiphereth.UserCount{Count: res}, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

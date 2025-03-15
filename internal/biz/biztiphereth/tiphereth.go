package biztiphereth

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/service/supervisor"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTiphereth,
	NewUserCountCache,
	NewPorterInstanceCache,
	NewPorterContextCache,
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
	UpsertPorters(context.Context, []*modelsupervisor.PorterInstance) error
	ListPorters(context.Context, model.Paging) ([]*modelsupervisor.PorterInstance, int64, error)
	FetchPorterByAddress(context.Context, string) (*modelsupervisor.PorterInstance, error)
	UpdatePorterStatus(context.Context, model.InternalID,
		modeltiphereth.UserStatus) (*modelsupervisor.PorterInstance, error)
	CreatePorterContext(context.Context, model.InternalID, *modelsupervisor.PorterContext) error
	GetEnabledPorterContexts(context.Context) ([]*modelsupervisor.PorterContext, error)
	ListPorterContexts(context.Context, model.InternalID, model.Paging) ([]*modelsupervisor.PorterContext, int64, error)
	UpdatePorterContext(context.Context, model.InternalID, *modelsupervisor.PorterContext) error
	FetchPorterContext(context.Context, model.InternalID) (*modelsupervisor.PorterContext, error)
	CreateDevice(context.Context, model.InternalID, *modeltiphereth.DeviceInfo, *string) (model.InternalID, error)
	ListUserSessions(context.Context, model.InternalID) ([]*modeltiphereth.UserSession, error)
	DeleteUserSession(context.Context, model.InternalID, model.InternalID) error
	FetchDeviceInfo(context.Context, model.InternalID) (*modeltiphereth.DeviceInfo, error)
	CreateUserSession(context.Context, *modeltiphereth.UserSession) error
	FetchUserSession(context.Context, model.InternalID, string) (*modeltiphereth.UserSession, error)
	UpdateUserSession(context.Context, *modeltiphereth.UserSession) error
	ListDevices(context.Context, model.InternalID) ([]*modeltiphereth.DeviceInfo, error)
	ListPorterGroups(context.Context, []modeltiphereth.UserStatus) ([]*modelsupervisor.PorterGroup, error)
}

type Tiphereth struct {
	app                 *libapp.Settings
	auth                *libauth.Auth
	repo                TipherethRepo
	supv                *supervisor.Supervisor
	id                  *libidgenerator.IDGenerator
	search              libsearch.Search
	pullAccount         *libmq.Topic[modeltiphereth.PullAccountInfo]
	userCountCache      *libcache.Key[modeltiphereth.UserCount]
	porterInstanceCache *libcache.Map[string, modelsupervisor.PorterInstance]
}

func NewTiphereth(
	app *libapp.Settings,
	repo TipherethRepo,
	auth *libauth.Auth,
	supv *supervisor.Supervisor,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo],
	cron *libcron.Cron,
	userCountCache *libcache.Key[modeltiphereth.UserCount],
	porterInstanceCache *libcache.Map[string, modelsupervisor.PorterInstance],
) (*Tiphereth, error) {
	t := &Tiphereth{
		app:                 app,
		auth:                auth,
		repo:                repo,
		supv:                supv,
		id:                  id,
		search:              search,
		pullAccount:         pullAccount,
		userCountCache:      userCountCache,
		porterInstanceCache: porterInstanceCache,
	}
	err := cron.Duration(
		"TipherethUpdatePorter",
		supv.GetHeartbeatInterval(),
		t.updatePorters, context.Background(),
	)
	if err != nil {
		return nil, err
	}
	return t, nil
}

const (
	demoAdminUserName = "admin"
	demoAdminPassword = "admin"
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
	id, err := t.id.New()
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

func NewPorterInstanceCache(
	t TipherethRepo,
	store libcache.Store,
) *libcache.Map[string, modelsupervisor.PorterInstance] {
	return libcache.NewMap[string, modelsupervisor.PorterInstance](
		store,
		"PorterInstanceCache",
		func(s string) string {
			return s
		},
		func(ctx context.Context, s string) (*modelsupervisor.PorterInstance, error) {
			return t.FetchPorterByAddress(ctx, s)
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

func NewPorterContextCache(
	t TipherethRepo,
	store libcache.Store,
) *libcache.Map[model.InternalID, modelsupervisor.PorterContext] {
	return libcache.NewMap[model.InternalID, modelsupervisor.PorterContext](
		store,
		"PorterContextCache",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		func(ctx context.Context, k model.InternalID) (*modelsupervisor.PorterContext, error) {
			return t.FetchPorterContext(ctx, k)
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

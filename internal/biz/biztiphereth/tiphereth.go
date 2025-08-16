package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTiphereth,
	NewUserCountCache,
)

type Tiphereth struct {
	app            *libapp.Settings
	auth           *libauth.Auth
	repo           *data.TipherethRepo
	supv           *data.SupervisorRepo
	id             *libidgenerator.IDGenerator
	search         libsearch.Search
	pullAccount    *libmq.Topic[model.PullAccountInfo]
	userCountCache *libcache.Key[model.UserCount]
}

func NewTiphereth(
	app *libapp.Settings,
	repo *data.TipherethRepo,
	auth *libauth.Auth,
	supv *data.SupervisorRepo,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	pullAccount *libmq.Topic[model.PullAccountInfo],
	userCountCache *libcache.Key[model.UserCount],
) (*Tiphereth, error) {
	t := &Tiphereth{
		app:            app,
		auth:           auth,
		repo:           repo,
		supv:           supv,
		id:             id,
		search:         search,
		pullAccount:    pullAccount,
		userCountCache: userCountCache,
	}
	return t, nil
}

const (
	demoAdminUserName = "admin"
	demoAdminPassword = "admin"
)

func (t *Tiphereth) CreateConfiguredAdmin() {
	ctx := context.Background()
	if !t.app.EnvExist(libapp.EnvDemoMode) && !t.app.EnvExist(libapp.EnvCreateAdminUserName) {
		return
	}
	user := &model.User{
		ID:       0,
		Username: demoAdminUserName,
		Password: demoAdminPassword,
		Type:     model.UserTypeAdmin,
		Status:   model.UserStatusActive,
	}
	if username, err := t.app.Env(libapp.EnvCreateAdminUserName); err == nil && username != "" {
		user.Username = username
	}
	if password, err := t.app.Env(libapp.EnvCreateAdminPassword); err == nil && password != "" {
		user.Password = password
	}
	password, err := t.auth.GeneratePassword(user.Password)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return
	}
	user.Password = password
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
	t *data.TipherethRepo,
	store libcache.Store,
) *libcache.Key[model.UserCount] {
	return libcache.NewKey[model.UserCount](
		store,
		"UserCount",
		func(ctx context.Context) (*model.UserCount, error) {
			res, err := t.GetUserCount(ctx)
			if err != nil {
				return nil, err
			}
			return &model.UserCount{Count: res}, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

package angelaweb

import (
	"context"
	"embed"
	"net/http"

	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAngelaWeb)

type AngelaWeb struct {
	apiHandler  *api.Handler
	pageBuilder *page.Builder
	auth        *libauth.Auth
	app         *fiber.App
}

// Embed view
//
//go:embed view/*
var embedDirView embed.FS

// Embed static
//
//go:embed static/*
var embedDirStatic embed.FS

func NewAngelaWeb(
	auth *libauth.Auth,
	t *biztiphereth.Tiphereth,
	userCountCache *libcache.Key[model.UserCount],
) *AngelaWeb {
	viewsEngine := html.NewFileSystem(http.FS(embedDirView), ".html")
	viewsEngine.Directory = "view"

	app := fiber.New(fiber.Config{ //nolint:exhaustruct // no need
		Views:       viewsEngine,
		ViewsLayout: "layout/default",
	})

	app.Use(logger.New())

	res := &AngelaWeb{
		apiHandler:  api.NewHandler(t, userCountCache),
		pageBuilder: page.NewBuilder(t, userCountCache),
		auth:        auth,
		app:         app,
	}
	res.setupRoutes()
	return res
}

func (a *AngelaWeb) Start(ctx context.Context) error {
	return a.app.Listen(":3000")
}

func (a *AngelaWeb) Stop(ctx context.Context) error {
	return a.app.ShutdownWithContext(ctx)
}

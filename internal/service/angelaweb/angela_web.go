package angelaweb

import (
	"context"
	"embed"
	"fmt"
	"io"
	"net/http"

	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
	"github.com/google/wire"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/samber/lo"
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

// Embed locales
//
//go:embed locales/*
var embedDirLocales embed.FS

func NewAngelaWeb(
	settings *libapp.Settings,
	digests []*conf.ConfigDigest,
	auth *libauth.Auth,
	t *biztiphereth.Tiphereth,
	userCountCache *libcache.Key[model.UserCount],
) *AngelaWeb {
	viewsEngine := html.NewFileSystem(http.FS(embedDirView), ".html")
	viewsEngine.Directory = "view"

	viewsEngine.AddFunc("localize", func(c *fiber.Ctx, args ...any) string {
		strArgs := lo.Map(args, func(v any, _ int) string {
			return fmt.Sprint(v)
		})
		if len(strArgs) == 0 {
			return ""
		}
		key := strArgs[0]
		if len(strArgs) == 1 {
			localize, err := fiberi18n.Localize(c, key)
			if err != nil {
				return key
			}
			return localize
		}
		if len(strArgs)%2 == 0 {
			return key
		}
		data := make(map[string]string)
		for i := 1; i < len(strArgs); i += 2 {
			data[strArgs[i]] = strArgs[i+1]
		}
		localize, err := fiberi18n.Localize(c, &i18n.LocalizeConfig{ //nolint:exhaustruct // no need
			MessageID:    key,
			TemplateData: data,
		})
		if err != nil {
			return key
		}
		return localize
	})

	fiberlog.SetOutput(io.Discard)

	app := fiber.New(fiber.Config{ //nolint:exhaustruct // no need
		Views:       viewsEngine,
		ViewsLayout: "layout/default",
	})

	res := &AngelaWeb{
		apiHandler:  api.NewHandler(t, userCountCache),
		pageBuilder: page.NewBuilder(t, digests, userCountCache),
		auth:        auth,
		app:         app,
	}
	res.setupMiddlewares(settings)
	res.setupRoutes()
	return res
}

func (a *AngelaWeb) Start(ctx context.Context) error {
	return a.app.Listen(":3000")
}

func (a *AngelaWeb) Stop(ctx context.Context) error {
	return a.app.ShutdownWithContext(ctx)
}

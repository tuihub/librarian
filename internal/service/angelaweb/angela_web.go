package angelaweb

import (
	"context"
	"embed"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/biz/bizangela"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/bizsupervisor"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libobserve"
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
	addr        string
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
	c *conf.Server,
	settings *libapp.Settings,
	digests []*conf.ConfigDigest,
	auth *libauth.Auth,
	a *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	s *bizsupervisor.Supervisor,
	userCountCache *libcache.Key[model.UserCount],
	observer *libobserve.Observe,
) *AngelaWeb {
	viewsEngine := html.NewFileSystem(http.FS(embedDirView), ".go.html")
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
		Views:                 viewsEngine,
		ViewsLayout:           "layout/default",
		DisableStartupMessage: true,
	})

	res := &AngelaWeb{
		apiHandler:  api.NewHandler(a, t, g, userCountCache, observer),
		pageBuilder: page.NewBuilder(settings, a, t, g, s, digests, userCountCache),
		auth:        auth,
		app:         app,
		addr:        net.JoinHostPort(c.Admin.Host, strconv.Itoa(int(c.Admin.Port))),
	}
	res.setupMiddlewares(settings)
	res.setupRoutes()
	return res
}

func (a *AngelaWeb) Start(ctx context.Context) error {
	return a.app.Listen(a.addr)
}

func (a *AngelaWeb) Stop(ctx context.Context) error {
	return a.app.ShutdownWithContext(ctx)
}

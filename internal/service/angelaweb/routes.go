package angelaweb

import (
	"net/http"

	"github.com/tuihub/librarian/internal/service/angelaweb/locales"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func (a *AngelaWeb) setupRoutes() {
	// 静态文件
	a.app.Use("/static", filesystem.New(filesystem.Config{ //nolint: exhaustruct // no need
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
		Browse:     true,
	}))

	// i18n Middleware
	a.app.Use(fiberi18n.New(&fiberi18n.Config{ //nolint:exhaustruct // no need
		RootPath:         "locales",
		FormatBundleFile: "toml",
		UnmarshalFunc:    toml.Unmarshal,
		Loader: &fiberi18n.EmbedLoader{
			FS: embedDirLocales,
		},
		DefaultLanguage: locales.DefaultLanguage(),
		AcceptLanguages: locales.SupportedLanguages(),
		LangHandler:     locales.LangHandler,
	}))

	// CORS
	a.app.Use(cors.New())

	// API路由
	api := a.app.Group("/api")
	api.Post("/login", a.apiHandler.Login)

	api.Use(tokenMiddleware(a.auth, a.pageBuilder))

	// 受保护的API路由
	api.Get("/users", a.apiHandler.ListUsers)
	api.Post("/users", a.apiHandler.CreateUser)
	api.Get("/users/:id", a.apiHandler.GetUser)
	api.Put("/users/:id", a.apiHandler.UpdateUser)

	api.Get("/porters", a.apiHandler.ListPorters)
	api.Put("/porters/:id/status", a.apiHandler.UpdatePorterStatus)

	api.Get("/dashboard/stats", a.apiHandler.GetDashboardStats)

	// 页面路由
	a.app.Get("/login", a.pageBuilder.Login)

	// 受保护的页面路由
	auth := a.app.Group("/")
	auth.Use(tokenMiddleware(a.auth, a.pageBuilder))

	auth.Get("/", a.pageBuilder.Dashboard)
	auth.Get("/users", a.pageBuilder.UserList)
	auth.Get("/users/new", a.pageBuilder.UserForm)
	auth.Get("/users/edit/:id", a.pageBuilder.UserForm)

	auth.Get("/porters", a.pageBuilder.PorterList)

	auth.Get("/config", a.pageBuilder.ConfigList)

	// 404 处理 - 放在所有路由之后
	a.app.Use(a.pageBuilder.NotFound)
}

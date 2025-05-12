package angelaweb

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (a *AngelaWeb) setupRoutes() {
	a.app.Use("/static", filesystem.New(filesystem.Config{ //nolint: exhaustruct // no need
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
		Browse:     true,
	}))

	api := a.app.Group("/api")
	api.Post("/login", a.apiHandler.Login)

	api.Use(tokenMiddleware(a.auth, a.pageBuilder))

	api.Get("/users", a.apiHandler.ListUsers)
	api.Post("/users", a.apiHandler.CreateUser)
	api.Get("/users/:id", a.apiHandler.GetUser)
	api.Put("/users/:id", a.apiHandler.UpdateUser)

	api.Get("/porters", a.apiHandler.ListPorters)
	api.Put("/porters/:id/status", a.apiHandler.UpdatePorterStatus)

	api.Post("/server-info", a.apiHandler.UpdateServerInfo)

	api.Get("/dashboard/stats", a.apiHandler.GetDashboardStats)

	api.Get("/sentinels", a.apiHandler.ListSentinels)
	api.Post("/sentinels", a.apiHandler.CreateSentinel)
	api.Get("/sentinels/:id", a.apiHandler.GetSentinel)
	api.Put("/sentinels/:id", a.apiHandler.UpdateSentinel)
	api.Post("/sentinels/:id/sessions", a.apiHandler.CreateSentinelSession)
	api.Put("/sentinels/sessions/:id/status", a.apiHandler.UpdateSentinelSessionStatus)
	api.Delete("/sentinels/sessions/:id", a.apiHandler.DeleteSentinelSession)

	page := a.app.Group("/")
	page.Get("/login", a.pageBuilder.Login)

	page.Use(tokenMiddleware(a.auth, a.pageBuilder))

	page.Get("/", a.pageBuilder.Dashboard)
	page.Get("/users", a.pageBuilder.UserList)
	page.Get("/users/new", a.pageBuilder.UserForm)
	page.Get("/users/edit/:id", a.pageBuilder.UserForm)

	page.Get("/porters", a.pageBuilder.PorterList)

	page.Get("/config", a.pageBuilder.ConfigList)
	page.Get("/server-info", a.pageBuilder.ServerInfoForm)
	page.Get("/monitor", monitor.New())

	page.Get("/sentinels", a.pageBuilder.SentinelList)
	page.Get("/sentinels/new", a.pageBuilder.SentinelForm)
	page.Get("/sentinels/edit/:id", a.pageBuilder.SentinelForm)
	page.Get("/sentinels/:id", a.pageBuilder.SentinelDetail)

	a.app.Use(a.pageBuilder.NotFound)
}

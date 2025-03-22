package angelaweb

import (
	"net/http"

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

	// CORS
	a.app.Use(cors.New())

	// API路由
	api := a.app.Group("/api")
	api.Post("/login", a.apiHandler.Login)

	api.Use(tokenMiddleware(a.auth))

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
	auth.Use(tokenMiddleware(a.auth))

	auth.Get("/", a.pageBuilder.Dashboard)
	auth.Get("/users", a.pageBuilder.UserList)
	auth.Get("/users/new", a.pageBuilder.UserForm)
	auth.Get("/users/edit/:id", a.pageBuilder.UserForm)

	auth.Get("/porters", a.pageBuilder.PorterList)
}

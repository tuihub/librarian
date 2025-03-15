package angelaweb

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func (a *AngelaWeb) setupRoutes() {
	// 静态文件
	a.app.Static("/static", "./static")

	// CORS
	a.app.Use(cors.New())

	// API路由
	api := a.app.Group("/api")
	api.Post("/login", a.apiHandler.Login)

	// JWT中间件
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("your-secret-key"),
	}))

	// 受保护的API路由
	api.Get("/users", a.apiHandler.ListUsers)
	api.Post("/users", a.apiHandler.CreateUser)
	api.Get("/users/:id", a.apiHandler.GetUser)
	api.Put("/users/:id", a.apiHandler.UpdateUser)
	api.Delete("/users/:id", a.apiHandler.DeleteUser)
	api.Get("/dashboard/stats", a.apiHandler.GetDashboardStats)

	// 页面路由
	a.app.Get("/login", a.pageBuilder.Login)

	// 受保护的页面路由
	auth := a.app.Group("/")
	auth.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("your-secret-key"),
	}))

	auth.Get("/", a.pageBuilder.Dashboard)
	auth.Get("/users", a.pageBuilder.UserList)
	auth.Get("/users/new", a.pageBuilder.UserForm)
	auth.Get("/users/edit/:id", a.pageBuilder.UserForm)
}

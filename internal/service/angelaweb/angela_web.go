package angelaweb

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"
)

type AngelaWeb struct {
	apiHandler  *api.Handler
	pageBuilder *page.Builder
	app         *fiber.App
}

func NewAngelaWeb(handler *api.Handler, builder *page.Builder) *AngelaWeb {
	viewsEngine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{
		Views:       viewsEngine,
		ViewsLayout: "layout/default",
	})

	res := &AngelaWeb{
		apiHandler:  handler,
		pageBuilder: builder,
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

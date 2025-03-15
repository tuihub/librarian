package angelaweb

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"
)

type AngelaWeb struct {
	handler *api.Handler
	builder *page.Builder
	app     *fiber.App
}

func NewAngelaWeb(handler *api.Handler, builder *page.Builder) *AngelaWeb {
	viewsEngine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{
		Views:       viewsEngine,
		ViewsLayout: "layout/default",
	})

	app.Static("/static", "./static")

	initRoutes(app)

	return &AngelaWeb{
		handler: handler,
		builder: builder,
		app:     app,
	}
}

func (a *AngelaWeb) Start(ctx context.Context) error {
	return a.app.Listen(":3000")
}

func (a *AngelaWeb) Stop(ctx context.Context) error {
	return a.app.ShutdownWithContext(ctx)
}

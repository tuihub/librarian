package angelaweb

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal"
)

type AngelaWeb struct {
	handler *internal.Handler
	app     *fiber.App
}

func NewAngelaWeb(handler *internal.Handler) *AngelaWeb {
	viewsEngine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{
		Views:       viewsEngine,
		ViewsLayout: "layout/default",
	})

	app.Static("/static", "./static")

	initRoutes(app)

	return &AngelaWeb{
		handler: handler,
		app:     app,
	}
}

func (a *AngelaWeb) Start(ctx context.Context) error {
	return a.app.Listen(":3000")
}

func (a *AngelaWeb) Stop(ctx context.Context) error {
	return a.app.ShutdownWithContext(ctx)
}

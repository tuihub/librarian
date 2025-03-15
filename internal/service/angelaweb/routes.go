package angelaweb

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
)

func initRoutes(app *fiber.App) {
	handler := &api.Handler{}

	app.Post("/api/login", handler.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
}

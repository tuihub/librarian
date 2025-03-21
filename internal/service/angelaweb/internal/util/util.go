package util

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/gofiber/fiber/v2"
)

func BizContext(c *fiber.Ctx) context.Context {
	return libauth.RawToContext(c.Context(), c.Cookies("access_token"))
}

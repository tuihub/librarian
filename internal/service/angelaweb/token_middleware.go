package angelaweb

import (
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/gofiber/fiber/v2"
)

func tokenMiddleware(auth *libauth.Auth, builder *page.Builder) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token")
		claims, err := libauth.FromString(accessToken, auth.KeyFunc(libauth.ClaimsTypeAccessToken))
		if err != nil {
			// Handle token issues - render error page for token expiration
			return builder.TokenExpired(c)
		}
		ctx := c.UserContext()
		ctx = libauth.RawToContext(ctx, c.Cookies("access_token"))
		ctx = jwt.NewContext(ctx, claims)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

package angelaweb

import (
	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/gofiber/fiber/v2"
)

func tokenMiddleware(auth *libauth.Auth) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token")
		claims, err := libauth.FromString(accessToken, auth.KeyFunc(libauth.ClaimsTypeAccessToken))
		if err != nil {
			return err
		}
		ctx := c.UserContext()
		ctx = libauth.RawToContext(ctx, c.Cookies("access_token"))
		ctx = jwt.NewContext(ctx, claims)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

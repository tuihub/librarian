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
			// Handle token issues - render error page for token expiration
			return c.Status(fiber.StatusUnauthorized).Render("error", fiber.Map{
				"Title":     "授权已过期",
				"Message":   "您的登录会话已过期，请重新登录",
				"ErrorType": "token_expired",
			})
		}
		ctx := c.UserContext()
		ctx = libauth.RawToContext(ctx, c.Cookies("access_token"))
		ctx = jwt.NewContext(ctx, claims)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

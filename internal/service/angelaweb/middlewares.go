package angelaweb

import (
	"time"

	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"
	"github.com/tuihub/librarian/internal/service/angelaweb/locales"

	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/samber/lo"
)

func (a *AngelaWeb) setupMiddlewares(settings *libapp.Settings) {
	a.app.Use(loggerMiddleware())

	a.app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        60, //nolint:mnd // no need
		Expiration: 1 * time.Minute,
	}))

	a.app.Use(fiberi18n.New(&fiberi18n.Config{
		RootPath:         "locales",
		FormatBundleFile: "toml",
		UnmarshalFunc:    toml.Unmarshal,
		Loader: &fiberi18n.EmbedLoader{
			FS: embedDirLocales,
		},
		DefaultLanguage: locales.DefaultLanguage(),
		AcceptLanguages: locales.SupportedLanguages(),
		LangHandler:     locales.LangHandler,
	}))

	a.app.Use(helmet.New())
	if settings.BuildType == libapp.BuildTypeDebug {
		a.app.Use(pprof.New())
	}
	if settings.EnablePanicRecovery {
		a.app.Use(recover.New())
	}
}

func tokenMiddleware(auth *libauth.Auth, builder *page.Builder) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token")
		claims, err := libauth.FromString(accessToken, auth.KeyFunc(libauth.ClaimsTypeAccessToken))
		if err != nil {
			// Handle token issues - render error page for token expiration
			return builder.TokenExpired(c)
		}

		// Check if token needs refresh
		var newAccessToken string
		var newClaims *libauth.Claims
		if claims.ExpiresAt != nil && time.Until(claims.ExpiresAt.Time) <= libtime.Hour { //nolint:nestif // no need
			// Generate new access token using the same user information from current claims
			newAccessToken, err = auth.GenerateToken(
				claims.UserID,
				claims.PorterID,
				libauth.ClaimsTypeAccessToken,
				claims.UserType,
				libtime.Hour,
			)
			if err != nil {
				// If refresh fails, log the error but continue with the current token
				log.Log(log.LevelWarn, "msg", "Failed to refresh access token", "error", err.Error())
			} else {
				// Set the new token in cookie
				c.Cookie(&fiber.Cookie{
					Name:     "access_token",
					Value:    newAccessToken,
					Expires:  time.Now().Add(libtime.Day),
					HTTPOnly: true,
					Secure:   true,
					SameSite: "strict",
					Path:     "/",
				})

				// Reparse the new token to get updated claims
				newClaims, err = libauth.FromString(accessToken, auth.KeyFunc(libauth.ClaimsTypeAccessToken))
				if err != nil {
					log.Log(log.LevelWarn, "msg", "Failed to parse refreshed token", "error", err.Error())
				} else {
					claims = newClaims
					log.Log(log.LevelInfo, "msg", "Token successfully refreshed", "user_id", claims.UserID)
				}
			}
		}

		ctx := c.UserContext()
		ctx = libauth.RawToContext(ctx, c.Cookies("access_token"))
		ctx = jwt.NewContext(ctx, claims)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

func loggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		kv := make(map[any]any)
		kv["method"] = c.Method()
		kv["path"] = c.Path()
		kv["ip"] = c.IP()
		startTime := time.Now()
		kv["time"] = startTime.Format(time.RFC3339)
		err := c.Next()
		kv["status"] = c.Response().StatusCode()
		kv["latency"] = time.Since(startTime).Milliseconds()
		log.Log(log.LevelInfo, lo.Flatten(lo.MapToSlice(kv, func(key any, value any) []any {
			return []any{key, value}
		}))...)
		return err
	}
}

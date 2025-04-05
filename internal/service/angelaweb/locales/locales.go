package locales

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/samber/lo"
	"golang.org/x/text/language"
)

func DefaultLanguage() language.Tag {
	return language.English
}

func SupportedLanguages() []language.Tag {
	return []language.Tag{
		language.English,
		language.Chinese,
	}
}

func LangHandler(c *fiber.Ctx, defaultLang string) string {
	if c == nil || c.Request() == nil {
		return defaultLang
	}
	var lang string
	supportedLanguages := lo.Map(SupportedLanguages(), func(tag language.Tag, _ int) string {
		return tag.String()
	})
	lang = utils.CopyString(c.Query("lang", ""))
	if lang != "" && lo.Contains(supportedLanguages, lang) {
		return lang
	}
	acceptLanguage := utils.CopyString(c.Get("Accept-Language", ""))
	if acceptLanguage != "" {
		langs := strings.Split(acceptLanguage, ",")
		if len(langs) > 0 {
			for _, l := range langs {
				parts := strings.Split(strings.TrimSpace(l), ";")
				if len(parts) > 0 && lo.Contains(supportedLanguages, parts[0]) {
					return parts[0]
				}
			}
		}
	}

	return defaultLang
}

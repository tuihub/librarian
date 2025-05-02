package api

import (
	"net/http"

	"github.com/tuihub/librarian/internal/model/modelangela"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UpdateServerInfo(c *fiber.Ctx) error {
	serverInfo := &modelangela.ServerInstanceSummary{
		Name:               c.FormValue("name"),
		Description:        c.FormValue("description"),
		WebsiteURL:         c.FormValue("website_url"),
		LogoImageURL:       c.FormValue("logo_image_url"),
		BackgroundImageURL: c.FormValue("background_image_url"),
	}

	err := h.a.SetServerInstanceSummary(c.UserContext(), serverInfo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiberi18n.MustLocalize(c, "UpdateServerInfoFailed"),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fiberi18n.MustLocalize(c, "UpdateServerInfoSuccess"),
	})
}

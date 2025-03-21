package api

import (
	"net/http"
	"time"

	"github.com/tuihub/librarian/internal/service/angelaweb/internal/util"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	accessToken, _, err := h.t.GetToken(util.BizContext(c), req.Username, req.Password, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Message})
	}

	c.Cookie(&fiber.Cookie{ //nolint:exhaustruct // no need
		Name:     "access_token",
		Value:    string(accessToken),
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "strict",
		Path:     "/",
	})

	return c.JSON(LoginResponse{})
}

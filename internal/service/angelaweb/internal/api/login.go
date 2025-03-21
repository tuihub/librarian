package api

import (
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	model2 "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	var user model.User
	if err := h.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	hashedPassword, err := h.auth.GeneratePassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error processing password"})
	}

	if hashedPassword != user.Password {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	token, err := h.auth.GenerateToken(model2.InternalID(user.ID), 0, libauth.ClaimsTypeAccessToken, model2.UserTypeAdmin, nil, 24*time.Hour)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error generating token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "strict",
		Path:     "/",
	})

	return c.JSON(model.LoginResponse{})
}

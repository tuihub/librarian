package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/model"
	"gorm.io/gorm"
)

type Handler struct {
	db   *gorm.DB
	auth *libauth.Auth
}

func NewHandler(db *gorm.DB, auth *libauth.Auth) *Handler {
	return &Handler{
		db:   db,
		auth: auth,
	}
}

func (h *Handler) GetDashboardStats(c *fiber.Ctx) error {
	var userCount int64
	if err := h.db.Model(&model.User{}).Count(&userCount).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching stats"})
	}
	return c.JSON(fiber.Map{"user_count": userCount})
}

package api

import (
	"net/http"

	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/model"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	t              *biztiphereth.Tiphereth
	userCountCache *libcache.Key[model.UserCount]
}

func NewHandler(
	t *biztiphereth.Tiphereth,
	userCountCache *libcache.Key[model.UserCount],
) *Handler {
	return &Handler{
		t:              t,
		userCountCache: userCountCache,
	}
}

func (h *Handler) GetDashboardStats(c *fiber.Ctx) error {
	userCount, err := h.userCountCache.Get(c.UserContext())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching stats"})
	}
	return c.JSON(fiber.Map{"user_count": userCount.Count})
}

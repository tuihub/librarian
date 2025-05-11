package api

import (
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/biz/bizangela"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	a              *bizangela.Angela
	t              *biztiphereth.Tiphereth
	g              *bizgebura.Gebura
	userCountCache *libcache.Key[model.UserCount]
}

func NewHandler(
	a *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	userCountCache *libcache.Key[model.UserCount],
) *Handler {
	return &Handler{
		a:              a,
		t:              t,
		g:              g,
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

func (h *Handler) ListSentinels(c *fiber.Ctx) error {
	pageNum, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize := 10
	sentinels, total, err := h.g.ListSentinels(c.UserContext(), &model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fiberi18n.MustLocalize(c, "ErrorFetchingSentinels"),
		})
	}
	return c.JSON(fiber.Map{
		"data":  sentinels,
		"total": total,
	})
}

func (h *Handler) CreateSentinel(c *fiber.Ctx) error {
	var sentinel modelgebura.Sentinel
	if err := c.BodyParser(&sentinel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := h.g.CreateSentinel(c.UserContext(), &sentinel)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fiberi18n.MustLocalize(c, "SentinelErrorCreating"),
		})
	}
	return c.JSON(sentinel)
}

func (h *Handler) GetSentinel(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	sentinel, err := h.g.GetSentinel(c.UserContext(), model.InternalID(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Sentinel not found",
		})
	}
	return c.JSON(sentinel)
}

func (h *Handler) UpdateSentinel(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	var sentinel modelgebura.Sentinel
	if err = c.BodyParser(&sentinel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	sentinel.ID = model.InternalID(id)
	err = h.g.UpdateSentinel(c.UserContext(), &sentinel)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fiberi18n.MustLocalize(c, "SentinelErrorUpdating"),
		})
	}
	return c.JSON(sentinel)
}

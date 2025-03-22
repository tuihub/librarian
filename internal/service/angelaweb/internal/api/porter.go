package api

import (
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) ListPorters(c *fiber.Ctx) error {
	// Parse pagination parameters from query string
	pageNum, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	porters, total, err := h.t.ListPorters(c.UserContext(), model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching porters"})
	}

	// Calculate pagination information
	totalPages := (int(total) + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	return c.JSON(fiber.Map{
		"porters": porters,
		"pagination": fiber.Map{
			"currentPage": pageNum,
			"totalPages":  totalPages,
			"totalItems":  total,
			"pageSize":    pageSize,
		},
	})
}

func (h *Handler) UpdatePorterStatus(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	var body struct {
		Status int `json:"status"`
	}

	if err = c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	var status model.UserStatus
	if body.Status == 1 {
		status = model.UserStatusActive
	} else {
		status = model.UserStatusBlocked
	}

	err = h.t.UpdatePorterStatus(c.UserContext(), model.InternalID(id), status)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

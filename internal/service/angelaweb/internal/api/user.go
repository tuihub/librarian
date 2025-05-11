package api

import (
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	// Parse pagination parameters from query string
	pageNum, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	users, total, err := h.t.ListUsers(c.UserContext(), model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	}, nil, nil)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching users"})
	}

	// Calculate pagination information
	totalPages := (int(total) + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	return c.JSON(fiber.Map{
		"users": users,
		"pagination": fiber.Map{
			"currentPage": pageNum,
			"totalPages":  totalPages,
			"totalItems":  total,
			"pageSize":    pageSize,
		},
	})
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}
	user, err := h.t.GetUser(c.UserContext(), lo.ToPtr(model.InternalID(id)))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}
	user.Status = model.UserStatusActive

	var err error
	res, err := h.t.CreateUser(c.UserContext(), &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(res)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}
	var user model.User
	if err = c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}
	user.ID = model.InternalID(id)

	err = h.t.UpdateUser(c.UserContext(), &user, "")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(user)
}

package api

import (
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	users, _, err := h.t.ListUsers(util.BizContext(c), model.Paging{
		PageSize: 1,
		PageNum:  20,
	}, nil, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching users"})
	}
	return c.JSON(users)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}
	user, err := h.t.GetUser(util.BizContext(c), lo.ToPtr(model.InternalID(id)))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user *model.User
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	var err error
	user, err = h.t.CreateUser(util.BizContext(c), user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(user)
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

	err = h.t.UpdateUser(util.BizContext(c), &user, "")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(user)
}

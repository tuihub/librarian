package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/model"
)

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	var users []model.User
	if err := h.db.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching users"})
	}
	return c.JSON(users)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	if err := h.db.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	// 加密密码
	hashedPassword, err := h.auth.GeneratePassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error processing password"})
	}
	user.Password = hashedPassword

	if err := h.db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating user"})
	}
	return c.JSON(user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var updates model.User
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	var user model.User
	if err := h.db.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	// 如果提供了新密码，则加密
	if updates.Password != "" {
		hashedPassword, err := h.auth.GeneratePassword(user.Password)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Error processing password"})
		}
		updates.Password = hashedPassword
	} else {
		updates.Password = user.Password
	}

	if err := h.db.Model(&user).Updates(updates).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error updating user"})
	}
	return c.JSON(user)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.db.Delete(&model.User{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error deleting user"})
	}
	return c.JSON(fiber.Map{"message": "User deleted"})
}

package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	var user model.User
	if err := h.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// 生成JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("your-secret-key")) // 在实际应用中使用环境变量
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Could not login"})
	}

	return c.JSON(model.LoginResponse{Token: t})
}

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error processing password"})
	}
	user.Password = string(hashedPassword)

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
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updates.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Error processing password"})
		}
		updates.Password = string(hashedPassword)
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

func (h *Handler) GetDashboardStats(c *fiber.Ctx) error {
	var userCount int64
	if err := h.db.Model(&model.User{}).Count(&userCount).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error fetching stats"})
	}
	return c.JSON(fiber.Map{"user_count": userCount})
}

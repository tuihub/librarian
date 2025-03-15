package page

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/model"
	"gorm.io/gorm"
)

type Builder struct {
	db *gorm.DB
}

func NewBuilder(db *gorm.DB) *Builder {
	return &Builder{db: db}
}

func (b *Builder) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "登录",
	})
}

func (b *Builder) Dashboard(c *fiber.Ctx) error {
	var userCount int64
	if err := b.db.Model(&model.User{}).Count(&userCount).Error; err != nil {
		return c.Status(500).SendString("Error fetching dashboard data")
	}

	return c.Render("dashboard", fiber.Map{
		"Title":     "Dashboard",
		"UserCount": userCount,
	})
}

func (b *Builder) UserList(c *fiber.Ctx) error {
	var users []model.User
	if err := b.db.Find(&users).Error; err != nil {
		return c.Status(500).SendString("Error fetching users")
	}

	return c.Render("user", fiber.Map{
		"Title": "用户管理",
		"Users": users,
	})
}

func (b *Builder) UserForm(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	var title string
	var action string
	var method string

	if id != "" {
		if err := b.db.First(&user, id).Error; err != nil {
			return c.Status(404).SendString("User not found")
		}
		title = "编辑用户"
		action = "/api/users/" + id
		method = "PUT"
	} else {
		title = "创建用户"
		action = "/api/users"
		method = "POST"
	}

	return c.Render("user_form", fiber.Map{
		"Title":  title,
		"User":   user,
		"Action": action,
		"Method": method,
	})
}

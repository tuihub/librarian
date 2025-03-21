package page

import (
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type Builder struct {
	t              *biztiphereth.Tiphereth
	userCountCache *libcache.Key[model.UserCount]
}

func NewBuilder(
	t *biztiphereth.Tiphereth,
	userCountCache *libcache.Key[model.UserCount],
) *Builder {
	return &Builder{
		t:              t,
		userCountCache: userCountCache,
	}
}

func (b *Builder) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "登录",
	})
}

func (b *Builder) Dashboard(c *fiber.Ctx) error {
	userCount, err := b.userCountCache.Get(c.UserContext())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching dashboard data")
	}

	return c.Render("dashboard", fiber.Map{
		"Title":     "Dashboard",
		"UserCount": userCount.Count,
	})
}

func (b *Builder) UserList(c *fiber.Ctx) error {
	users, _, err := b.t.ListUsers(c.UserContext(), model.Paging{
		PageNum:  1,
		PageSize: 20,
	}, nil, nil)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching users")
	}

	return c.Render("user", fiber.Map{
		"Title": "用户管理",
		"Users": users,
	})
}

func (b *Builder) UserForm(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var user *model.User
	var title string
	var action string
	var method string

	if idStr != "" {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid ID")
		}
		user, err = b.t.GetUser(c.UserContext(), lo.ToPtr(model.InternalID(id)))
		if err != nil {
			return c.Status(http.StatusNotFound).SendString("User not found")
		}
		title = "编辑用户"
		action = "/api/users/" + idStr
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

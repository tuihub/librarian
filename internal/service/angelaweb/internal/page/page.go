package page

import (
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/biz/bizangela"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/service/angelaweb/locales"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type Builder struct {
	a              *bizangela.Angela
	t              *biztiphereth.Tiphereth
	g              *bizgebura.Gebura
	configDigests  []*conf.ConfigDigest
	userCountCache *libcache.Key[model.UserCount]
}

func NewBuilder(
	a *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	configDigests []*conf.ConfigDigest,
	userCountCache *libcache.Key[model.UserCount],
) *Builder {
	return &Builder{
		a:              a,
		t:              t,
		g:              g,
		configDigests:  configDigests,
		userCountCache: userCountCache,
	}
}

// addCommonData adds common data used across templates.
func addCommonData(c *fiber.Ctx, data fiber.Map) fiber.Map {
	if data == nil {
		data = fiber.Map{}
	}
	// Add context for localization
	data["Ctx"] = c
	data["Lang"] = locales.LangHandler(c, locales.DefaultLanguage().String())
	return data
}

func (b *Builder) NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("error", addCommonData(c, fiber.Map{
		"ErrorType": "not_found",
	}))
}

func (b *Builder) TokenExpired(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).Render("error", addCommonData(c, fiber.Map{
		"ErrorType": "token_expired",
	}))
}

func (b *Builder) Login(c *fiber.Ctx) error {
	return c.Render("login", addCommonData(c, nil))
}

func (b *Builder) Dashboard(c *fiber.Ctx) error {
	userCount, err := b.userCountCache.Get(c.UserContext())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingData"))
	}

	return c.Render("dashboard", addCommonData(c, fiber.Map{
		"UserCount": userCount.Count,
	}))
}

func (b *Builder) UserList(c *fiber.Ctx) error {
	pageNum, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 10
	}

	pageSize := 10 // Users per page

	users, total, err := b.t.ListUsers(c.UserContext(), model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	}, nil, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching users")
	}
	// Calculate pagination information
	totalPages := (int(total) + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}
	return c.Render("user", addCommonData(c, fiber.Map{
		"Users": users,
		"Pagination": fiber.Map{
			"CurrentPage": pageNum,
			"TotalPages":  totalPages,
			"HasPrev":     pageNum > 1,
			"HasNext":     pageNum < totalPages,
			"PrevPage":    pageNum - 1,
			"NextPage":    pageNum + 1,
		},
	}))
}

func (b *Builder) UserForm(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var user *model.User
	var action, method string

	if idStr != "" {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid ID")
		}
		user, err = b.t.GetUser(c.UserContext(), lo.ToPtr(model.InternalID(id)))
		if err != nil {
			return c.Status(http.StatusNotFound).SendString("User not found")
		}
		action = "/api/users/" + idStr
		method = "PUT"
	} else {
		action = "/api/users"
		method = "POST"
		user = new(model.User)
	}
	return c.Render("user_form", addCommonData(c, fiber.Map{
		"User":   user,
		"Action": action,
		"Method": method,
	}))
}

func (b *Builder) PorterList(c *fiber.Ctx) error {
	pageNum, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize := 10 // Porters per page
	porters, total, err := b.t.ListPorters(c.UserContext(), model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingPorters"))
	}
	// Calculate pagination information
	totalPages := (int(total) + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}
	return c.Render("porter", addCommonData(c, fiber.Map{
		"Porters": porters,
		"Pagination": fiber.Map{
			"CurrentPage": pageNum,
			"TotalPages":  totalPages,
			"HasPrev":     pageNum > 1,
			"HasNext":     pageNum < totalPages,
			"PrevPage":    pageNum - 1,
			"NextPage":    pageNum + 1,
		},
	}))
}

func (b *Builder) ConfigList(c *fiber.Ctx) error {
	// Get server instance summary
	serverInfo, err := b.a.GetServerInstanceSummary(c.UserContext())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingServerInfo"))
	}

	return c.Render("config", addCommonData(c, fiber.Map{
		"Configs":    b.configDigests,
		"ServerInfo": serverInfo,
	}))
}

func (b *Builder) ServerInfoForm(c *fiber.Ctx) error {
	serverInfo, err := b.a.GetServerInstanceSummary(c.UserContext())
	if err != nil {
		serverInfo = new(modelangela.ServerInstanceSummary)
	}

	return c.Render("server_info_form", addCommonData(c, fiber.Map{
		"ServerInfo": serverInfo,
	}))
}

func (b *Builder) SentinelList(c *fiber.Ctx) error {
	pageNum, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize := 10 // Sentinels per page
	sentinels, total, err := b.g.ListSentinels(c.UserContext(), &model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingSentinels"))
	}
	// Calculate pagination information
	totalPages := (int(total) + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}
	return c.Render("sentinel", addCommonData(c, fiber.Map{
		"Sentinels": sentinels,
		"Pagination": fiber.Map{
			"CurrentPage": pageNum,
			"TotalPages":  totalPages,
			"HasPrev":     pageNum > 1,
			"HasNext":     pageNum < totalPages,
			"PrevPage":    pageNum - 1,
			"NextPage":    pageNum + 1,
		},
	}))
}

func (b *Builder) SentinelForm(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var sentinel *modelgebura.Sentinel
	var action, method string

	if idStr != "" {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid ID")
		}
		sentinel, err = b.g.GetSentinel(c.UserContext(), model.InternalID(id))
		if err != nil {
			return c.Status(http.StatusNotFound).SendString("Sentinel not found")
		}
		action = "/api/sentinels/" + idStr
		method = "PUT"
	} else {
		action = "/api/sentinels"
		method = "POST"
		sentinel = new(modelgebura.Sentinel)
	}
	return c.Render("sentinel_form", addCommonData(c, fiber.Map{
		"Sentinel": sentinel,
		"Action":   action,
		"Method":   method,
	}))
}

func (b *Builder) SentinelDetail(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	sentinel, err := b.g.GetSentinel(c.UserContext(), model.InternalID(id))
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("Sentinel not found")
	}

	pageNum, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize := 10 // Sessions per page

	sessions, total, err := b.g.ListSentinelSessions(c.UserContext(), &model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	}, sentinel.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingData"))
	}

	// Calculate pagination information
	totalPages := (int(total) + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	return c.Render("sentinel_detail", addCommonData(c, fiber.Map{
		"Sentinel": sentinel,
		"Sessions": sessions,
		"Pagination": fiber.Map{
			"CurrentPage": pageNum,
			"TotalPages":  totalPages,
			"HasPrev":     pageNum > 1,
			"HasNext":     pageNum < totalPages,
			"PrevPage":    pageNum - 1,
			"NextPage":    pageNum + 1,
		},
	}))
}

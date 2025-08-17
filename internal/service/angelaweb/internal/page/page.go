package page

import (
	"net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/biz/bizangela"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/service/angelaweb/locales"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

type Builder struct {
	app            *libapp.Settings
	a              *bizangela.Angela
	t              *biztiphereth.Tiphereth
	g              *bizgebura.Gebura
	configDigests  []*conf.ConfigDigest
	userCountCache *libcache.Key[model.UserCount]
}

func NewBuilder(
	app *libapp.Settings,
	a *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	configDigests []*conf.ConfigDigest,
	userCountCache *libcache.Key[model.UserCount],
) *Builder {
	return &Builder{
		app:            app,
		a:              a,
		t:              t,
		g:              g,
		configDigests:  configDigests,
		userCountCache: userCountCache,
	}
}

const (
	DefaultPageSize = 10
	DefaultPageNum  = 1
)

func getPaginationParams(c *fiber.Ctx) (int, int) {
	pageNum, err := strconv.Atoi(c.Query("page", strconv.Itoa(DefaultPageNum)))
	if err != nil || pageNum < 1 {
		pageNum = DefaultPageNum
	}
	return pageNum, DefaultPageSize
}

func parsePaginationData(pageNum, pageSize, total int) fiber.Map {
	totalPages := (total + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}
	return fiber.Map{
		"CurrentPage": pageNum,
		"TotalPages":  totalPages,
		"HasPrev":     pageNum > 1,
		"HasNext":     pageNum < totalPages,
		"PrevPage":    pageNum - 1,
		"NextPage":    pageNum + 1,
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
	// Add current path
	data["Path"] = c.Path()
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
	pageNum, pageSize := getPaginationParams(c)

	users, total, err := b.t.ListUsers(c.UserContext(), model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	}, nil, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching users")
	}

	return c.Render("user", addCommonData(c, fiber.Map{
		"Users":      users,
		"Pagination": parsePaginationData(pageNum, pageSize, int(total)),
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
	pageNum, pageSize := getPaginationParams(c)

	porters, total, err := b.t.ListPorters(c.UserContext(), model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingPorters"))
	}
	return c.Render("porter", addCommonData(c, fiber.Map{
		"Porters":    porters,
		"Pagination": parsePaginationData(pageNum, pageSize, int(total)),
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
	pageNum, pageSize := getPaginationParams(c)

	sentinels, total, err := b.g.ListSentinels(c.UserContext(), &model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingSentinels"))
	}

	return c.Render("sentinel", addCommonData(c, fiber.Map{
		"Sentinels":  sentinels,
		"Pagination": parsePaginationData(pageNum, pageSize, int(total)),
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
		var bizErr *errors.Error
		sentinel, bizErr = b.g.GetSentinel(c.UserContext(), model.InternalID(id))
		if bizErr != nil {
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

	var bizErr *errors.Error
	sentinel, bizErr := b.g.GetSentinel(c.UserContext(), model.InternalID(id))
	if bizErr != nil {
		return c.Status(http.StatusNotFound).SendString("Sentinel not found")
	}

	pageNum, pageSize := getPaginationParams(c)

	sessions, total, bizErr := b.g.ListSentinelSessions(c.UserContext(), &model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	}, sentinel.ID)
	if bizErr != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingData"))
	}

	return c.Render("sentinel_detail", addCommonData(c, fiber.Map{
		"Sentinel":   sentinel,
		"Sessions":   sessions,
		"Pagination": parsePaginationData(pageNum, pageSize, int(total)),
	}))
}

func (b *Builder) StoreAppList(c *fiber.Ctx) error {
	pageNum, pageSize := getPaginationParams(c)

	storeApps, total, err := b.g.ListStoreApps(c.UserContext(), &model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fiberi18n.MustLocalize(c, "ErrorFetchingStoreApps"))
	}

	return c.Render("store_app", addCommonData(c, fiber.Map{
		"StoreApps":  storeApps,
		"Pagination": parsePaginationData(pageNum, pageSize, int(total)),
	}))
}

func (b *Builder) StoreAppBinaryList(c *fiber.Ctx) error {
	pageNum, pageSize := getPaginationParams(c)

	storeAppBinaries, total, err := b.g.ListStoreAppBinaries(c.UserContext(), &model.Paging{
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	}, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			SendString(fiberi18n.MustLocalize(c, "ErrorFetchingStoreAppBinaries"))
	}

	return c.Render("store_app_binary", addCommonData(c, fiber.Map{
		"StoreAppBinaries": storeAppBinaries,
		"Pagination":       parsePaginationData(pageNum, pageSize, int(total)),
	}))
}

// Monitor 显示监控页面.
func (b *Builder) Monitor(c *fiber.Ctx) error {
	return c.Render("monitor", addCommonData(c, fiber.Map{
		"Debug": b.app.BuildType == libapp.BuildTypeDebug,
	}))
}

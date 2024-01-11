package modeltiphereth

import (
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
)

type User struct {
	ID       model.InternalID
	UserName string
	PassWord string
	Type     libauth.UserType
	Status   UserStatus
}

type UserStatus int

const (
	UserStatusUnspecified UserStatus = iota
	UserStatusActive
	UserStatusBlocked
)

type AccessToken string
type RefreshToken string

type Account struct {
	ID                model.InternalID
	Platform          string
	PlatformAccountID string
	Name              string
	ProfileURL        string
	AvatarURL         string
	LatestUpdateTime  time.Time
}

type PullAccountInfo struct {
	ID                model.InternalID
	Platform          string
	PlatformAccountID string
}

type UserSession struct {
	ID           model.InternalID
	UserID       model.InternalID
	RefreshToken string
	DeviceInfo   *DeviceInfo
	CreateAt     time.Time
	ExpireAt     time.Time
}

type DeviceInfo struct {
	ID                      model.InternalID
	DeviceModel             string
	SystemVersion           string
	ClientName              string
	ClientSourceCodeAddress string
	ClientVersion           string
}

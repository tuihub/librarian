package modeltiphereth

import (
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/model"
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
	Platform          AccountPlatform
	PlatformAccountID string
	Name              string
	ProfileURL        string
	AvatarURL         string
	LatestUpdateTime  time.Time
}

type AccountPlatform int

const (
	AccountPlatformUnspecified AccountPlatform = iota
	AccountPlatformSteam
)

type PullAccountInfo struct {
	ID                model.InternalID
	Platform          AccountPlatform
	PlatformAccountID string
}

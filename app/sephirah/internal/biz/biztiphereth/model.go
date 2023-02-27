package biztiphereth

import (
	"github.com/tuihub/librarian/internal/lib/libauth"
)

type User struct {
	InternalID int64
	UserName   string
	PassWord   string
	Type       libauth.UserType
	Status     UserStatus
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
	InternalID        int64
	Platform          AccountPlatform
	PlatformAccountID string
	Name              string
	ProfileURL        string
	AvatarURL         string
}

type AccountPlatform int

const (
	AccountPlatformUnspecified AccountPlatform = iota
	AccountPlatformSteam
)

type PullAccountInfo struct {
	InternalID        int64
	Platform          AccountPlatform
	PlatformAccountID string
}

package model

import (
	"time"
)

type User struct {
	ID       InternalID
	Username string
	Password string
	Type     UserType
	Status   UserStatus
}

type UserStatus int

const (
	UserStatusUnspecified UserStatus = iota
	UserStatusActive
	UserStatusBlocked
)

type UserType int

const (
	UserTypeUnspecified UserType = iota
	UserTypeAdmin
	UserTypeNormal
	UserTypeSentinel
	UserTypePorter
)

type AccessToken string
type RefreshToken string

type Account struct {
	ID                InternalID
	Platform          string
	PlatformAccountID string
	Name              string
	ProfileURL        string
	AvatarURL         string
	LatestUpdateTime  time.Time
}

type PullAccountInfo struct {
	ID                InternalID
	Config            *FeatureRequest
	Platform          string
	PlatformAccountID string
}

type Session struct {
	ID           InternalID
	UserID       InternalID
	RefreshToken string
	Device       *Device
	CreateAt     time.Time
	ExpireAt     time.Time
}

type Device struct {
	ID                      InternalID
	DeviceName              string
	SystemType              SystemType
	SystemVersion           string
	ClientName              string
	ClientSourceCodeAddress string
	ClientVersion           string
}

type SystemType int

const (
	SystemTypeUnspecified SystemType = iota
	SystemTypeAndroid
	SystemTypeIOS
	SystemTypeWindows
	SystemTypeMacOS
	SystemTypeLinux
	SystemTypeWeb
)

type CaptchaQue struct {
	ID    string
	Image []byte
}

type CaptchaAns struct {
	ID    string
	Value string
}

type UserCount struct {
	Count int
}

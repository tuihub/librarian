package model

import (
	"database/sql/driver"
	"errors"
	"time"
)

type User struct {
	ID        InternalID `gorm:"primaryKey;autoIncrement:false"`
	Username  string     `gorm:"uniqueIndex"`
	Password  string
	Type      UserType
	Status    UserStatus
	CreatorID InternalID
	CreatedAt time.Time
	UpdatedAt time.Time
	Sessions  []Session `gorm:"foreignKey:UserID"`
	Account   []Account `gorm:"foreignKey:BoundUserID"`
	// Relations to other packages removed to avoid circular dependencies
}

func (User) TableName() string {
	return "users"
}

type UserStatus int

const (
	UserStatusUnspecified UserStatus = iota
	UserStatusActive
	UserStatusBlocked
)

func (s UserStatus) Value() (driver.Value, error) {
	switch s {
	case UserStatusActive:
		return "active", nil
	case UserStatusBlocked:
		return "blocked", nil
	default:
		return "", nil
	}
}

func (s *UserStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for UserStatus")
	}
	switch v {
	case "active":
		*s = UserStatusActive
	case "blocked":
		*s = UserStatusBlocked
	default:
		*s = UserStatusUnspecified
	}
	return nil
}

type UserType int

const (
	UserTypeUnspecified UserType = iota
	UserTypeAdmin
	UserTypeNormal
	UserTypeSentinel
	UserTypePorter
)

func (t UserType) Value() (driver.Value, error) {
	switch t {
	case UserTypeAdmin:
		return "admin", nil
	case UserTypeNormal:
		return "normal", nil
	case UserTypeSentinel:
		return "sentinel", nil
	case UserTypePorter:
		return "porter", nil
	default:
		return "", nil
	}
}

func (t *UserType) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for UserType")
	}
	switch v {
	case "admin":
		*t = UserTypeAdmin
	case "normal":
		*t = UserTypeNormal
	case "sentinel":
		*t = UserTypeSentinel
	case "porter":
		*t = UserTypePorter
	default:
		*t = UserTypeUnspecified
	}
	return nil
}

type AccessToken string
type RefreshToken string

type Account struct {
	ID                InternalID `gorm:"primaryKey;autoIncrement:false"`
	Platform          string     `gorm:"index:idx_account_platform_id,priority:1"`
	PlatformAccountID string     `gorm:"index:idx_account_platform_id,priority:2"`
	Name              string
	ProfileURL        string
	AvatarURL         string
	LatestUpdateTime  time.Time
	BoundUserID       InternalID `gorm:"index"`
	BoundUser         *User      `gorm:"foreignKey:BoundUserID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (Account) TableName() string {
	return "accounts"
}

type PullAccountInfoConfig struct {
	Platform          string `json:"platform"`
	PlatformAccountID string `json:"platform_account_id"`
}

type PullAccountInfo struct {
	ID     InternalID
	Config *FeatureRequest
}

type Session struct {
	ID           InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID       InternalID `gorm:"index:idx_session_user_id_device_id,priority:1"`
	DeviceID     InternalID `gorm:"index:idx_session_user_id_device_id,priority:2"`
	RefreshToken string     `gorm:"uniqueIndex"`
	Device       *Device    `gorm:"foreignKey:DeviceID"`
	User         *User      `gorm:"foreignKey:UserID"`
	CreateAt     time.Time  `gorm:"column:created_at"` // Map to created_at to match GORM default
	ExpireAt     time.Time
	UpdatedAt    time.Time
}

func (Session) TableName() string {
	return "sessions"
}

type Device struct {
	ID                      InternalID `gorm:"primaryKey;autoIncrement:false"`
	DeviceName              string
	SystemType              SystemType
	SystemVersion           string
	ClientName              string
	ClientSourceCodeAddress string
	ClientVersion           string
	ClientLocalID           string
	CreatedAt               time.Time
	UpdatedAt               time.Time
	Sessions                []Session `gorm:"foreignKey:DeviceID"`
	// App relation removed
}

func (Device) TableName() string {
	return "devices"
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

func (t SystemType) Value() (driver.Value, error) {
	switch t {
	case SystemTypeIOS:
		return "ios", nil
	case SystemTypeAndroid:
		return "android", nil
	case SystemTypeWeb:
		return "web", nil
	case SystemTypeWindows:
		return "windows", nil
	case SystemTypeMacOS:
		return "macos", nil
	case SystemTypeLinux:
		return "linux", nil
	default:
		return "", nil
	}
}

func (t *SystemType) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for SystemType")
	}
	switch v {
	case "ios":
		*t = SystemTypeIOS
	case "android":
		*t = SystemTypeAndroid
	case "web":
		*t = SystemTypeWeb
	case "windows":
		*t = SystemTypeWindows
	case "macos":
		*t = SystemTypeMacOS
	case "linux":
		*t = SystemTypeLinux
	default:
		*t = SystemTypeUnspecified
	}
	return nil
}

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

type Tag struct {
	ID          InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserTag     InternalID `gorm:"index"` // UserID
	Name        string
	Description string
	Public      bool
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *User `gorm:"foreignKey:UserTag"`
}

func (Tag) TableName() string {
	return "tags"
}

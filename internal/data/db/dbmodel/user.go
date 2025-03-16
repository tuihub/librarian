package dbmodel

type UserStatus string

const (
	UserStatusUnspecified UserStatus = "unspecified"
	UserStatusActive      UserStatus = "active"
	UserStatusBlocked     UserStatus = "blocked"
)

type UserType string

const (
	UserTypeUnspecified UserType = "unspecified"
	UserTypeAdmin       UserType = "admin"
	UserTypeNormal      UserType = "normal"
	UserTypeSentinel    UserType = "sentinel"
	UserTypePorter      UserType = "porter"
)

type User struct {
	Model
	Username string
	Password string
	Status   UserStatus
	Type     UserType
}

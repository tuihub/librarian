package dbmodel

import "time"

type Session struct {
	Model
	UserID   ID
	DeviceID ID
	Token    string
	ExpireAt time.Time
}

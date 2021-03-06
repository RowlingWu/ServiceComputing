package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	Uid        int `orm:"id,auto-inc"` //语义标签
	Username   string
	Departname string
	Created    *time.Time
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.Username) == 0 {
		panic("UserName shold not null!")
	}
	if u.Created == nil {
		t := time.Now()
		u.Created = &t
	}
	return &u
}

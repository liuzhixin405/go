package model

import "time"

type User struct {
	UserName    string
	PassWord    string
	Email       string
	Phone       int64
	CreatedTime time.Time
}

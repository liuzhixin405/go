package model

import "time"

type Dto struct {
	Id         string `gorm:"primary_key"`
	CreateTime *time.Time
	CreatorId  string
	Deleted    bool
}

package model


type Dto struct {
	
	Id         string `gorm:"primary_key"`
	CreateTime int64
	CreatorId  string
	Deleted    bool
}

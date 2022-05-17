package src

import "time"

type Topics struct {
	TopicID         int       `json:"id" gorm:"primary_key"`
	TopicTitle      string    `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string    `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string    `json:"ip" binding:"ipv4"`
	TopicScore      int       `json:"score" binding:"omitempty,gte=5"`
	TopicUrl        string    `json:"url" binding:"omitempty,topicurlnew"`
	TopicDate       time.Time `json:"date"`
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required" `
	PageSize int    `json:"pagesize" form:"pagesize"`
}

type TopicArray struct {
	TopicList     []Topics `json:"topics" binding:"gt=0,lt=3,dive"`
	TopicListSize int      `json:"size"`
}

type Model struct {
	ID uint `gorm:"primaryKey"`
}

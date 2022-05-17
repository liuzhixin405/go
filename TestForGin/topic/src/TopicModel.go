package src

type Topic struct {
	TopicID         int    `json:"id" gorm:"primary_key"`
	TopicTitle      string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string `json:"ip" binding:"ipv4"`
	TopicScore      int    `json:"score" binding:"omitempty,gte=5"`
	TopicUrl        string `json:"url" binding:"omitempty,topicurlnew"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title, "testtest", "123.123.123.123", 6, "url"}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required" `
	PageSize int    `json:"pagesize" form:"pagesize"`
}

type Topics struct {
	TopicList     []Topic `json:"topics" binding:"gt=0,lt=3,dive"`
	TopicListSize int     `json:"size"`
}

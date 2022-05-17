package src

import "github.com/gin-gonic/gin"

func MustLogin() gin.HandlerFunc {
	//登录
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); status != true {
			c.JSON(401, gin.H{
				"status": 401,
				"msg":    "请登录",
			})
			c.Abort()
			return
		}
	}
}

func GetTopicDetail(c *gin.Context) {
	c.JSON(200, CreateTopic(101, "帖子标题"))
}

func GetTopicList(c *gin.Context) {
	query := TopicQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(400, "参数错误: %s", err.Error())
	} else {
		c.JSON(200, query)
	}
}

func NewTopic(c *gin.Context) {
	topic := Topic{}
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(400, "参数错误:%s", err.Error())
	} else {
		c.JSON(200, topic)
	}

}

func NewTopics(c *gin.Context) {
	topics := Topics{}
	err := c.BindJSON(&topics)
	if err != nil {
		c.String(400, "参数错误:%s", err.Error())
	} else {
		c.JSON(200, topics)
	}

}

func UpdateTopic(c *gin.Context) {
	c.String(200, "更新帖子")
}

func DeleteTopic(c *gin.Context) {
	c.String(200, "新删除帖子")
}

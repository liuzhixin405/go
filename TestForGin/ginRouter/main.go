package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/v1/topics/:topic_id", func(c *gin.Context) {
		c.String(200, "获取帖子ID为：%s", c.Param("topic_id"))
	}) //curl http://localhost:8080/v1/topics/345

	//对的设计
	router.GET("/v1/topics", func(c *gin.Context) {
		if c.Query("username") == "" {
			c.String(200, "获取帖子列表...") //curl http://localhost:8080/v1/topics
		} else {
			c.String(200, "获取用户名=%s的帖子列表", c.Query("username")) //curl http://localhost:8080/v1/topics?username=lzx
		}

	})

	router.Run(":8080")
}

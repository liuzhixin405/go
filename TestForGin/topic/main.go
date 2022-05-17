package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	. "topic.lzxtest.com/src"
)

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurlnew", TopicUrlNew)
	}
	v1 := router.Group("/v1/topics")
	{
		v1.GET("/:topic_id", GetTopicDetail) //curl http://localhost:8080/v1/topics/345

		//对的设计
		// v1.GET("", func(c *gin.Context) {
		// 	if c.Query("username") == "" {
		// 		c.String(200, "获取帖子列表...") //curl http://localhost:8080/v1/topics
		// 	} else {
		// 		c.String(200, "获取用户名=%s的帖子列表", c.Query("username")) //curl http://localhost:8080/v1/topics?username=lzx
		// 	}
		// })

		v1.GET("", GetTopicList) // (cmd)(curl "http://localhost:8080/v1/topics?username=lzx&page=1&pagesize=10" ) cmd下不带双引号无法识别&后面的参数

		v1.Use(MustLogin())
		{
			v1.POST("", NewTopic)                // (cmd下有效，powershell下无效) curl "http://localhost:8080/v1/topics?token=123" -X POST  -d "{\"title\":\"test\",\"stitle\":\"testshort\",\"ip\":\"127.0.0.3\",\"score\":6,\"url\":\"abcd_nihao\"}" -H "Content-Type:application/json;charset-utf-8" -H "ACcept:application/json"
			v1.DELETE("/:topic_id", DeleteTopic) //curl -X DELETE http://localhost:8080/v1/topics/345?token=123
		}
	}
	v2 := router.Group("/v1/mtopics")
	{
		v2.Use(MustLogin())
		{
			v2.POST("", NewTopics)
		}
		/*
			curl "http://localhost:8080/v1/mtopics?token=123" -X POST  -d  "{\"topics\":[{\"title\":\"test\",\"stitle\":\"testshort\",\"ip\":\"127.0.0.3\",\"score\":6,\"url\":\"abcd_nihao\"}, {\"title\":\"test\",\"stitle\":\"testshort\",\"ip\":\"127.0.0.3\",\"score\":6,\"url\":\"abcd_nihao\"} ],\"size\":2}" -H "Content-Type:application/json;charset-utf-8" -H "ACcept:application/json"
		*/
	}

	router.Run(":8080")
}

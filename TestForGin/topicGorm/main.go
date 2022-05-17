package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	. "topicGorm.eiza.com/src"
)

var DB *gorm.DB

//var err error  //坑 不管是全局还是gorm都会报错

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

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

	//router.Run(":8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go (func() {
		InitDB()
	})()

	go (func() {
		err := server.ListenAndServe()
		if err != nil {
			//log.Fatal("服务其启动失败")
			ShutDownServer(err)
		}
	})()

	ServerNotify()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {

		log.Fatal("服务器关闭")
	}
	log.Println("服务器有序退出")
}

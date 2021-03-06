package src

import (
	"time"

	"github.com/gin-gonic/gin"
)

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

// func GetTopicDetail(c *gin.Context) {

// 	tid := c.Param("topic_id")
// 	topics := Topics{}

// 	conn := RedisDefaultPool.Get()
// 	redisKey := "topic_" + tid
// 	defer conn.Close()

// 	ret, err := redis.Bytes(conn.Do("get", redisKey))
// 	if err != nil {
// 		DBHelper.Find(&topics, tid)
// 		retData, _ := json.Marshal(topics)
// 		if topics.TopicID == 0 {
// 			conn.Do("setex", redisKey, 20, retData) //空数据防止缓存穿透
// 		} else {
// 			conn.Do("setex", redisKey, 50, retData) //正常数据
// 		}

// 		log.Println("从数据库返回")
// 	} else {
// 		json.Unmarshal(ret, &topics)
// 		log.Println("从redis返回")

// 	}
// 	c.JSON(200, topics)

// }

func GetTopicDetail(c *gin.Context) {

	tid := c.Param("topic_id")
	topics := Topics{}
	DBHelper.Find(&topics, tid)
	c.Set("dbResult", topics)
	// conn := RedisDefaultPool.Get()
	// redisKey := "topic_" + tid
	// defer conn.Close()

	// ret, err := redis.Bytes(conn.Do("get", redisKey))
	// if err != nil {
	// 	DBHelper.Find(&topics, tid)
	// 	retData, _ := json.Marshal(topics)
	// 	if topics.TopicID == 0 {
	// 		conn.Do("setex", redisKey, 20, retData) //空数据防止缓存穿透
	// 	} else {
	// 		conn.Do("setex", redisKey, 50, retData) //正常数据
	// 	}

	// 	log.Println("从数据库返回")
	// } else {
	// 	json.Unmarshal(ret, &topics)
	// 	log.Println("从redis返回")

	// }
	// c.JSON(200, topics)

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
	topic := Topics{}
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(400, "参数错误:%s", err.Error())
	} else {
		c.JSON(200, topic)
	}

}

func NewTopics(c *gin.Context) {
	// topics := Topics{}
	// err := c.BindJSON(&topics)
	// if err != nil {
	// 	c.String(400, "参数错误:%s", err.Error())
	// } else {
	// 	c.JSON(200, topics)
	// }

	// if err != nil {

	topics := Topics{
		TopicTitle:      "TopicTitle",
		TopicShortTitle: "TopicShortTitle",
		UserIP:          "127.0.0.1",
		TopicScore:      0,
		TopicUrl:        "testUrl",
		TopicDate:       time.Now(),
	}
	row := DBHelper.Create(&topics).RowsAffected
	c.String(200, "新增帖子成功,id:%d , 行数：%d", topics.TopicID, row)

	// } else {
	// 	fmt.Println(err.Error())
	// }

}

func UpdateTopic(c *gin.Context) {
	c.String(200, "更新帖子")
}

func DeleteTopic(c *gin.Context) {
	c.String(200, "新删除帖子")
}

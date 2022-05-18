package src

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

//缓存装饰器
func CacheDecorator(h gin.HandlerFunc, param string, redKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//redis if

		getID := context.Param(param)
		redisKey := fmt.Sprintf(redKeyPattern, getID)
		conn := RedisDefaultPool.Get()
		defer conn.Close()
		ret, err := redis.Bytes(conn.Do("get", redisKey))
		if err != nil {
			h(context)
			dbResult, exists := context.Get("dbResult")
			if !exists {
				dbResult = empty
			}
			retData, _ := json.Marshal(dbResult)
			conn.Do("setex", redisKey, 20, retData)
		} else {

			json.Unmarshal(ret, &empty)
			log.Println("从redis返回")
		}
		context.JSON(200, empty)
	}
}

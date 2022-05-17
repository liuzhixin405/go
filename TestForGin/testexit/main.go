package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	//1. Create a channel to listen for an interrupt or terminate signal from the OS.
	// count := 0
	// for {
	// 	fmt.Println("执行", count)
	// 	count++
	// 	time.Sleep(time.Second * 1)
	// }

	go func() {
		count := 0
		for {
			fmt.Println("执行", count)
			count++
			time.Sleep(time.Second * 1)
		}

	}() //这里的go就是开启一个协程,不是协程interrupt不了
	//2. Create a channel to listen for an interrupt or terminate signal from the OS.
	// 1.2.调换效果出其不意
	c := make(chan os.Signal)

	go func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		select {
		case <-ctx.Done():
			c <- os.Interrupt
		}
	}()

	signal.Notify(c)
	s := <-c
	fmt.Println(s)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // 慢 SQL 阈值
	// 		LogLevel:                  logger.Silent, // 日志级别
	// 		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  false,         // 禁用彩色打印
	// 	},
	// )

	// dsn := "root1:1230@/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// DBHelper, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: newLogger,
	// })
	// DBHelper.DB()

	// if err != nil {
	// 	fmt.Println(err)
	// }

}

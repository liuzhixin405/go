package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New(cron.WithSeconds())
	spec := "*/5 * * * * *"
	c.AddFunc(spec, func() {

		log.Printf(" print time = %d\n", time.Now().Unix())
	})
	c.AddFunc("*/1 * * * * *", func() { // 可以随时添加多个定时任务
		fmt.Println("222")
	})

	defer c.Stop()
	c.Start()
	select {} //阻塞主线程退出
}

// /go test -v main_test.go

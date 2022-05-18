package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {

	defer wg.Done()
FORLOOP: //标签 退出的位置  否则退出当前循环会进入下一次循环
	for {
		fmt.Println("执行中...")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait() //通知子goroutine退出
}

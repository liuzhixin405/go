package main

import (
	"fmt"
	"time"
)

func main() {

	msg := make(chan string)
	go send(msg)
	go receive(msg)
	time.Sleep(time.Second)
}

func send(msg chan string) {
	str := "hello world"
	fmt.Println("ready send msg:", str)
	msg <- str
}
func receive(msg chan string) {
	recMsg, ok := <-msg
	if !ok {
		panic("error")
	}
	fmt.Println("receive msg:", recMsg)
}

//现实中的错误写法
// package main

// import "fmt"

// func main() {

// 	msg := make(chan string)

// 	msg <- "hello world"

// 	recMsg, ok := <-msg
// 	if !ok {
// 		panic("err")
// 	}
// 	fmt.Println("get message:", recMsg)
// }

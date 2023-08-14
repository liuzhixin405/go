package main

import (
	"fmt"
)

type Message struct {
	message string
} //消息类

type Greeter struct {
	Message Message
} //打招呼结构体
type Event struct {
	Greeter Greeter
} //事件

func NewMessage(msg string) Message {
	return Message{
		message: msg,
	}
} //消息构造器
func (g Greeter) Greet() Message {
	return g.Message //问候的方法
}
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
} //招呼的结构体
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
} //事件构造器
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
} //时间的方法

/* func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()
}
*/ //old

func main() {
	event := InitializeEvent("hi")
	event.Start()
}

//  go run main.go wire_gen.go   否则运行报错.\main.go:46:11: undefined: InitializeEvent

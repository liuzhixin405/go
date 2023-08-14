package main

import "fmt"

type Message string //消息类

type Greeter struct {
	Message Message
} //打招呼结构体
type Event struct {
	Greeter Greeter
} //事件

func NewMessage() Message {
	return Message("hi here!")
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

func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}

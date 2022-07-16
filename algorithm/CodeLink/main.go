package main

import "fmt"

//链式栈
func main_linkstack() {

	mystack := NewStack()
	for i := 0; i < 100000000; i++ {
		mystack.Push(i)
	}
	for data := mystack.Pop(); data != nil; data = mystack.Pop() {
		fmt.Println(data)
	}
}

//链式队列
func main() {
	myq := NewLinkQueue()
	for i := 0; i < 1000000; i++ {
		myq.Enqueue(i)
	}
	for data := myq.Dequeue(); data != nil; data = myq.Dequeue() {
		fmt.Println(data)
	}
}

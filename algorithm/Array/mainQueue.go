package main

import (
	"fmt"
	"main/criclequeue"
	//"main/queue"
)

// func main_bak() {
// 	myqueue := queue.NewQueue()
// 	myqueue.EnQueue(1)
// 	myqueue.EnQueue(2)
// 	myqueue.EnQueue(3)
// 	myqueue.EnQueue(4)

// 	fmt.Println(myqueue.Dequeue())
// 	fmt.Println(myqueue.Dequeue())
// 	fmt.Println(myqueue.Dequeue())
// 	fmt.Println(myqueue.Dequeue())
// 	fmt.Println(myqueue.Dequeue())
// }

func main() {
	var myqueue criclequeue.CricleQueue
	criclequeue.InitQueue(&myqueue)

	criclequeue.EnQueue(&myqueue,1)
	criclequeue.EnQueue(&myqueue,2)
	criclequeue.EnQueue(&myqueue,3)
	criclequeue.EnQueue(&myqueue,4)

	fmt.Println(criclequeue.Dequeue(&myqueue))
	fmt.Println(criclequeue.Dequeue(&myqueue))
	fmt.Println(criclequeue.Dequeue(&myqueue))
	fmt.Println(criclequeue.Dequeue(&myqueue))
	fmt.Println(criclequeue.Dequeue(&myqueue))
}

package main

import (
	"fmt"
	"time"
)

// import (
// 	"sync"
// 	"time"
// 	"fmt"
// )

//one test
// func main() {
// 	go count("sheep")
// 	go count("fish")
// 	//fmt.Scanln() //Console.ReadLine()  防止结束go相当于task.Run
// 	//time.Sleep(time.Second * 5)   防止结束task.Run
// }

// func count(thing string) {
// 	for i := 1; true; i++ {
// 		fmt.Println(i, thing)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

//two test
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		count("sheep")
// 		wg.Done()
// 	}()
// 	wg.Wait() //未收到done之前阻塞主线程
// }

// func count(thing string) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i, thing)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

//three test

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan string)
// 	go count("sheep", c)
// 	for msg := range c {
// 		fmt.Println(msg)
// 	}

// }

// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }

//four
// func main() {
// 	c := make(chan string, 2) //缓冲
// 	c <- "hello"
// 	c <- "world"
// 	msg := <-c
// 	fmt.Println(msg)
// 	msg = <-c
// 	fmt.Println(msg)

// }

//deadlock!
// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		c1 <- "Every 500ms"
// 		time.Sleep(time.Millisecond * 500)
// 	}()
// 	go func() {
// 		for {
// 			c2 <- "Every two seconds"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()

// 	for {
// 		fmt.Println(<-c1)
// 		fmt.Println(<-c2)
// 	}
// }

//正解

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		c1 <- "Every 500ms"
// 		time.Sleep(time.Millisecond * 500)
// 	}()
// 	go func() {
// 		for {
// 			c2 <- "Every two seconds"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()

// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

func main() {
	startTime := time.Now()
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for count := 0; count < 10; count++ {
		go worker(jobs, results)
	}

	for i := 0; i < 45; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 45; j++ {
		fmt.Println(<-results)
	}
	endTime := time.Now()
	fmt.Println("channel run time:", endTime.Sub(startTime), "ms")
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

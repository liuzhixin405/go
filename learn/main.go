package main

import (
	"sync"
)

func add(i *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	wg.Done()
	*i++
}

func main() {
	i := 0
	count := 100000000
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(count)
	for j := 0; j < count; j++ {
		go add(&i, &mu, &wg)
	}
	wg.Wait()
	println(i)
}

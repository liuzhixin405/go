package main

import (
	"fmt"
	"sync"
)

var store sync.Map

func mainBak() {

	store.Store("key", "hello cache")
	writer, found := store.Load("key1")
	if found {
		fmt.Println(writer.(string))
	} else {
		//赋值
	}

	store.LoadOrStore("key2", "hello cache2")
	writer, _ = store.Load("key2")
	fmt.Println(writer.(string))

	writer, _ = store.Load("key2")
	fmt.Println(writer.(string))
}

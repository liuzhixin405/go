package SINGLETON

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingltonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Singleton Object")
		singleInstance = new(Singleton)

	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingltonObj()
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

//go test -run TestGetSingletonObj -v      单元测试运行

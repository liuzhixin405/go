package main

import (
	"fmt"
	"strinf/myinterface/myimplement"
)

func main() {
	fmt.Println("go start")

	myimplement.SharedStore().Test()
}

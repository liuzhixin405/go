package main

import (
	"fmt"

	"github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}
type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func main() {
	s := Staff{}
	inj := inject.New()
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)

	inj.Apply(&s)
	fmt.Printf("s = %v\n", s)
}

// func Format(name string, company S1, level S2, age int) {
// 	fmt.Printf("name= %s, company=%s,level=%s, age=%d\n", name, company, level, age)
// }

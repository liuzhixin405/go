package myimplement

import (
	"fmt"
	"strinf/myinterface"
)

var sto myinterface.Store

type Store struct {
}

func SharedStore() myinterface.Store {
	sto = NewStore()
	return sto
}

func NewStore() *Store {
	return &Store{}
}
func (s *Store) Test() {
	fmt.Println("test 实现了 myi.store接口到mysql.store的实现")
}

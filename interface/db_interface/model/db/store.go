package db

import (
	"fmt"
	"sync"
	"testimplement/model"
)

var store model.Store
var storeOnce sync.Once
var msg string

type Store struct {
	message string
}

func SharedStore() model.Store {
	storeOnce.Do(func() {
		err := initThis()
		if err != nil {
			panic(err)
		}
		store = NewStore(msg)
	})
	return store
}

func NewStore(message string) *Store {
	return &Store{
		message: message,
	}
}

func initThis() error {
	fmt.Println("init success")
	return nil
}

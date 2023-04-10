package service

import (
	"fmt"
	"login/model/psql"
)

type Service struct{}

func (l Service) Login(userName string, password string) bool {

	tx, err := psql.NewStore()
	if err != nil {
		panic(err)
	}
	defer tx.Close()
	fmt.Println("isLogin start")
	isLogin, err := tx.UserLogin(userName, password)
	fmt.Println("isLogin=", isLogin)
	if err != nil {
		panic(err)
	}
	return isLogin
}

type LoginService interface {
	Login(userName string, password string) bool
}

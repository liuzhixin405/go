package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("添加用户测试:")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}

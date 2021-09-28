package model

import (
	"fmt"
	"web01_db/utils"
)

//User 结构体
type User struct {
	ID       int
	Username string
	Pwd      string
	Email    string
}

//添加User方法

func (user *User) AddUser() error {
	sqlStr := "insert into users(username,pwd,email) values (?,?,?)"
	intSmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}
	_, err2 := intSmt.Exec("admin11", "11111", "asd@qq.com")
	if err2 != nil {
		fmt.Println("插入出现异常", err2)
		return err2
	}
	return nil
}
func (user *User) AddUser2() error {
	sqlStr := "insert into users(username,pwd,email) values (?,?,?)"

	_, err2 := utils.Db.Exec(sqlStr, "admin9", "66666", "11113333@163.com")
	if err2 != nil {
		fmt.Println("插入出现异常", err2)
		return err2
	}
	return nil
}

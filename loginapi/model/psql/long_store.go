package psql

import (
	"fmt"
)

func (s *Store) UserLogin(userName string, passWord string) (bool, error) {
	fmt.Println("sql=>username=", userName, "password=", passWord)
	query := "SELECT COUNT(*) FROM User WHERE username = ? AND password = ?" //表名是区分大小写的

	// 执行查询语句并获取结果
	var count int
	err := s.db.QueryRow(query, userName, passWord).Scan(&count)

	if err != nil {
		return false, err
	}

	// 如果count为1，则意味着找到了匹配的用户
	if count == 1 {
		return true, nil
	}

	return false, nil
	// _, err := s.db.Exec("Select * FROM Login WHERE userName=? and passWord=?", userName, passWord)
	// if err != nil {
	// 	return false, err
	// }
	// return true, nil
}

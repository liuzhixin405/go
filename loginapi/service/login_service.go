package service

import "login/model/psql"

func Login(userName string, password string) bool {
	tx, err := psql.SharedStore().BeginTx()
	if err != nil {
		return false
	}
	defer func() { _ = tx.Rollback() }()

	isLogin := tx.Login(userName, password)
	if isLogin != true {
		return false
	}
	return true
}

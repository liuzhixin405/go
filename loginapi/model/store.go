package model

type Store interface {
	UserLogin(user string, pwd string) (bool, error)
}

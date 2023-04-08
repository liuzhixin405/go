package model

type Store interface {
	Login(user string, pwd string) bool
	CommitTx() error
	Rollback() error
	BeginTx() (Store, error)
}

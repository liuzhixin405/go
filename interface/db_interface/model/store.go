package model

type Store interface {
	Notify(message string)
}

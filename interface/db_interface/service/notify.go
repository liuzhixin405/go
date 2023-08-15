package service

import (
	"testimplement/model/db"
)

func Notify(message string) {
	db.SharedStore().Notify(message)
}

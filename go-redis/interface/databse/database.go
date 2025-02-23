package database

import "go-redis/interface/resp"

type CnmdLine = [][]byte
type Database interface {
	Exec(client resp.Connection, args [][]byte) resp.Reply
	Close()
	AfterClientClise(c resp.Connection)
}

type DataEntity struct {
	Data interface{}
}

package database

import (
	"go-redis/datastruct/dict"
	databaseface "go-redis/interface/databse"
	"go-redis/interface/resp"
	"go-redis/resp/reply"
	"strings"
)

type DB struct {
	index  int
	data   dict.Dict
	addAof func(line CmdLine)
}
type ExecFunc func(db *DB, args [][]byte) resp.Reply
type CmdLine = [][]byte

func MakeDB() *DB {
	return &DB{
		data: dict.MakeSyncDict(),
		addAof: func(line CmdLine) {

		},
	}
}

func (db *DB) Exec(c resp.Connection, cmdLine CmdLine) resp.Reply {
	cmdName := strings.ToLower(string(cmdLine[0]))
	cmd, ok := cmdTable[cmdName]
	if !ok {
		return reply.MakeErrReply("ERR unkonwn command " + cmdName)
	}
	if !validateArity(cmd.arity, cmdLine) {
		return reply.MakeArgNumErrReply(cmdName)
	}
	fun := cmd.exector
	return fun(db, cmdLine[1:])
}

func validateArity(arity int, cmdArgs [][]byte) bool {
	argNum := len(cmdArgs)
	if arity >= 0 {
		return argNum == arity
	}
	return argNum >= -arity
}

func (db *DB) GetEntity(key string) (*databaseface.DataEntity, bool) {
	raw, ok := db.data.Get(key)
	if !ok {
		return nil, false
	}
	entity, _ := raw.(*databaseface.DataEntity)
	return entity, true
}

func (db *DB) PutEntity(key string, entity *databaseface.DataEntity) int {
	return db.data.Put(key, entity)
}
func (db *DB) PutIfExists(key string, entity *databaseface.DataEntity) int {
	return db.data.PutIfExists(key, entity)
}
func (db *DB) PutIfAbsent(key string, entity *databaseface.DataEntity) int {
	return db.data.PutIfAbsent(key, entity)
}

func (db *DB) Remove(key string) int {
	return db.data.Remove(key)
}
func (db *DB) Removes(keys ...string) (deleted int) {
	deleted = 0
	for _, key := range keys {
		_, exists := db.data.Get(key)
		if exists {
			db.Remove(key)
			deleted++
		}
	}
	return deleted
}

func (db *DB) Flush() {
	db.data.Clear()
}

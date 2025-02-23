package database

import (
	databaseface "go-redis/interface/databse"
	"go-redis/interface/resp"
	"go-redis/lib/utils"
	"go-redis/resp/reply"
)

//get
//set
//setnx
//getset
//strlen

func execGet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	entity, exits := db.GetEntity(key)
	if !exits {
		return reply.MakeNullBulkReply()
	}
	bytes := entity.Data.([]byte)
	return reply.MakeBulkReply(bytes)
}

func execSet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := string(args[1])
	entity := &databaseface.DataEntity{
		Data: value,
	}
	db.PutEntity(key, entity)
	db.addAof(utils.ToCmdLine2("set", args...))
	return reply.MakeOkReply()
}

func execSetnx(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := string(args[1])
	entity := &databaseface.DataEntity{
		Data: value,
	}
	result := db.PutIfAbsent(key, entity)
	db.addAof(utils.ToCmdLine2("setnx", args...))
	return reply.MakeIntReply(int64(result))
}

func execGetSet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := string(args[1])
	entity, exits := db.GetEntity(key)

	db.PutEntity(key, &databaseface.DataEntity{Data: value})
	db.addAof(utils.ToCmdLine2("getset", args...))
	if !exits {
		return reply.MakeNullBulkReply()
	}
	return reply.MakeBulkReply(entity.Data.([]byte))
}
func execStrLen(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	entity, exists := db.GetEntity(key)
	if !exists {
		return reply.MakeNullBulkReply()
	}

	bytes := entity.Data.([]byte)
	return reply.MakeIntReply(int64(len(bytes)))
}

func init() {
	RegisterCommand("Get", execGet, 2)
	RegisterCommand("set", execSet, 3)
	RegisterCommand("SetNx", execSetnx, 3) //flushdb a b c
	RegisterCommand("StrLen", execStrLen, 2)
	RegisterCommand("GetSet", execGetSet, 3)
}

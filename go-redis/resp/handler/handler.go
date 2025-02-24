package handler

import (
	"context"
	"go-redis/cluster"
	"go-redis/config"
	"go-redis/database"
	databaseface "go-redis/interface/databse"
	"go-redis/lib/logger"
	"go-redis/lib/sync/atomic"
	"go-redis/resp/connection"
	"go-redis/resp/parser"
	"go-redis/resp/reply"
	"io"
	"net"
	"strings"
	"sync"
)

var (
	unknownErrReplyBytes = []byte("-ERR unknown\r\n")
)

type RespHandler struct {
	activeConn sync.Map
	db         databaseface.Database
	closing    atomic.Boolean
}

func MakeHandler() *RespHandler {
	var db databaseface.Database
	if config.Properties.Self != "" && len(config.Properties.Peers) > 0 {
		db = cluster.MakeClusterDatabase()
	} else {
		db = database.NewStandaloneDatabase()
	}
	//db = database.NewEchoDatabase() //测试
	return &RespHandler{
		db: db,
	}
}
func (r *RespHandler) closeClient(client *connection.Connection) {
	client.Close()
	r.db.AfterClientClise(client)
	r.activeConn.Delete(client)
}
func (r *RespHandler) Handle(ctc context.Context, conn net.Conn) {
	if r.closing.Get() {
		_ = conn.Close()
	}
	client := connection.NewConn(conn)
	r.activeConn.Store(client, struct{}{})
	ch := parser.ParseStream(conn)
	for payload := range ch {
		if payload.Err != nil {
			if payload.Err == io.EOF || payload.Err == io.ErrUnexpectedEOF || strings.Contains(payload.Err.Error(), "use of closed newwork connection") {
				r.closeClient(client)
				logger.Info("connection closed:" + client.RemoteAddr().String())
				return
			}

			errReply := reply.MakeErrReply(payload.Err.Error())
			err := client.Write(errReply.ToBytes())
			if err != nil {
				r.closeClient(client)
				logger.Info("connection closed:" + client.RemoteAddr().String())
				return
			}
			continue
		}
		//exec
		if payload.Data == nil {
			continue
		}
		reply, ok := payload.Data.(*reply.MultiBulkReply)
		if !ok {
			logger.Error("require multi bulk reply")
			continue
		}
		result := r.db.Exec(client, reply.Args)
		if result != nil {
			client.Write(result.ToBytes())
		} else {
			_ = client.Write(unknownErrReplyBytes)
		}
	}
}

func (r *RespHandler) Close() error {
	logger.Info("handler shutting down")
	r.activeConn.Range(func(key interface{}, value interface{}) bool {
		client := key.(*connection.Connection)
		_ = client.Close()
		return true
	})
	r.db.Close()
	return nil
}

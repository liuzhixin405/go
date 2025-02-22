package tcp

import (
	"context"
	"net"
)

type Handler interface {
	Handle(ctc context.Context, conn net.Conn)
	Close() error
}

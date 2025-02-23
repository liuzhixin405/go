package reply

type PongReply struct {
}
type OkReply struct{}

var pongbytes = []byte("+PONG\r\n")
var okBytes = []byte("+OK\r\n")

func (r *OkReply) ToBytes() []byte {
	return okBytes
}
func (r PongReply) ToBytes() []byte {
	return pongbytes
}

func MakePongReply() *PongReply {
	return &PongReply{}
}

var theOkReply = new(OkReply)

func MakeOkReply() *OkReply {
	return theOkReply
}

type NullBulkReply struct{}

var nullBulkbytes = []byte("$-1\r\n")

func (n NullBulkReply) ToBytes() []byte {
	return nullBulkbytes
}

func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

type EmptyMultiBulkReply struct{}

var emptyMultiBulkbytes = []byte("*0\r\n")

func (e EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkbytes
}
func MakeEmptyMultiBulkReply() *EmptyMultiBulkReply {
	return &EmptyMultiBulkReply{}
}

type NoReply struct{}

var noreplybytes = []byte("")

func (e NoReply) ToBytes() []byte {
	return noreplybytes
}
func MakeNoReply() *NoReply {
	return &NoReply{}
}

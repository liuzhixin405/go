package reply

type UnknowErrReply struct {
}

var unknownErrBytes = []byte("-Err unknown\r\n")

func (u UnknowErrReply) Error() string {
	return "Err unknown"
}

func (u UnknowErrReply) ToBytes() []byte {
	return unknownErrBytes
}

type ArgNumErrReply struct {
	Cmd string
}

func (a *ArgNumErrReply) Error() string {
	return "-ERR wrong number of arguments for '" + a.Cmd + "' command"
}

func (a *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of arguments for '" + a.Cmd + "' command\r\n")
}

func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}

type SyntaxErrReply struct{}

var syntaxErrBytes = []byte("-Err syntax error\r\n")
var theSyntaxErrReply = &SyntaxErrReply{}

func MakeSyntaxErrReply() *SyntaxErrReply {
	return theSyntaxErrReply
}
func (s *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}
func (s *SyntaxErrReply) Error() string {
	return "-Err syntax error"
}

type WrongTypeErrReply struct{}

func (w WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key holding the wrong kind of value"
}

func (w WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")

func MakeWrongTypeErrReply() *WrongTypeErrReply {
	return &WrongTypeErrReply{}
}

type ProtocolErrReply struct {
	Msg string
}

func (p *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error: '" + p.Msg + "'\r\n")
}

func (p *ProtocolErrReply) Error() string {
	return "ERR Protocol error: " + p.Msg
}

func MakeProtocolErrReply(msg string) *ProtocolErrReply {
	return &ProtocolErrReply{
		Msg: msg,
	}
}

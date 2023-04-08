package rest

import "github.com/coreos/etcd/error"

type LoginDto struct {
	userName string
	passWord string
}

type messageVo struct {
	Message string `json:"message"`
}

func newMessageVo(err error) *messageVo {
	return &messageVo{
		Message: error.Error(),
	}
}

package rest

// type LoginDto struct {
// 	userName string
// 	passWord string
// }

type loginDto struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type messageVo struct {
	Message string `json:"message"`
}

// func newMessageVo(err error) *messageVo {
// 	return &messageVo{
// 		Message: error.Error(),
// 	}
// }

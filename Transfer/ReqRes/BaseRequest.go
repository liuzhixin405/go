package reqres

type BaseRequest struct {
	CustomerId string
	IP         string
	Token      string
	ClientType ClientType
}

type ClientType int

const (
	Web ClientType = iota
	Android
	Ios
	API
	H5
	Other
)

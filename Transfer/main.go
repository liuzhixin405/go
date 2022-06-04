package main

import (
	"transferasset/rest"
)

func main() {
	rest.StartServer()

	select {}
}

//curl -X POST http://localhost:8080/v1/transfer/getavailablequantity  -d coin=1  -d customerId=test001
//curl -X POST http://localhost:8080/v1/transfer/confirmtransfer
// curl -X POST http://localhost:8080/v1/transfer/transferasset

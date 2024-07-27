package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"samplebtc/tran"
	"time"
)

// 签名
func main() {
	priv, err := tran.GenerateKeyPair()
	if err != nil {
		log.Fatal(err)
	}
	pubKeyBytes := tran.PublicKeyToBytes(&priv.PublicKey)
	tx := tran.Transaction{
		Sender:   pubKeyBytes,
		Receiver: []byte("recevier-address"),
		Amount:   10.0,
		DTime:    time.Now().Unix(),
	}

	signature, err := tran.SignTransaction(&tx, priv)
	if err != nil {
		log.Fatal(err)
	}
	tx.Signature = signature

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("error connecting to server", err)
	}
	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	if err := encoder.Encode(tx); err != nil {
		log.Fatal("Error encoding transaction:", err)
	}
	fmt.Println("Transaction sent")
}

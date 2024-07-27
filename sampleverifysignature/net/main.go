package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"samplebtc/tran"
)

// 验证签名
func handleConnection(conn net.Conn) {
	defer conn.Close()
	decoder := gob.NewDecoder(conn)
	var tx tran.Transaction
	if err := decoder.Decode(&tx); err != nil {
		log.Println("Error decoding transaction:", err)
		return
	}
	fmt.Printf("Received transaction %+v\n", tx)
	senderPubKey := tran.BytesToPublicKey(tx.Sender)
	isValid := tran.VerifySignature(&tx, senderPubKey)
	fmt.Printf("Signature verified: %t\n", isValid)
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Listening on :8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

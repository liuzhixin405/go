package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Transaction struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Amount    int    `json:"amount"`
	Timestamp string `json:"timestamp"`
}

func main() {
	serverURL := "http://localhost:8001"

	// 1. 发送10笔交易
	for i := 1; i <= 10; i++ {
		transaction := Transaction{
			Sender:    fmt.Sprintf("Sender%d", i),
			Receiver:  fmt.Sprintf("Receiver%d", i),
			Amount:    i * 10,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		transactionData, err := json.Marshal(transaction)
		if err != nil {
			log.Fatalf("Failed to marshal transaction %d: %v", i, err)
		}

		resp, err := http.Post(fmt.Sprintf("%s/transaction", serverURL), "application/json", bytes.NewBuffer(transactionData))
		if err != nil {
			log.Fatalf("Failed to send transaction %d: %v", i, err)
		}
		resp.Body.Close()
		fmt.Printf("Transaction %d added, response status: %s\n", i, resp.Status)
	}

	// 2. 启动三个节点并挖矿
	nodePorts := []int{8002, 8003, 8004}

	for _, port := range nodePorts {
		go func(port int) {
			nodeURL := fmt.Sprintf("http://localhost:%d", port)
			resp, err := http.Post(fmt.Sprintf("%s/mine", nodeURL), "application/json", nil)
			if err != nil {
				log.Fatalf("Failed to mine on node %d: %v", port, err)
			}
			defer resp.Body.Close()

			var minedResponse map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&minedResponse); err != nil {
				log.Fatalf("Failed to decode mine response from node %d: %v", port, err)
			}

			fmt.Printf("Node %d mine response: %v\n", port, minedResponse)
		}(port)
	}

	// 防止主协程退出过早
	time.Sleep(10 * time.Second)
}

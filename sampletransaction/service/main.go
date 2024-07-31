package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	difficulty    = 3                // 简化的 PoW 难度
	blockInterval = 30 * time.Second // Interval for creating new blocks
)

var (
	port            int
	blockchain      = []Block{}
	blockchainMu    sync.Mutex
	knownPeers      = []string{}
	knownPeersMu    sync.Mutex
	transactionPool []Transaction //pool
	transactionMu   sync.Mutex
)

type Transaction struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Amount    int    `json:"amount"`
	Timestamp string `json:"timestamp"`
}

type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"` // 使用 Transaction 类型
	Proof        int           `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}

func main() {
	flag.IntVar(&port, "port", 8080, "Port to run the service on")
	flag.Parse()

	http.HandleFunc("/mine", handleMineRequest)
	http.HandleFunc("/block", handleBlockRequest)
	http.HandleFunc("/sync", handleSyncRequest)
	http.HandleFunc("/peers", handlePeersRequest)
	http.HandleFunc("/addpeer", handleAddPeerRequest)
	http.HandleFunc("/transaction", handleTransactionRequest) // 添加处理交易的端点
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting blockchain node on port %d", port)
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func handleTransactionRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var transaction Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		transactionMu.Lock()
		transactionPool = append(transactionPool, transaction)
		transactionMu.Unlock()

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Transaction added: %s", transaction)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func handleMineRequest(w http.ResponseWriter, r *http.Request) {
	blockchainMu.Lock()
	defer blockchainMu.Unlock()

	var lastBlock Block
	if len(blockchain) > 0 {
		lastBlock = blockchain[len(blockchain)-1]
	} else {
		lastBlock = Block{Index: 0,
			PreviousHash: "0",
			Timestamp:    time.Now().String(),
			Transactions: []Transaction{ // 交易池为空时的默认交易
				{Sender: "System", Receiver: "Genesis", Amount: 0, Timestamp: time.Now().String()},
			},
			Proof: 0,
		}
	}

	lastProof := lastBlock.Proof
	proof := proofOfWork(lastProof)
	for !validProof(lastProof, proof) {
		proof++
	}

	timestamp := time.Now().String()

	newBlock := Block{
		Index:        len(blockchain) + 1,
		PreviousHash: calculateHash(lastBlock),
		Timestamp:    timestamp,
		Transactions: extractTransactions(),
		Proof:        proof,
	}
	// 清空交易池
	transactionPool = []Transaction{}
	// Calculate the hash for the new block
	blockHash := calculateHash(newBlock)

	blockchain = append(blockchain, newBlock)

	//fmt.Fprintf(w, "New block mined: %+v\n", newBlock)
	//fmt.Fprintf(w, "Block hash: %s\n", blockHash)
	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(map[string]interface{}{
		"message": "New block mined",
		"block":   newBlock,
		"hash":    blockHash,
	})
	if err != nil {
		http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
		return
	}
	// 在写入 HTTP 响应之前，确保没有其他输出
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("Failed to write response:", err)
	}
}

func handleBlockRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var block Block
		if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		blockchainMu.Lock()
		blockchain = append(blockchain, block)
		blockchainMu.Unlock()

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Block added: %v", block)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func handleSyncRequest(w http.ResponseWriter, r *http.Request) {
	blockchainMu.Lock()
	defer blockchainMu.Unlock()

	response := struct {
		Blocks []Block `json:"blocks"`
	}{
		Blocks: blockchain,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func handlePeersRequest(w http.ResponseWriter, r *http.Request) {
	knownPeersMu.Lock()
	defer knownPeersMu.Unlock()

	response := struct {
		Peers []string `json:"peers"`
	}{
		Peers: knownPeers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func handleAddPeerRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var peer string
		if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		knownPeersMu.Lock()
		knownPeers = append(knownPeers, peer)
		knownPeersMu.Unlock()

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Peer added: %s", peer)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func proofOfWork(lastProof int) int {
	proof := 0
	for !validProof(lastProof, proof) {
		proof++
	}
	return proof
}

func validProof(lastProof, proof int) bool {
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := sha256.Sum256([]byte(guess))
	prefix := make([]byte, difficulty) // 创建难度字节前缀

	// 检查哈希值前缀是否符合难度要求
	return bytes.HasPrefix(guessHash[:], prefix)
}

func calculateHash(block Block) string {
	data := fmt.Sprintf("%d%s%s%d%s", block.Index, block.Timestamp, block.Transactions, block.Proof, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func broadcastNewBlock(newBlock Block) {
	knownPeersMu.Lock()
	defer knownPeersMu.Unlock()

	for _, peer := range knownPeers {
		url := fmt.Sprintf("http://%s/block", peer)
		data, err := json.Marshal(newBlock)
		if err != nil {
			log.Printf("Failed to marshal new block: %v", err)
			continue
		}

		_, err = http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Printf("Failed to broadcast block to %s: %v", peer, err)
		}
	}
}

func extractTransactions() []Transaction {
	var transactions []Transaction
	for _, tx := range transactionPool {
		transactions = append(transactions, tx)
	}
	return transactions
}

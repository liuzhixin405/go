package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func main() {
	// 定义一个区块链切片，用于存储区块链中的所有区块
	var blockchain []Block

	// 设置挖矿的难度，表示哈希值前面需要有多少个零
	difficulty := 4

	// 创建创世区块，并将其添加到区块链中
	genesisBlock := createGenesisBlock()
	blockchain = append(blockchain, genesisBlock)

	// 初始化前一个区块为创世区块
	prevBlock := genesisBlock

	// 挖掘10个新的区块
	for i := 1; i <= 10; i++ {
		// 创建一个新的区块，基于前一个区块和当前的区块编号
		newBlock := createBlock(prevBlock, fmt.Sprintf("Block #%d", i))

		// 挖掘新的区块，找到符合难度要求的哈希值
		minedBlock := mineBlock(newBlock, difficulty)

		// 将挖掘完成的区块添加到区块链中
		blockchain = append(blockchain, minedBlock)

		// 更新前一个区块为刚刚挖掘完成的区块
		prevBlock = minedBlock

		// 输出当前挖掘的区块编号和其哈希值
		fmt.Printf("Block #%d mined with hash: %s,nonce:%d\n", minedBlock.Index, minedBlock.Hash, minedBlock.Nonce)
	}
}

// 计算哈希
func calculateHash(block Block) string {
	recod := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(recod))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 创建区块
func createBlock(oldBlock Block, data string) Block {

	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = ""
	newBlock.Nonce = 0
	return newBlock
}

// 创建创世块
func createGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      calculateHash(Block{}),
		Nonce:     0,
	}
}

// 工作量算法
func mineBlock(block Block, difficulty int) Block {
	startTime := time.Now()
	fmt.Printf("Mineing block #%d , start time, %s\n", block.Index, startTime)
	for {
		block.Hash = calculateHash(block)
		if isValidHashHard(block.Hash, difficulty) {
			break
		}
		block.Nonce++
	}
	endTime := time.Now()
	fmt.Printf("Time taken to mine block: %s\n", endTime.Sub(startTime))
	return block
}

// 难度适中平均10几秒
func isValidHashHard(hash string, difficulty int) bool {
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}
	// 检查前导零
	if hash[:difficulty] != prefix {
		return false
	}
	// 检查特定字符模式，例如：哈希值的第10位到第15位必须是"abcde" 或者10-11 ax  或者第6位是d
	//if hash[9:15] != "abcde" {
	//	return false
	//}
	//if hash[9:12] != "ax" {
	//	return false
	//}
	if hash[5:6] != "d" {
		return false
	}
	// 检查哈希值范围，要求小于一个特定的十六进制值
	//target := "00000fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	target := "00000e0000000000000000000000000000000000000000000000000000000000"
	if hash >= target {
		return false
	}
	return true
}
func isValidHash(hash string, difficulty int) bool {
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}
	return hash[:difficulty] == prefix
}

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	PreHash   []byte
	Data      []byte
}

type BlockChain struct {
	Blocks []*Block
}

func main() {

	blockChain := CreateBlockChain()
	time.Sleep(time.Second)
	blockChain.AddBlock("After genesis,i have  4 BTC")
	time.Sleep(time.Second)
	blockChain.AddBlock("i lost")
	time.Sleep(time.Second)
	blockChain.AddBlock("i have no money")
	time.Sleep(time.Second)

	for _, block := range blockChain.Blocks {
		fmt.Printf("Timestamp:%d\n", block.Timestamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("previous hash : %x\n", block.PreHash)
		fmt.Printf("data:%s\n", block.Data)
	}
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{ToHexInt(b.Timestamp), b.PreHash, b.Data}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]

}

func ToHexInt(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func CreateBlock(preHash, data []byte) *Block {
	block := Block{time.Now().Unix(), []byte{}, preHash, data}
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	genesisWords := "hello ,blockchain!"
	return CreateBlock([]byte{}, []byte(genesisWords))
}

func (bc *BlockChain) AddBlock(data string) {
	newBlock := CreateBlock(bc.Blocks[len(bc.Blocks)-1].Hash, []byte(data))
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateBlockChain() *BlockChain {
	blockChain := BlockChain{}
	blockChain.Blocks = append(blockChain.Blocks, GenesisBlock())
	return &blockChain
}

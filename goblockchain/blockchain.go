package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// 型定義
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// 新規のブロックを作成する関数（nonceとpreviousHashが引数、Blockが返り値）
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// プリント関数の実装（表示形式を整えるため）
func (b *Block) Print() {
	fmt.Printf("timestamp %d\n", b.timestamp)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("previous hash %s\n", b.previousHash)
	fmt.Printf("transactions %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	// jsonの文字列にしてやる
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64
		Nonce        int
		PreviousHash string
		Transactions []string
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

// ブロックを詰め込むストラクト
type Blockchain struct {
	tranzactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

// createBlock関数
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s \n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
	// blockChain := NewBlockchain()
	// fmt.Println(blockChain)
	// blockChain.Print()
	// fmt.Println("--------------")
	// blockChain.CreateBlock(5, "hash 1")
	// blockChain.Print()
	// fmt.Println("--------------")
	// blockChain.CreateBlock(2, "hash 2")
	// blockChain.Print()
	// b := NewBlock(0, "init hash")
	// fmt.Println("--------------")
	// fmt.Println("そのままの場合")
	// fmt.Println(b)
	// fmt.Println("--------------")
	// fmt.Println("--------------")
	// fmt.Println("整えた場合")
	// b.Print()
	// fmt.Println("--------------")
}

// ブロックのハッシュを求める
// sha256を使用

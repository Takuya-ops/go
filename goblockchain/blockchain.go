package main

import (
	"fmt"
	"log"
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

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "init hash")
	fmt.Println("--------------")
	fmt.Println("そのままの場合")
	fmt.Println(b)
	fmt.Println("--------------")
	fmt.Println("--------------")
	fmt.Println("整えた場合")
	b.Print()
	fmt.Println("--------------")

}

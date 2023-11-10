package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 完全な乱数の作成(ファイルが実行した時、毎回違う乱数になるようにする)
	rand.Seed(time.Now().Unix())

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d回目のおみくじ結果: ", i)

		// 乱数の作成
		number := rand.Intn(6)
		switch number {
		case 0:
			fmt.Println("凶です。")
		case 1, 2:
			fmt.Println("吉です。")
		case 3, 4:
			fmt.Println("中吉です。")
		case 5:
			fmt.Println("大吉です。")
		}
	}
}

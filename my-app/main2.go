package main

import (
	"fmt"
	"math/rand"
	"time"
)

func submain() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
func main() {
	// 初期値、範囲、処理
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// 乱数の作成
	for i := 1; i <= 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	// 完全な乱数の作成
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 5; i++ {
		fmt.Println(rand.Intn(10))
	}

	text := "文字"
	// ポインタ型の定義は型に*を付ける
	var name *string = &text
	fmt.Println(&text)
	fmt.Println(name)
	fmt.Println(*name)

}

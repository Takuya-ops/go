// package
package main

// 必要モジュール
import (
	"fmt"
)

// structの定義
type vertex struct {
	x int
	y int
}

// 面積を求める関数
func Area(v vertex) int {
	return v.x * v.y
}

// main関数
func main() {
	v := vertex{3, 5}
	result := Area(v)
	fmt.Println("Area:", result)
}

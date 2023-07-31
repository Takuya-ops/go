package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

// // 値レシーバーを持つメソッド
// func (r Rectangle) Area() int {
// 	return r.width * r.height
// }

// func main() {
// 	rect := Rectangle{width: 10, height: 5}
// 	fmt.Println("Area:", rect.Area()) // Area: 50

// 	// メソッド内で値を変更しても、元のオブジェクトには影響しない
// 	rect.width = 20
// 	fmt.Println("Area:", rect.Area()) // Area: 50
// }

// ポインタレシーバーを持つメソッド
func (r *Rectangle) DoubleSize() {
	r.width *= 2
	r.height *= 2
}

func main() {
	rect := &Rectangle{width: 10, height: 5}
	fmt.Println("Width:", rect.width, "Height:", rect.height) // Width: 10 Height: 5

	// メソッド内でオブジェクトの値を変更する
	rect.DoubleSize()
	fmt.Println("Width:", rect.width, "Height:", rect.height) // Width: 20 Height: 10
}

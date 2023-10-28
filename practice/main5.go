package main

import "fmt"

type vertex struct {
	x int
	y int
}

func square_calc(v vertex) int {
	return v.x * v.y
}

func (v vertex) square_calc() int {
	return v.x * v.y
}

func main() {
	// var v vertex = vertex{5, 8} //長方形
	// v := vertex{5,8}
	// v := vertex{5, 5} //正方形
	// v := vertex{0, 0} //存在しない
	v := vertex{-1, -6} //存在しない
	if (v.x == v.y) && ((v.x > 0) && (v.y > 0)) {
		// fmt.Printf("正方形の面積は、%vです。", square_calc(v))
		// 値レシーバー
		fmt.Printf("正方形の面積は、%vです。", v.square_calc())
	} else if v.x <= 0 && v.y <= 0 {
		fmt.Println("面積は存在しません。")
	} else {
		fmt.Printf("長方形の面積は、%vです。", square_calc(v))
	}
}

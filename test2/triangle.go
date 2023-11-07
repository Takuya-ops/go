package main

import (
	"fmt"
	"math"
)

func calculateDistance(x1, y1, x2, y2 int) float64 {
	// 2点間の距離を計算する関数
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func main() {
	var N int
	fmt.Scan(&N)

	// 点の座標を保存するスライスを作成
	points := make([][2]int, N)

	// 点の座標を入力から読み取る
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		points[i] = [2]int{x, y}
	}

	// 最大の線分の長さを初期化
	maxDistance := 0.0

	// 2点を選んで線分の長さを計算し、最大値を更新
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			distance := calculateDistance(points[i][0], points[i][1], points[j][0], points[j][1])
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	// 結果を出力
	fmt.Printf("%.8f\n", maxDistance)
}

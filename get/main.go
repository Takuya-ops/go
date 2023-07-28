package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// データの構造体 (仮のデータ)
type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// APIエンドポイントのハンドラを登録
	http.HandleFunc("/api/data", getData)

	// サーバーを起動 (ポート: 8080)
	fmt.Println("サーバーを起動します...")
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("サーバーの起動に失敗しました:", err)
		}
	}()
	fmt.Println("サーバーは起動しています。")

	// サーバーが起動している間、メインの処理をブロックする
	select {}
}

// テスト用のデータを返すAPIハンドラ
func getData(w http.ResponseWriter, r *http.Request) {
	// テストデータを作成
	data := Data{
		ID:   1,
		Name: "サンプルデータ",
	}

	// データをJSON形式にエンコードしてレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "データのエンコードに失敗しました", http.StatusInternalServerError)
		return
	}
}

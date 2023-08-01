package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// RequestBody はPOSTリクエストのJSONデータを格納する構造体。
type RequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ResponseBody はレスポンスのJSONデータを格納する構造体。
type ResponseBody struct {
	Message string `json:"message"`
}

func main() {
	// ルーティング設定
	http.HandleFunc("/api", handleAPI)

	// サーバーを起動し、ポート8080でリクエストを待ち受ける
	port := ":8080"
	fmt.Printf("サーバーをポート%sで起動中...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %s", err)
	}
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	// POSTメソッド以外のリクエストを受け付けないようにする
	if r.Method != http.MethodPost {
		http.Error(w, "POSTメソッドのみ受け付けています", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディの読み取り
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}

	// データ検証
	if requestBody.Name == "" || requestBody.Email == "" {
		http.Error(w, "nameとemailの両方のフィールドが必須です", http.StatusBadRequest)
		return
	}

	// 受け取ったデータを使って何らかの処理を行う（ここでは簡単な例としてそのままレスポンスを作成）
	responseMessage := fmt.Sprintf("Hello, %s! Your email is %s.", requestBody.Name, requestBody.Email)

	// レスポンスの作成
	responseBody := ResponseBody{
		Message: responseMessage,
	}

	// レスポンスをJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "レスポンスの作成に失敗しました", http.StatusInternalServerError)
		return
	}
}

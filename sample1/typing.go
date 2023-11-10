package main

import "fmt"

func main() {
	var total int = 0
	calculate(1, "dog", &total)
	calculate(2, "cat", &total)
	calculate(3, "fish", &total)
	fmt.Printf("あなたのスコアは、%d点です。", total)
}

func calculate(number int, word string, score *int) int {
	var input string
	fmt.Printf("[質問%d]次の文字を入力してください：%s\n", number, word)
	// &をつけないと入力されず以降のif文の処理が走ってしまう。
	fmt.Scan(&input)
	if input == word {
		fmt.Println("正解です。")
		*score += 10
	} else {
		fmt.Println("不正解です。")
	}
	return *score
}

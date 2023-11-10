package main

import "fmt"

func main() {
	ask(1, "dragon")
	ask(2, "monkey")
	ask(3, "clokodile")

}

func ask(number int, question string) {
	var input string
	fmt.Printf("[質問%d]次の文字を入力してください：%s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%sと入力されました。", input)
}

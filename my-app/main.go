package main

import "fmt"

const (
	Apple = iota + iota
	Orange
	Banana = iota + 3
)

// func FindUser(name string) (*User, error) {
// 	user := findUserFromGroup("my-group", name)
// 	if user == nil {
// 		return nil, ErrNotFound
// 	}
// 	return nil
// }

func main() {
	fmt.Println(Apple)
	fmt.Println(Orange)
	fmt.Println(Banana)

	// 文字列の書き換え(Goのstringはイミュータブルなためバイト列に変換する必要がある)
	s := "Hello"
	fmt.Println(s)
	b := []byte(s)
	b[0] = 'h'
	s = string(b)
	fmt.Println(s)

	S := "こんにちわ世界"
	fmt.Println(S)
	rs := []rune(S)
	rs[4] = 'は'
	S = string(rs)
	fmt.Println(S)

	// map
	m := map[string]int{
		"john": 21,
		"Bob":  18,
		"Mark": 33,
	}
	fmt.Println(m)
	// mapから要素を削除
	delete(m, "Mark")
	fmt.Println(m)

	// mapのキーと値の列挙
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

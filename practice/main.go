package main

import "fmt"

// 偶数追加の関数
func getEvenOddNumbers(numbers []int) ([]int, []int) {
	var evenNumbers []int
	var oddnumbers []int
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
		if number%2 != 0 {
			oddnumbers = append(oddnumbers, number)
		}
	}
	return evenNumbers, oddnumbers
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers, oddnumbers := getEvenOddNumbers(numbers)
	fmt.Println("元のスライス：", numbers)
	fmt.Println("偶数のみのスライス：", evenNumbers)
	fmt.Println("奇数のみのスライス：", oddnumbers)
}

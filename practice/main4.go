package main

import "fmt"

// 偶数のみのスライスを作成
func getEvenNumbers(numbers []int) []int {
	var evenNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	return evenNumbers
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evennumbers := getEvenNumbers(numbers)
	fmt.Println("偶数のみのスライス:", evennumbers)
}

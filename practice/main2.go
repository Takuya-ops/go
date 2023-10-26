package main

import "fmt"

func getOddNumbers(numbers []int) []int {
	var oddnumbers []int
	for _, number := range numbers {
		if number%2 != 0 {
			oddnumbers = append(oddnumbers, number)
		}
	}
	return oddnumbers
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddnumbers := getOddNumbers(numbers)
	fmt.Println("奇数のスライス：", oddnumbers)
}

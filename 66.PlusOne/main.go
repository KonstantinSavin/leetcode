package main

import "fmt"

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	digits := []int{9, 9, 9, 9}
	fmt.Println(digits)
	fmt.Println(plusOne(digits))
}

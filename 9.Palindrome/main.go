package main

import (
	"fmt"
	"math"
)

// Given an integer x, return true if x is a palindrome, and false otherwise.

func isPalindrome(x int) bool {
	num := x

	if x < 0 {
		return false
	}

	numbers := []int{}
	// fmt.Println(numbers, len(numbers), cap(numbers))
	for x > 0 {
		d := x % 10
		x = x / 10
		// fmt.Println(d, x)
		numbers = append(numbers, d)
	}
	fmt.Println(numbers)
	var revers int
	for i := 0; i < len(numbers); i++ {
		revers += numbers[i] * int(math.Pow(10, float64(len(numbers)-1-i)))
	}
	fmt.Println(revers)

	return num == revers
}

func main() {
	fmt.Println(isPalindrome(1221))
}

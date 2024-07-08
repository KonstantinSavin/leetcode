package main

import "fmt"

func isPerfectSquare(num int) bool {
	right := num
	left := 1
	for left < right {
		mid := (left + right) / 2
		if mid >= num/mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left*left == num
}

func main() {
	fmt.Println(isPerfectSquare(16))
}

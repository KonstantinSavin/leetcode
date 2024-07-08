package main

import "fmt"

func mySqrt(x int) int {
	right := x
	left := 1
	for left < right {
		mid := (left + right) / 2
		if mid >= x/mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println(mySqrt(16))
}

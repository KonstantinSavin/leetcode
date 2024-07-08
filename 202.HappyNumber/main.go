package main

import "fmt"

func isHappy(n int) bool {
	slow, fast := n, squeresSum(n)
	for fast != 1 && slow != fast {
		slow = squeresSum(slow)
		fast = squeresSum(squeresSum(fast))
	}
	return fast == 1
}

func squeresSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		sum += digit * digit
	}
	return sum
}

func main() {
	fmt.Println(isHappy(2))
}

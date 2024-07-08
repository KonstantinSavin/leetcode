package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(3))
}

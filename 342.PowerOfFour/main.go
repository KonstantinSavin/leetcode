package main

import (
	"fmt"
	"math/bits"
)

func isPowerOfFour(n int) bool {
	return n > 0 && (n-1)%3 == 0 && bits.OnesCount(uint(n)) == 1
}

func main() {
	fmt.Println(isPowerOfFour(1), isPowerOfFour(2), isPowerOfFour(15))
}

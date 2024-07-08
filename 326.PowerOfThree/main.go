package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {
	return n > 0 && int(math.Pow(3, 19))%n == 0
}

func main() {
	fmt.Println(isPowerOfThree(81), isPowerOfThree(0), isPowerOfThree(15))
}

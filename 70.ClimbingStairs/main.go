package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	ways := 0
	prev1 := 1
	prev2 := 2

	for i := 2; i < n; i++ {
		ways = prev1 + prev2
		fmt.Println(ways)
		prev1 = prev2
		prev2 = ways
	}

	return ways
}

func main() {
	fmt.Println(climbStairs(2), climbStairs(3), climbStairs(20))
}

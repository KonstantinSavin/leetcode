package main

import "fmt"

func countBits(n int) []int {
	ans := make([]int, n+1, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = i&1 + ans[i/2]
	}
	return ans
}

func main() {
	fmt.Println(countBits(1), countBits(5))
}

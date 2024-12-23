package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(7))
}

func arrangeCoins(n int) int {
	k := 0
	for n >= k+1 {
		k++
		n -= k
	}
	return k
}

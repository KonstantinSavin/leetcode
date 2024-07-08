package main

import "fmt"

func hammingWeight(num uint32) int {
	res := 0

	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}

func main() {
	fmt.Println(hammingWeight(uint32(7)), uint32(7))
}

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(readBinaryWatch(0))
	fmt.Println(readBinaryWatch(1))
	fmt.Println(readBinaryWatch(2))
	fmt.Println(readBinaryWatch(3))
}

func readBinaryWatch(turnedOn int) []string {
	var res []string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if bits.OnesCount(uint(i))+bits.OnesCount(uint(j)) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return res
}

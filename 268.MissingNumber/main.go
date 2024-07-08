package main

import "fmt"

func missingNumber(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		res += i + 1
		res -= nums[i]
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4}
	fmt.Println(missingNumber(nums))
}

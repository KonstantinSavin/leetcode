package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(singleNumber(nums))
}

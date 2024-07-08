package main

import "fmt"

func moveZeroes(nums []int) {
	i := 0
	fmt.Println(nums)
	for j := range nums {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}

func main() {
	// nums := []int{0, 1, 0, 3, 12, 0}
}

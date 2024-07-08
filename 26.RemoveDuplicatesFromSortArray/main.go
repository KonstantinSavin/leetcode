package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	duplicateCount := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			duplicateCount += 1
		} else {
			nums[i-duplicateCount] = nums[i]
		}
	}

	return len(nums) - duplicateCount

}

func main() {
	nums := []int{1, 1, 2, 2, 2}

	fmt.Println(nums)
	fmt.Println(removeDuplicates(nums))
}

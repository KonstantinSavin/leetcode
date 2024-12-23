package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 2, 3, 1, 4, 4, 6}))
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

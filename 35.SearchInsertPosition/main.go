package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println(searchInsert(nums, 3))
}

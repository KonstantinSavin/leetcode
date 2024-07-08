package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	L := 0

	for i := range nums {
		if i-L > k {
			delete(m, nums[L])
			L++
		}
		_, ok := m[nums[i]]
		if ok {
			return true
		}
		m[nums[i]] = i
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}

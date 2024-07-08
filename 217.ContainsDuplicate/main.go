package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if hashmap[nums[i]] == 0 {
			hashmap[nums[i]]++
		} else if hashmap[nums[i]] == 1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 4, 5}
	fmt.Println(containsDuplicate(nums))
}

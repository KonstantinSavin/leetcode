package main

import (
	"fmt"
)

// в тупую
func twoSumTwoLoops(nums []int, target int) []int {
	var answer [2]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				answer[0] = i
				answer[1] = j
			}
		}
	}
	return answer[:]
}

// с мапами
func twoSumHashMap(nums []int, target int) []int {
	var answer [2]int
	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = i
	}
	fmt.Println(hashmap)

	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		j, ok := hashmap[key]
		if ok && i != j {
			answer[0] = i
			answer[1] = j
		}

	}

	return answer[:]
}

func main() {
	nums := []int{3, 5, 2, -4, 8, 11}
	target := 10
	fmt.Println(twoSumTwoLoops(nums, target))
	fmt.Println(twoSumHashMap(nums, target))
}

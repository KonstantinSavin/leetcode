package main

import "fmt"

func twoSum(nums []int, target int) []int {
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
	nums := []int{2, 3, 7, 11, 15}
	fmt.Println(twoSum(nums, 17))
}

package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	maj, cnt := 0, 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			maj = nums[i]
			cnt++
		} else if nums[i] == maj {
			cnt++
		} else {
			cnt--
		}
	}
	return maj
}

func main() {
	nums := []int{1, 1, 1, 2, 3, 3, 3, 3}
	fmt.Println(majorityElement(nums))
}

package main

import "fmt"

func removeElement(nums []int, val int) int {
	i, count := 0, len(nums)
	for i < count {
		if nums[i] == val {
			count--
			nums[i] = nums[count]
			// fmt.Println(i, count)
		} else {
			i++
		}
	}
	return count
}

func main() {
	nums := []int{1, 1, 2, 3, 2}
	fmt.Println(nums)
	fmt.Println(removeElement(nums, 2))

}

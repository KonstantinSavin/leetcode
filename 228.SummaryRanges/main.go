package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	i := 0
	for i < len(nums) {
		start := nums[i]
		for i < len(nums)-1 && nums[i] == nums[i+1]-1 {
			i++
		}
		end := nums[i]
		if start == end {
			res = append(res, fmt.Sprint(start))
		} else {
			res = append(res, fmt.Sprint(start)+"->"+fmt.Sprint(end))
		}
		i++
	}
	return res
}

func main() {
	nums := []int{1, 2, 4, 5, 6, 8}
	fmt.Println(summaryRanges(nums))
}

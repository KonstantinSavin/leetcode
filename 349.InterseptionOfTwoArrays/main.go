package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]bool)
	for _, n := range nums1 {
		m[n] = true
	}
	fmt.Println(m)
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; ok {
			res = append(res, nums2[i])
			delete(m, nums2[i])
		}
	}
	return res
}

func main() {
	num1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 6, 6}
	num2 := []int{6, 2, 9, 10, 6, 2}
	fmt.Println(intersection(num1, num2))
}

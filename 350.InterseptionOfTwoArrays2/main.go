package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]int)
	for k := 0; k < len(nums1); k++ {
		m[nums1[k]]++
	}
	fmt.Println(m)
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; ok {
			res = append(res, nums2[i])
			m[nums2[i]]--
			if m[nums2[i]] == 0 {
				delete(m, nums2[i])
			}
		}
	}
	return res
}

func main() {
	num1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 6, 6}
	num2 := []int{6, 2, 9, 10, 6, 2}
	fmt.Println(intersect(num1, num2))
}

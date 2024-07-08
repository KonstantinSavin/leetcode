package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	index := m + n - 1
	i := m - 1
	j := n - 1

	for ; i >= 0 && j >= 0; index-- {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
	fmt.Println(nums1)

}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	nums3 := []int{0}
	nums4 := []int{1}
	fmt.Println(nums1, len(nums1), nums2, len(nums2))
	fmt.Println(nums3, len(nums3), nums4, len(nums4))
	merge(nums1, len(nums1)-len(nums2), nums2, len(nums2))
	merge(nums3, len(nums3)-len(nums4), nums4, len(nums4))
}

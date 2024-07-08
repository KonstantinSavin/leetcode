package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := (l + r) / 2
	return &TreeNode{nums[m], build(nums, l, m-1), build(nums, m+1, r)}
}

func main() {
	nums := []int{-1, 0, 5, 6, 8, 10}
	fmt.Println(sortedArrayToBST(nums), sortedArrayToBST(nums).Left, sortedArrayToBST(nums).Right)
}

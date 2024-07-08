package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	result = append(result, left...)
	result = append(result, right...)
	result = append(result, root.Val)
	return result
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	output := postorderTraversal(&root)
	fmt.Println(output)
}

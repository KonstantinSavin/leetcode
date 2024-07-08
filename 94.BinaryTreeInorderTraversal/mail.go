package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	result = append(result, left...)
	result = append(result, root.Val)
	result = append(result, right...)
	return result
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	output := inorderTraversal(&root)
	fmt.Println(output)
}

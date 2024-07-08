package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	result := TreeNode{root.Val, root.Right, root.Left}
	result.Left = invertTree(root.Right)
	result.Right = invertTree(root.Left)

	return &result
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println(invertTree(&root), invertTree(root.Left), invertTree(root.Right))
	fmt.Println(&root, root.Left, root.Right)
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Left.Left = &TreeNode{Val: 1}
	root.Right.Right.Left = &TreeNode{Val: 4}
	fmt.Println(minDepth(&root))

	root2 := TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 1}
	root2.Right.Right = &TreeNode{Val: 1}
	root2.Right.Right.Right = &TreeNode{Val: 1}
	fmt.Println(minDepth(&root2))
}

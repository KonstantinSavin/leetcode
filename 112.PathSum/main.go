package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	fmt.Println(root, targetSum)
	return sub(root, targetSum)
}

func sub(root *TreeNode, sum int) bool {
	fmt.Println(root, sum)
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && sum == root.Val {
		return true
	} else {
		return sub(root.Left, sum-root.Val) || sub(root.Right, sum-root.Val)
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

	fmt.Println(hasPathSum(&root, 5))

}

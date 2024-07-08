package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetric2(root, root)
}

func isSymmetric2(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && isSymmetric2(p.Left, q.Right) && isSymmetric2(p.Right, q.Left)
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 4}

	fmt.Println(isSymmetric(&root))
}

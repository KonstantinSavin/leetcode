package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Left.Left = &TreeNode{Val: 4}
	p.Right = &TreeNode{Val: 3}
	p.Right.Left = &TreeNode{Val: 5}
	p.Right.Right = &TreeNode{Val: 6}

	q := TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Left.Left = &TreeNode{Val: 4}
	q.Right = &TreeNode{Val: 3}
	q.Right.Left = &TreeNode{Val: 5}
	q.Right.Right = &TreeNode{Val: 6}

	fmt.Println(isSameTree(&p, &q))
}

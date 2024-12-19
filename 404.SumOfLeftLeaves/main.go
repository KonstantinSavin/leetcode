package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{}
	sumOfLeftLeaves(&root)
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	isLeaf := func(node *TreeNode) bool {
		return node != nil && node.Left == nil && node.Right == nil
	}

	var res int
	if isLeaf(root.Left) {
		res += root.Left.Val
	} else {
		res += sumOfLeftLeaves(root.Left)
	}

	res += sumOfLeftLeaves(root.Right)

	return res
}

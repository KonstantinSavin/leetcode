package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, sb string, res *[]string) {
	if node == nil {
		return
	}

	if sb == "" {
		sb = fmt.Sprintf("%d", node.Val)
	} else {
		sb = fmt.Sprintf("%s->%d", sb, node.Val)
	}

	if node.Left == nil && node.Right == nil {
		*res = append(*res, sb)
		return
	}

	dfs(node.Left, sb, res)
	dfs(node.Right, sb, res)
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	dfs(root, "", &result)
	return result
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println(binaryTreePaths(&root))
}

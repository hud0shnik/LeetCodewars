package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return recr(root, 0)
}

func recr(root *TreeNode, result int) int {
	if root == nil {
		return result
	}

	left := recr(root.Left, result+1)
	right := recr(root.Right, result+1)

	if left > right {
		return left
	}
	return right
}

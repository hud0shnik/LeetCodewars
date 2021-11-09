package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return recr(root.Left, root.Right)
}

func recr(left *TreeNode, right *TreeNode) bool {
	if left == right {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && recr(left.Left, right.Right) && recr(left.Right, right.Left)
}

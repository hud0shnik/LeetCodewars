package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(node *TreeNode, val int) *TreeNode {
	if node == nil || node.Val == val {
		return node
	}
	if val < node.Val {
		return searchBST(node.Left, val)
	}
	if val > node.Val {
		return searchBST(node.Right, val)
	}
	return nil
}

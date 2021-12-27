package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	recr(root, &result)
	return result
}

func recr(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		recr(root.Left, result)
		recr(root.Right, result)
	}
}

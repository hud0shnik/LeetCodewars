package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	recr(root, &result)
	return result
}

func recr(root *TreeNode, result *[]int) {
	if root != nil {
		recr(root.Left, result)
		recr(root.Right, result)
		*result = append(*result, root.Val)
	}
}

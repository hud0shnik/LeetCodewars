package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recr(elemenet *TreeNode) []int {
	if elemenet == nil {
		return nil
	}

	result := []int{elemenet.Val}
	result = append(recr(elemenet.Left), result...)
	result = append(result, recr(elemenet.Right)...)
	return result
}

func inorderTraversal(root *TreeNode) []int {
	return recr(root)
}

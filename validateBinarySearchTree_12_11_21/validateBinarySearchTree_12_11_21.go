package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isOk(node *TreeNode, minimum *int, maximum *int) bool {
	if node == nil {
		return true
	}
	if minimum != nil && node.Val <= *minimum || maximum != nil && node.Val >= *maximum {
		return false
	}
	return isOk(node.Left, minimum, &node.Val) && isOk(node.Right, &node.Val, maximum)
}

func isValidBST(root *TreeNode) bool {
	return isOk(root, nil, nil)
}

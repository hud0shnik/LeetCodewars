package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	recr(root, &result, -1)
	return result
}

func recr(node *TreeNode, m *[][]int, level int) {
	if node == nil {
		return
	}

	currentLevel := level + 1

	for len(*m) < currentLevel+1 {
		*m = append(*m, []int{})
	}

	(*m)[currentLevel] = append((*m)[currentLevel], node.Val)

	recr(node.Left, m, currentLevel)
	recr(node.Right, m, currentLevel)
}

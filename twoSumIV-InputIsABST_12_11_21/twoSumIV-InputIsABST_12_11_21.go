package main

import (
	"sort"
)

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

func twoSum(numbers []int, target int) bool {
	i, j := 0, len(numbers)-1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			break
		}
	}
	return !(i == j)
}

func findTarget(root *TreeNode, k int) bool {
	m := recr(root)
	sort.Ints(m)
	if len(m) == 1 {
		return false
	}
	return twoSum(m, k)
}

/*
Ugly, but working
I used 167's solution and 94's solution
Because it's very funny :^)
*/

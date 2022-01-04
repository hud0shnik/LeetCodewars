package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	result, current := 0, 0
	xc := []int{1, 0, -1, 0}
	yc := []int{0, 1, 0, -1}
	var explore func(int, int)

	explore = func(x int, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] == 0 {
			return
		}

		grid[x][y] = 0
		current++
		for i := 0; i < 4; i++ {
			explore(x+xc[i], y+yc[i])
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				explore(i, j)
			}
			if current > result {
				result = current
			}
			current = 0
		}
	}

	return result
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))

	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0, 0, 0}}))
	fmt.Println(maxAreaOfIsland([][]int{{0, 1, 1, 0, 0, 0, 0, 1}}))
}

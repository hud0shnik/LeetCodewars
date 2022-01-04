package main

import "fmt"

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rottens := [][]int{}
	fresh := 0
	result := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rottens = append(rottens, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	for fresh > 0 && len(rottens) > 0 {

		result++
		rottensL := len(rottens)

		for i := 0; i < rottensL; i++ {

			x := rottens[0][0]
			y := rottens[0][1]
			rottens = rottens[1:]

			if x+1 < len(grid) && grid[x+1][y] == 1 {
				grid[x+1][y] = 2
				fresh--
				rottens = append(rottens, []int{x + 1, y})
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				grid[x][y-1] = 2
				fresh--
				rottens = append(rottens, []int{x, y - 1})
			}
			if x-1 >= 0 && grid[x-1][y] == 1 {
				grid[x-1][y] = 2
				fresh--
				rottens = append(rottens, []int{x - 1, y})
			}
			if y+1 < len(grid[0]) && grid[x][y+1] == 1 {
				grid[x][y+1] = 2
				fresh--
				rottens = append(rottens, []int{x, y + 1})
			}

		}
	}

	if fresh > 0 {
		return -1
	}
	return result

}

func main() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 1},
		{0, 1, 2}}))
	/*fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))*/
}

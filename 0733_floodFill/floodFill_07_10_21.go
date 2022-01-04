package main

import "fmt"

func colorize(image [][]int, x, y, oldColor, newColor int) {
	if x < 0 || y < 0 || x >= len(image) || y >= len(image[0]) || image[x][y] != oldColor {
		return
	}
	image[x][y] = newColor
	colorize(image, x-1, y, oldColor, newColor)
	colorize(image, x+1, y, oldColor, newColor)
	colorize(image, x, y-1, oldColor, newColor)
	colorize(image, x, y+1, oldColor, newColor)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		colorize(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func main() {
	fmt.Println(floodFill([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1}}, 1, 1, 2))
	fmt.Println(floodFill([][]int{
		{0, 0, 0},
		{0, 0, 0}}, 0, 0, 2))
}

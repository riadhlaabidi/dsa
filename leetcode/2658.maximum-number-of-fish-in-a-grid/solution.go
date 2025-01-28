package main

import "fmt"

var directions = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func isValidWater(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] != 0
}

func dfs(grid [][]int, visited [][]bool, r, c int) int {
	visited[r][c] = true
	ans := grid[r][c]
	for _, dir := range directions {
		nextR := r + dir[0]
		nextC := c + dir[1]
		if isValidWater(grid, nextR, nextC) && !visited[nextR][nextC] {
			ans += dfs(grid, visited, nextR, nextC)
		}
	}
	return ans
}

func findMaxFish(grid [][]int) int {
	ans := 0

	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[0]))
	}

	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] != 0 && !visited[r][c] {
				maxFish := dfs(grid, visited, r, c)
				if maxFish > ans {
					ans = maxFish
				}
			}
		}
	}

	return ans
}

func main() {
	grid := [][]int{
		{0, 2, 1, 0},
		{4, 0, 0, 3},
		{1, 0, 0, 4},
		{0, 3, 2, 0},
	}

	expected := 7
	actual := findMaxFish(grid)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d\n", expected, actual)
		return
	}

	fmt.Printf("Correct %d\n", actual)
}

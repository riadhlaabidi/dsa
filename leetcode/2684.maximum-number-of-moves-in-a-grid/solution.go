package main

import "fmt"

func maxMoves(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([][]int, rows)

	for i := range dp {
		dp[i] = []int{1, 0}
	}

	moves := 0

	for j := 1; j < cols; j++ {
		for i := range rows {
			if i-1 >= 0 && grid[i][j] > grid[i-1][j-1] && dp[i-1][0] != 0 {
				dp[i][1] = max(dp[i][1], dp[i-1][0]+1)
			}

			if grid[i][j] > grid[i][j-1] && dp[i][0] != 0 {
				dp[i][1] = max(dp[i][1], dp[i][0]+1)
			}

			if i+1 < rows && grid[i][j] > grid[i+1][j-1] && dp[i+1][0] != 0 {
				dp[i][1] = max(dp[i][1], dp[i+1][0]+1)
			}

			moves = max(moves, dp[i][1]-1)
		}

		for k := range rows {
			dp[k][0] = dp[k][1]
			dp[k][1] = 0
		}
	}

	return moves
}

func main() {
	grid := [][]int{
		{2, 4, 3, 5},
		{5, 4, 9, 3},
		{3, 4, 2, 11},
		{10, 9, 13, 15},
	}

	expected := 3
	actual := maxMoves(grid)

	if actual != expected {
		fmt.Printf("Expected %d, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %d\n", actual)

	grid = [][]int{
		{3, 2, 4},
		{2, 1, 9},
		{1, 1, 7},
	}

	expected = 0
	actual = maxMoves(grid)

	if actual != expected {
		fmt.Printf("Expected %d, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %d", actual)
}

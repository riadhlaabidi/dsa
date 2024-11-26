package main

import "fmt"

func findChampion(grid [][]int) int {
	n := len(grid)
	for i := 0; i < n; i++ {
		j := 0
		for ; j < n; j++ {
			if i != j && grid[i][j] == 0 {
				break
			}
		}

		if j == n {
			return i
		}
	}

	return -1
}

func main() {
	grid := [][]int{
		{0, 0, 1},
		{1, 0, 1},
		{0, 0, 0},
	}

	expected := 1
	actual := findChampion(grid)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d\n", expected, actual)
		return
	}

	fmt.Printf("Correct %d\n", actual)
}

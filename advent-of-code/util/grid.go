package util

type Cell struct {
	Row int
	Col int
}

func NewCell(row, col int) Cell {
	return Cell{row, col}
}

var Directions = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func NewGrid[T any](m, n int) [][]T {
	grid := make([][]T, m)
	for i := range grid {
		grid[i] = make([]T, n)
	}
	return grid
}

func IsInBounds(m, n, r, c int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

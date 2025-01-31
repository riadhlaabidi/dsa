package main

import "fmt"

type Cell struct {
	r, c int
}

var directions = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func largestIsland(grid [][]int) int {
	islands := make(map[int]int)
	islandId := 2
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 1 {
				islands[islandId] = area(grid, r, c, islandId)
				islandId++
			}
		}
	}

	if len(islands) == 0 {
		return 1
	}

	if len(islands) == 1 {
		islandId--
		area := islands[islandId]
		if area == len(grid)*len(grid[0]) {
			return area
		}
		return area + 1
	}

	largest := 0

	for r := range grid {
		for c := range grid[0] {
			connecting := make(map[int]bool)
			if grid[r][c] == 0 {
				for _, dir := range directions {
					nextR := r + dir[0]
					nextC := c + dir[1]
					if isInBounds(grid, nextR, nextC) {
						connecting[grid[nextR][nextC]] = true
					}
				}
				if len(connecting) >= 2 {
					area := 1
					for id := range connecting {
						area += islands[id]
					}
					if area > largest {
						largest = area
					}
				}
			}
		}
	}

	return largest
}

func area(grid [][]int, r, c, islandId int) int {
	area := 1
	q := make([]Cell, 0)
	q = append(q, Cell{r, c})
	grid[r][c] = islandId

	for len(q) > 0 {
		cell := q[0]
		q = q[1:]

		for _, dir := range directions {
			nextR := cell.r + dir[0]
			nextC := cell.c + dir[1]

			if isInBounds(grid, nextR, nextC) && grid[nextR][nextC] == 1 {
				area++
				grid[nextR][nextC] = islandId
				q = append(q, Cell{nextR, nextC})
			}
		}
	}

	return area
}

func isInBounds(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func main() {
	grid := [][]int{{0, 0}, {0, 0}}
	expected := 1
	actual := largestIsland(grid)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d\n", expected, actual)
		return
	}

	fmt.Printf("Correct %d\n", actual)
}

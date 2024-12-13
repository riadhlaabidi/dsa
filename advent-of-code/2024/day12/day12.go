package main

import (
	"aoc/util"
	"fmt"
)

type Zone struct {
	perimeter int
	area      int
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func isInBounds(m, n, r, c int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func getZone(grid [][]byte, visited [][]bool, i, j int, z byte, perimeter *int, area *int) {
	if visited[i][j] {
		return
	}

	visited[i][j] = true
	*area++

	for _, dir := range dirs {
		nextR := i + dir[0]
		nextC := j + dir[1]

		if isInBounds(len(grid), len(grid[0]), nextR, nextC) && grid[nextR][nextC] == z {
			getZone(grid, visited, nextR, nextC, z, perimeter, area)
		} else {
			*perimeter++
		}
	}
}

func solvePartI(lines [][]byte) {
	ans := 0
	m := len(lines)
	n := len(lines[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i, line := range lines {
		for j, ch := range line {
			if !visited[i][j] {
				perimeter := 0
				area := 0
				getZone(lines, visited, i, j, ch, &perimeter, &area)
				ans += area * perimeter
			}
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func getZone2(grid [][]byte, visited [][]bool, i, j int, z byte, perimeter *int, area *int) {
	if visited[i][j] {
		return
	}

	visited[i][j] = true
	*area++

	for _, dir := range dirs {
		nextR := i + dir[0]
		nextC := j + dir[1]

		if isInBounds(len(grid), len(grid[0]), nextR, nextC) && grid[nextR][nextC] == z {
			getZone2(grid, visited, nextR, nextC, z, perimeter, area)
		} else {
			*perimeter++
		}
	}
}

func solvePartII(lines [][]byte) {
	ans := 0
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day12.test")
	inputLines := util.ReadInput("day12.input")
	fmt.Printf("---- Day 12 Test ----\n")
	solvePartI(testLines)
	// solvePartII(testLines)
	fmt.Printf("---- Day 12 Input ----\n")
	solvePartI(inputLines)
	// solvePartII(inputLines)
}

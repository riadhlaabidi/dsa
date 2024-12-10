package main

import (
	"aoc/util"
	"fmt"
)

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Cell struct {
	r int
	c int
}

func isInBounds(m, n, r, c int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func getScore(grid [][]byte, visited map[Cell]bool, m, n, i, j int) int {
	if visited[Cell{i, j}] {
		return 0
	}

	visited[Cell{i, j}] = true

	if grid[i][j] == '9' {
		return 1
	}

	score := 0

	for _, dir := range dirs {
		nextR := i + dir[0]
		nextC := j + dir[1]
		if isInBounds(m, n, nextR, nextC) && grid[nextR][nextC] == grid[i][j]+1 {
			score += getScore(grid, visited, m, n, nextR, nextC)
		}
	}

	return score
}

func getRating(grid [][]byte, m, n, i, j int) int {
	if grid[i][j] == '9' {
		return 1
	}

	rating := 0

	for _, dir := range dirs {
		nextR := i + dir[0]
		nextC := j + dir[1]
		if isInBounds(m, n, nextR, nextC) && grid[nextR][nextC] == grid[i][j]+1 {
			rating += getRating(grid, m, n, nextR, nextC)
		}
	}

	return rating
}

func solvePartI(lines [][]byte) {
	m := len(lines)
	n := len(lines[0])
	ans := 0

	for i, line := range lines {
		for j, ch := range line {
			if ch == '0' {
				visited := make(map[Cell]bool)
				ans += getScore(lines, visited, m, n, i, j)
			}
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	m := len(lines)
	n := len(lines[0])
	ans := 0

	for i, line := range lines {
		for j, ch := range line {
			if ch == '0' {
				ans += getRating(lines, m, n, i, j)
			}
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day10.test")
	inputLines := util.ReadInput("day10.input")
	fmt.Printf("---- Day 10 Test ----\n")
	solvePartI(testLines)  /* 36 */
	solvePartII(testLines) /* 81 */
	fmt.Printf("---- Day 10 Input ----\n")
	solvePartI(inputLines)  /* 674 */
	solvePartII(inputLines) /* 1372 */
}

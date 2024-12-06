package main

import (
	"aoc/util"
	"fmt"
)

var dirs = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func getInitialPosition(lines [][]byte) (int, int) {
	for i, line := range lines {
		for j, ch := range line {
			if ch == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isOutOfBounds(m, n, r, c int) bool {
	return r <= 0 || r >= m-1 || c <= 0 || c >= n-1
}

func countDistinctPositions(grid [][]byte, gr, gc int) int {
	m, n := len(grid), len(grid[0])
	positions := 0
	direction := 0
	ir := gr // initial guard row pos
	ic := gc // initial guard col pos

	for !isOutOfBounds(m, n, gr, gc) {
		dx := dirs[direction][0]
		dy := dirs[direction][1]

		if grid[gr+dx][gc+dy] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		gr += dx
		gc += dy

		if grid[gr][gc] != 'X' {
			grid[gr][gc] = 'X'
			positions++
		}
	}

	// restore initial guard position
	grid[ir][ic] = '^'

	return positions
}

func isInCycle(grid [][]byte, gr, gc int) bool {
	m, n := len(grid), len(grid[0])

	dejaVu := make([][]int, m)

	for i := 0; i < m; i++ {
		dejaVu[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dejaVu[i][j] = -1
		}
	}

	direction := 0

	for !isOutOfBounds(m, n, gr, gc) {
		dx := dirs[direction][0]
		dy := dirs[direction][1]

		if grid[gr+dx][gc+dy] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		gr += dx
		gc += dy

		if dejaVu[gr][gc] == direction {
			return true
		}

		dejaVu[gr][gc] = direction
	}

	return false
}

func solvePartI(lines [][]byte) {
	gr, gc := getInitialPosition(lines)
	if gr < 0 || gc < 0 {
		panic("Guard not found!")
	}

	ans := countDistinctPositions(lines, gr, gc)

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0

	gr, gc := getInitialPosition(lines)
	if gr < 0 || gc < 0 {
		panic("Guard not found!")
	}

	// Only test cells from the guard positions already
	// drawn in part 1 as X's
	for i, line := range lines {
		for j, ch := range line {
			if ch != 'X' || (i == gr && j == gc) {
				continue
			}

			lines[i][j] = '#'
			if isInCycle(lines, gr, gc) {
				ans++
			}
			lines[i][j] = '.'
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day6.test")
	inputLines := util.ReadInput("day6.input")
	fmt.Printf("---- Day 6 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 6 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

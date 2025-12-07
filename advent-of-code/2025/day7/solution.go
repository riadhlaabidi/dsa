package main

import (
	"aoc/util"
	"fmt"
)

func solvePartI(lines [][]byte) {
	ans := 0
	m := len(lines[0])
	beams := make([]bool, m)
	beams[m/2] = true // start
	for _, line := range lines[1:] {
		for j, char := range line {
			if char == '^' {
				if beams[j] {
					beams[j] = false
					beams[j-1] = true
					beams[j+1] = true
					ans++
				}
			}
		}
	}
	fmt.Printf("Part 1: %v\n", ans)
}

func countTimelines(grid [][]byte, beamRow int, beamCol int, memo map[[2]int]int) int {
	if beamRow >= len(grid) {
		return 0
	}
	if beamRow == len(grid)-1 {
		return 1
	}
	key := [2]int{beamRow, beamCol}
	if value, exists := memo[key]; exists {
		return value
	}
	if grid[beamRow][beamCol] == '.' {
		memo[key] = countTimelines(grid, beamRow+1, beamCol, memo)
	} else {
		left := countTimelines(grid, beamRow+1, beamCol-1, memo)
		right := countTimelines(grid, beamRow+1, beamCol+1, memo)
		memo[key] = left + right
	}
	return memo[key]
}

func solvePartII(lines [][]byte) {
	m := len(lines[0])
	memo := make(map[[2]int]int)
	ans := countTimelines(lines, 1, m/2, memo)
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 7 Test ----\n")
	solvePartI(testLines)  // 21
	solvePartII(testLines) // 40
	fmt.Printf("---- Day 7 Input ----\n")
	solvePartI(inputLines)  // 1711
	solvePartII(inputLines) // 36706966158365
}

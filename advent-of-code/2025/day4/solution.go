package main

import (
	"aoc/util"
	"fmt"
)

func solvePartI(lines [][]byte) {
	m := len(lines)
	n := len(lines[0])
	ans := 0
	for row, line := range lines {
		for col, char := range line {
			if char == '.' {
				continue
			}
			count := 0
			for _, dir := range util.EightAdjacentDirections {
				newRow := row + dir[0]
				newCol := col + dir[1]

				if util.IsInBounds(m, n, newRow, newCol) && lines[newRow][newCol] == '@' {
					count++
				}
			}
			if count < 4 {
				ans++
			}
		}
	}
	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	m := len(lines)
	n := len(lines[0])
	ans := 0
	for {
		removed := 0
		for row, line := range lines {
			for col, char := range line {
				if char == '.' {
					continue
				}
				count := 0
				for _, dir := range util.EightAdjacentDirections {
					newRow := row + dir[0]
					newCol := col + dir[1]

					if util.IsInBounds(m, n, newRow, newCol) && lines[newRow][newCol] == '@' {
						count++
					}
				}
				if count < 4 {
					removed++
					lines[row][col] = '.'
				}
			}
		}
		if removed == 0 {
			break
		}
		ans += removed
	}
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 4 Test ----\n")
	solvePartI(testLines)  //13
	solvePartII(testLines) // 43
	fmt.Printf("---- Day 4 Input ----\n")
	solvePartI(inputLines)  // 1569
	solvePartII(inputLines) // 9280
}

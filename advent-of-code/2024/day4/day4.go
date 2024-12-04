package main

import (
	"aoc/util"
	"fmt"
)

func isSafe(m, n, r, c int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func solvePartI(lines [][]byte) {
	dirs := [][]int{ /* all dierctions icluding diagonals */
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	ans := 0
	key := "MAS"

	for i, line := range lines {
		for j := range line {
			if lines[i][j] != 'X' {
				continue
			}

			for _, dir := range dirs {
				ind := 0
				newR := i + dir[0]
				newC := j + dir[1]

				for isSafe(len(lines), len(line), newR, newC) {
					if lines[newR][newC] != key[ind] {
						break
					}

					ind++
					newR += dir[0]
					newC += dir[1]

					if ind == len(key) {
						ans++
						break
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] != 'A' {
				continue
			}

			chars := string([]byte{
				lines[i-1][j-1], // top left
				lines[i-1][j+1], // top right
				lines[i+1][j-1], // bottom left
				lines[i+1][j+1], // bottom right
			})

			if chars == "MSMS" || chars == "SMSM" || chars == "SSMM" || chars == "MMSS" {
				ans++
			}
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day4.test")
	inputLines := util.ReadInput("day4.input")
	fmt.Printf("---- Day 4 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 4 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

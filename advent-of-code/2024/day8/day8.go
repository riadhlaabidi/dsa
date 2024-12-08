package main

import (
	"aoc/util"
	"fmt"
)

type Node struct {
	row, col int
}

func isInBounds(m, n, r, c int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func solvePartI(lines [][]byte) {
	antennas := make(map[byte][]Node)
	antinodes := make(map[Node]bool)

	m := len(lines)
	n := len(lines[0])

	for i, line := range lines {
		for j, ch := range line {
			if ch == '.' {
				continue
			}

			if locations, exists := antennas[ch]; exists {
				for _, pairAntenna := range locations {
					dx := i - pairAntenna.row
					dy := j - pairAntenna.col

					antinode1 := Node{pairAntenna.row - dx, pairAntenna.col - dy}
					antinode2 := Node{i + dx, j + dy}

					if isInBounds(m, n, antinode1.row, antinode1.col) {
						antinodes[antinode1] = true
					}
					if isInBounds(m, n, antinode2.row, antinode2.col) {
						antinodes[antinode2] = true
					}
				}
			}

			antennas[ch] = append(antennas[ch], Node{i, j})
		}
	}

	fmt.Printf("Part 1: %v\n", len(antinodes))
}

func solvePartII(lines [][]byte) {
	antennas := make(map[byte][]Node)
	antinodes := make(map[Node]bool)

	m := len(lines)
	n := len(lines[0])

	for i, line := range lines {
		for j, ch := range line {
			if ch == '.' {
				continue
			}

			if locations, exists := antennas[ch]; exists {
				for _, pairAntenna := range locations {
					dr := i - pairAntenna.row
					dc := j - pairAntenna.col

					r, c := pairAntenna.row, pairAntenna.col
					for isInBounds(m, n, r, c) {
						antinodes[Node{r, c}] = true
						r -= dr
						c -= dc
					}

					r, c = i, j
					for isInBounds(m, n, r, c) {
						antinodes[Node{r, c}] = true
						r += dr
						c += dc
					}
				}
			}

			antennas[ch] = append(antennas[ch], Node{i, j})
		}
	}

	fmt.Printf("Part 2: %v\n", len(antinodes))
}

func main() {
	testLines := util.ReadInput("day8.test")
	inputLines := util.ReadInput("day8.input")
	fmt.Printf("---- Day 8 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 8 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

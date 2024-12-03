package main

import (
	"aoc/util"
	"fmt"
)

func solvePartI(lines [][]byte) {
	floor := 0
	for _, line := range lines {
		for _, c := range line {
			if c == '(' {
				floor++
			} else if c == ')' {
				floor--
			}
		}
	}

	fmt.Printf("Part 1: %d\n", floor)
}

func solvePartII(lines [][]byte) {
	floor := 0
	pos := -1
	for _, line := range lines {
		for i, c := range line[:len(line)-1] {
			if c == '(' {
				floor++
			} else {
				floor--
			}

			if floor < 0 {
				pos = i + 1
				break
			}
		}
	}

	fmt.Printf("Part 2: %d\n", pos)
}

func main() {
	inputLines := util.ReadInput("day1.input")
	fmt.Printf("---- Day 1 ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

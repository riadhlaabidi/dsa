package main

import (
	"aoc/util"
	"fmt"
)

func solvePartI(lines [][]byte) {
	ans := 0
	maxPosition := 100
	position := 50

	for _, line := range lines {
		direction := 1

		if line[0] == 'L' {
			direction = -1
		}

		steps := util.GetInt(string(line[1:]))
		newPosition := position + direction*steps
		position = ((newPosition % maxPosition) + maxPosition) % maxPosition

		if position == 0 {
			ans++
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0
	maxPosition := 100
	position := 50

	for _, line := range lines {
		direction := 1
		stepsTillZero := (maxPosition - position) % maxPosition

		if line[0] == 'L' {
			direction = -1
			stepsTillZero = position
		}

		steps := util.GetInt(string(line[1:]))

		if steps > stepsTillZero {
			if position != 0 {
				ans++
			}
			ans += (steps - stepsTillZero - 1) / maxPosition
		}

		newPosition := position + direction*steps
		position = ((newPosition % maxPosition) + maxPosition) % maxPosition

		if position == 0 {
			ans++
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 1 Test ----\n")
	solvePartI(testLines)  // 3
	solvePartII(testLines) // 6
	fmt.Printf("---- Day 1 Input ----\n")
	solvePartI(inputLines)  // 995
	solvePartII(inputLines) // 5847
}

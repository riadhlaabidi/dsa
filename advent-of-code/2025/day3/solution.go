package main

import (
	"aoc/util"
	"fmt"
)

func solvePartI(lines [][]byte) {
	ans := 0

	for _, line := range lines {
		bestRight := 0
		maxJoltage := 0

		for i := len(line) - 1; i >= 0; i-- {
			digit := util.GetInt(string(line[i]))
			if bestRight != 0 {
				maxJoltage = max(maxJoltage, digit*10+bestRight)
			}
			bestRight = max(bestRight, digit)
		}

		ans += maxJoltage
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := int64(0)

	for _, line := range lines {
		n := len(line)
		maxJoltage := make([]byte, 12)
		idx := 0
		start := 0
		numBatteries := 12

		for numBatteries > 0 {
			end := n - numBatteries + 1
			bestDigit := byte('0')
			bestDigitPos := start

			for i := start; i < end; i++ {
				currentDigit := line[i]
				if currentDigit > bestDigit {
					bestDigit = currentDigit
					bestDigitPos = i
					if currentDigit == '9' {
						break
					}
				}
			}

			maxJoltage[idx] = bestDigit
			idx++
			start = bestDigitPos + 1
			numBatteries--
		}

		// fmt.Printf("line %s adding joltage %s\n", line, string(maxJoltage))
		ans += util.GetInt64(string(maxJoltage))
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 3 Test ----\n")
	solvePartI(testLines)  // 357
	solvePartII(testLines) // 3121910778619
	fmt.Printf("---- Day 3 Input ----\n")
	solvePartI(inputLines)  // 17359
	solvePartII(inputLines) // 172787336861064
}

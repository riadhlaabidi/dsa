package main

import (
	"aoc/util"
	"fmt"
)

func isFilled(line []byte) bool {
	for _, ch := range line {
		if ch != '#' {
			return false
		}
	}
	return true
}

func solvePartI(lines [][]byte) {
	ans := 0
	keys := make([][5]int, 0)
	locks := make([][5]int, 0)

	pins := [5]int{-1, -1, -1, -1, -1}
	for i, line := range lines {
		if len(line) == 0 {
			if isFilled(lines[i-1]) {
				keys = append(keys, pins)
			} else {
				locks = append(locks, pins)
			}
			pins = [5]int{-1, -1, -1, -1, -1}
			continue
		}

		for j, ch := range line {
			if ch == '#' {
				pins[j]++
			}
		}
	}

	if isFilled(lines[len(lines)-1]) {
		keys = append(keys, pins)
	} else {
		locks = append(locks, pins)
	}

	for _, lock := range locks {
		for _, key := range keys {
			fit := true
			for i, col := range key {
				if col+lock[i] > 5 {
					fit = false
					break
				}
			}
			if fit {
				ans++
			}
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day25.test")
	inputLines := util.ReadInput("day25.input")
	fmt.Printf("---- Day 25 Test ----\n")
	solvePartI(testLines)
	fmt.Printf("---- Day 25 Input ----\n")
	solvePartI(inputLines)
}

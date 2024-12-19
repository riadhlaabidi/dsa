package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func countWays(patterns []string, design string, cache map[string]int) int {
	if _, exists := cache[design]; !exists {
		if len(design) == 0 {
			return 1
		}

		ways := 0
		for _, pattern := range patterns {
			if strings.HasPrefix(design, pattern) {
				ways += countWays(patterns, design[len(pattern):], cache)
			}
		}
		cache[design] = ways
	}

	return cache[design]
}

func solvePartI(lines [][]byte) {
	ans := 0
	patterns := strings.Split(string(lines[0]), ", ")
	cache := make(map[string]int)

	for _, design := range lines[2:] {
		if countWays(patterns, string(design), cache) > 0 {
			ans++
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0
	patterns := strings.Split(string(lines[0]), ", ")
	cache := make(map[string]int)

	for _, design := range lines[2:] {
		ans += countWays(patterns, string(design), cache)
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day19.test")
	inputLines := util.ReadInput("day19.input")
	fmt.Printf("---- Day 19 Test ----\n")
	solvePartI(testLines)  /* 6 */
	solvePartII(testLines) /* 16 */
	fmt.Printf("---- Day 19 Input ----\n")
	solvePartI(inputLines)  /* 371 */
	solvePartII(inputLines) /* 650354687260341 */
}

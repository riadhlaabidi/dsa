package main

import (
	"aoc/util"
	"fmt"
)

type Sequence [4]int

func generate(secret int) int {
	secret = ((secret * 64) ^ secret) % 16777216
	secret = ((secret / 32) ^ secret) % 16777216
	secret = ((secret * 2048) ^ secret) % 16777216
	return secret
}

func solvePartI(lines [][]byte) {
	ans := 0

	for _, line := range lines {
		secret := util.GetInt(string(line))
		for range 2000 {
			secret = generate(secret)
		}
		ans += secret
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0
	sequences := make(map[Sequence]int)

	for _, line := range lines {
		secret := util.GetInt(string(line))
		changes := make([]int, 0)
		occured := make(map[Sequence]bool)
		previousPrice := secret % 10
		for range 2000 {
			secret = generate(secret)
			currentPrice := secret % 10
			changes = append(changes, currentPrice-previousPrice)
			previousPrice = currentPrice
			if len(changes) == 4 {
				sequence := Sequence(changes)
				if !occured[sequence] {
					occured[sequence] = true
					price := sequences[sequence] + currentPrice
					sequences[sequence] = price
					if price > ans {
						ans = price
					}
				}
				changes = changes[1:]
			}
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines1 := util.ReadInput("day22-1.test")
	testLines2 := util.ReadInput("day22-2.test")
	inputLines := util.ReadInput("day22.input")
	fmt.Printf("---- Day 22 Test ----\n")
	solvePartI(testLines1)  /* 37327623 */
	solvePartII(testLines2) /* 23 */
	fmt.Printf("---- Day 22 Input ----\n")
	solvePartI(inputLines)  /* 20215960478 */
	solvePartII(inputLines) /* 2221 */
}

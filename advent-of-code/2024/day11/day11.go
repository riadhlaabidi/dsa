package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

var cache = make(map[[2]int]int)

func hasEvenDigits(n int) bool {
	x := len(fmt.Sprintf("%d", n))
	return x&1 == 0
}

func split(number int) (int, int) {
	num := fmt.Sprintf("%d", number)
	n := len(num) >> 1
	x := util.GetInt(num[:n])
	y := util.GetInt(num[n:])
	return x, y
}

func getStones(x, n int, cache map[[2]int]int) int {
	if n == 0 {
		return 1
	}
	if _, exists := cache[[2]int{x, n}]; !exists {
		res := 0
		if x == 0 {
			res = getStones(1, n-1, cache)
		} else if hasEvenDigits(x) {
			a, b := split(x)
			res += getStones(a, n-1, cache)
			res += getStones(b, n-1, cache)
		} else {
			res = getStones(x*2024, n-1, cache)
		}
		cache[[2]int{x, n}] = res
	}

	return cache[[2]int{x, n}]
}

func solvePartI(lines [][]byte) {
	line := lines[0]
	stones := strings.Split(string(line), " ")
	ans := 0

	for _, stone := range stones {
		ans += getStones(util.GetInt(stone), 25, cache)
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	line := lines[0]
	stones := strings.Split(string(line), " ")

	ans := 0

	for _, stone := range stones {
		ans += getStones(util.GetInt(stone), 75, cache)
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day11.test")
	inputLines := util.ReadInput("day11.input")
	fmt.Printf("---- Day 11 Test ----\n")
	solvePartI(testLines)  /* 55312 */
	solvePartII(testLines) /* 65601038650482 */
	fmt.Printf("---- Day 11 Input ----\n")
	solvePartI(inputLines)  /* 203953 */
	solvePartII(inputLines) /* 242090118578155 */
}

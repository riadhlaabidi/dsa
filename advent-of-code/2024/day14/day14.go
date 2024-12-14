package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func solvePartI(lines [][]byte) {
	// test
	// m := 7
	// n := 11

	m := 103
	n := 101
	seconds := 100
	middlej := n / 2
	middlei := m / 2
	quadrants := make([]int, 4)
	ans := 1

	for _, line := range lines {
		pv := strings.Split(string(line), " ")
		p := strings.Split(strings.Split(pv[0], "=")[1], ",")
		v := strings.Split(strings.Split(pv[1], "=")[1], ",")
		pj := util.GetInt(p[0])
		pi := util.GetInt(p[1])
		vj := util.GetInt(v[0])
		vi := util.GetInt(v[1])

		finali := (pi + vi*seconds) % m
		finalj := (pj + vj*seconds) % n

		if finali < 0 {
			finali = m + finali
		}

		if finalj < 0 {
			finalj = n + finalj
		}

		if finali == middlei || finalj == middlej {
			continue
		}

		i, j := 0, 0

		if finali > middlei {
			i = 1
		}

		if finalj > middlej {
			j = 1
		}

		quadrants[2*i+j]++
	}

	for _, q := range quadrants {
		ans *= q
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day14.test")
	inputLines := util.ReadInput("day14.input")
	fmt.Printf("---- Day 14 Test ----\n")
	solvePartI(testLines) /* 12 */
	// solvePartII(testLines)
	fmt.Printf("---- Day 14 Input ----\n")
	solvePartI(inputLines) /* 236628054 */
	// solvePartII(inputLines)
}

package main

import (
	"aoc/util"
	"fmt"
)

func move(currR, currC int, direction byte) (int, int) {
	switch direction {
	case '^':
		currR--
	case '>':
		currC++
	case 'v':
		currR++
	case '<':
		currC--
	}

	return currR, currC
}

func solvePartI(lines [][]byte) {
	visited := make(map[[2]int]bool)
	r := 0
	c := 0

	visited[[2]int{r, c}] = true

	for _, ch := range lines[0] {
		r, c = move(r, c, ch)
		visited[[2]int{r, c}] = true
	}

	fmt.Printf("Part 1: %v\n", len(visited))
}

func solvePartII(lines [][]byte) {
	line := lines[0]
	visited := make(map[[2]int]bool)

	santaR := 0
	santaC := 0

	roboR := 0
	roboC := 0

	visited[[2]int{santaR, santaC}] = true

	for i := 0; i < len(line)-1; i += 2 {
		santaR, santaC = move(santaR, santaC, line[i])
		roboR, roboC = move(roboR, roboC, line[i+1])
		visited[[2]int{santaR, santaC}] = true
		visited[[2]int{roboR, roboC}] = true
	}

	fmt.Printf("Part 2: %v\n", len(visited))
}

func main() {
	testLines := util.ReadInput("day3.test")
	inputLines := util.ReadInput("day3.input")
	fmt.Printf("---- Day 3 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 3 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

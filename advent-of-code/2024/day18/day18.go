package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Cell struct {
	r int
	c int
}

type Pair struct {
	cell  Cell
	steps int
}

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func getBytesMap(lines [][]byte) map[Cell]bool {
	bytes := make(map[Cell]bool)

	for _, line := range lines {
		rc := strings.Split(string(line), ",")
		c := util.GetInt(rc[0])
		r := util.GetInt(rc[1])
		bytes[Cell{r, c}] = true
	}

	return bytes
}

func isInBounds(n, r, c int) bool {
	return r >= 0 && r < n && c >= 0 && c < n
}

func walk(bytes map[Cell]bool, n int) int {
	visited := make(map[Cell]bool)
	q := make([]Pair, 0)
	q = append(q, Pair{Cell{0, 0}, 0})

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		if visited[c.cell] {
			continue
		}

		visited[c.cell] = true

		for _, dir := range directions {
			nextR := c.cell.r + dir[0]
			nextC := c.cell.c + dir[1]
			nextCell := Cell{nextR, nextC}
			_, corrupted := bytes[nextCell]
			_, seen := visited[nextCell]
			if isInBounds(n, nextR, nextC) && !seen && !corrupted {
				if nextR == n-1 && nextC == n-1 {
					return c.steps + 1
				}
				q = append(q, Pair{nextCell, c.steps + 1})
			}
		}
	}

	return 0
}

func getFirstBlockingByte(lines [][]byte, n int) string {
	left := 0
	right := len(lines)

	for left < right {
		middle := (left + right) / 2
		bytes := getBytesMap(lines[:middle])
		steps := walk(bytes, n)
		if steps == 0 {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return string(lines[right-1])
}

func solvePartI(lines [][]byte) {
	if len(lines) == 25 {
		// test lines
		bytes := getBytesMap(lines[:12])
		steps := walk(bytes, 7)
		fmt.Printf("Part 1: %v\n", steps)

	} else {
		// input lines
		bytes := getBytesMap(lines[:1024])
		steps := walk(bytes, 71)
		fmt.Printf("Part 1: %v\n", steps)
	}
}

func solvePartII(lines [][]byte) {
	l := len(lines)
	if l == 25 {
		// test lines
		ans := getFirstBlockingByte(lines, 7)
		fmt.Printf("Part 2: %v\n", ans)
	} else {
		// input lines
		ans := getFirstBlockingByte(lines, 71)
		fmt.Printf("Part 2: %v\n", ans)
	}
}

func main() {
	testLines := util.ReadInput("day18.test")
	inputLines := util.ReadInput("day18.input")
	fmt.Printf("---- Day 18 Test ----\n")
	solvePartI(testLines)  /* 22 */
	solvePartII(testLines) /* 6,1 */
	fmt.Printf("---- Day 18 Input ----\n")
	solvePartI(inputLines)  /* 296 */
	solvePartII(inputLines) /* 28,44 */
}

package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func solvePartI(lines [][]byte) {
	n := len(lines)
	redTiles := make([]util.Cell, n)
	for i, line := range lines {
		coord := strings.Split(string(line), ",")
		row := util.GetInt(coord[0])
		col := util.GetInt(coord[1])
		redTiles[i] = util.NewCell(row, col)
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			dh := redTiles[i].Row - redTiles[j].Row
			if dh < 0 {
				dh = -dh
			}
			dw := redTiles[i].Col - redTiles[j].Col
			if dw < 0 {
				dw = -dw
			}
			area := (dh + 1) * (dw + 1)
			ans = max(ans, area)
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

//func solvePartII(lines [][]byte) {
//	ans := 0
//
//	for i, line := range lines {
//	}
//
//	fmt.Printf("Part 2: %v\n", ans)
//}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 9 Test ----\n")
	solvePartI(testLines) // 50
	//solvePartII(testLines)
	fmt.Printf("---- Day 9 Input ----\n")
	solvePartI(inputLines) // 4759930955
	//solvePartII(inputLines)
}

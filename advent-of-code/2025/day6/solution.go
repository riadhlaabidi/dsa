package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func solvePartI(lines [][]byte) {
	ans := int64(0)
	n := len(lines)
	operators := strings.Fields(string(lines[n-1]))
	operatorsLen := len(operators)
	results := make([]int64, operatorsLen)
	for i := range results {
		if operators[i] == "*" {
			results[i] = 1
		}
	}
	for _, line := range lines[:n-1] {
		operands := strings.Fields(string(line))
		for j, operandStr := range operands {
			operand := util.GetInt64(operandStr)
			switch operators[j] {
			case "*":
				results[j] *= operand
			case "+":
				results[j] += operand
			}
		}
	}
	for _, result := range results {
		ans += result
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
	fmt.Printf("---- Day 6 Test ----\n")
	solvePartI(testLines) // 4277556
	// solvePartII(testLines)
	fmt.Printf("---- Day 6 Input ----\n")
	solvePartI(inputLines) // 6299564383938
	//solvePartII(inputLines)
}

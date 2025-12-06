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
	results := make([]int64, len(operators))

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

func solvePartII(lines [][]byte) {
	ans := int64(0)
	n := len(lines)
	operatorsLine := lines[n-1]
	result := int64(0)
	var operator byte

	for col, c := range operatorsLine {
		next := col + 1
		if next < len(operatorsLine) && operatorsLine[next] != ' ' {
			ans += result
			result = 0
			continue
		}
		if c != ' ' {
			if c == '*' {
				result = 1
			}
			operator = c
		}

		num := int64(0)
		for row := range n - 1 {
			digit := lines[row][col]
			if digit != ' ' {
				num = num*10 + int64(digit-'0')
			}
		}

		switch operator {
		case '*':
			result *= num
		case '+':
			result += num
		}
	}
	ans += result
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 6 Test ----\n")
	solvePartI(testLines)  // 4277556
	solvePartII(testLines) // 3263827
	fmt.Printf("---- Day 6 Input ----\n")
	solvePartI(inputLines)  // 6299564383938
	solvePartII(inputLines) // 11950004808442
}

package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func getEquation(line []byte) (int64, []int64) {
	l := strings.Split(string(line), ": ")
	testValue := util.GetInt64(l[0])
	strs := strings.Split(l[1], " ")
	operands := make([]int64, len(strs))

	for i, str := range strs {
		operands[i] = util.GetInt64(str)
	}

	return testValue, operands
}

func concat(left, right int64) int64 {
	r := right

	// pad left operand with zeros at the end
	for r != 0 {
		r /= 10
		left = left * 10
	}

	return left + right
}

func couldBeTrue(operands []int64, i int, test, rollingRes int64, checkConcat bool) bool {
	if rollingRes > test {
		return false
	}

	if i == len(operands) {
		if rollingRes == test {
			return true
		}
		return false
	}

	return (checkConcat &&
		couldBeTrue(operands, i+1, test, concat(rollingRes, operands[i]), checkConcat)) ||
		couldBeTrue(operands, i+1, test, rollingRes+operands[i], checkConcat) ||
		couldBeTrue(operands, i+1, test, rollingRes*operands[i], checkConcat)
}

func solvePartI(lines [][]byte) {
	ans := int64(0)

	for _, line := range lines {
		testValue, operands := getEquation(line)

		if couldBeTrue(operands, 1, testValue, operands[0], false) {
			ans += testValue
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := int64(0)

	for _, line := range lines {
		testValue, operands := getEquation(line)

		if couldBeTrue(operands, 1, testValue, operands[0], true) {
			ans += testValue
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day7.test")
	inputLines := util.ReadInput("day7.input")
	fmt.Printf("---- Day 7 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 7 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

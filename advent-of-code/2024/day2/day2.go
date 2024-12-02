package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getInt(str string) int {
	n, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse int from \"%s\"", str))
	}

	return int(n)
}

func isSafe(nums []int) bool {
	increasing := (nums[1] - nums[0]) > 0

	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]

		if increasing && (diff < 1 || diff > 3) {
			return false
		}

		if !increasing && (diff > -1 || diff < -3) {
			return false
		}
	}

	return true
}

func solvePartI(lines [][]byte) {
	safeCount := 0

	for _, line := range lines {
		strNums := strings.Split(string(line), " ")

		nums := []int{}
		for _, str := range strNums {
			nums = append(nums, getInt(str))
		}

		if isSafe(nums) {
			safeCount++
		}
	}

	fmt.Printf("Part 1: %v\n", safeCount)
}

func solvePartII(lines [][]byte) {
	safeCount := 0

	for _, line := range lines {
		strNums := strings.Split(string(line), " ")

		nums := []int{}
		for _, str := range strNums {
			nums = append(nums, getInt(str))
		}

		if isSafe(nums) {
			safeCount++
			continue
		}

		for idx := 0; idx < len(nums); idx++ {
			newNums := slices.Concat(nums[:idx], nums[idx+1:])
			if isSafe(newNums) {
				safeCount++
				break
			}
		}
	}

	fmt.Printf("Part 2: %v\n", safeCount)
}

func main() {
	testLines := util.ReadInput("day2.test")
	inputLines := util.ReadInput("day2.input")
	fmt.Printf("---- Day 1 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 1 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func solvePartI(lines [][]byte) {
	n := len(lines)
	left := make([]int, n)
	right := make([]int, n)

	for i, line := range lines {
		nums := strings.Split(string(line), "   ")
		left[i] = util.GetInt(nums[0])
		right[i] = util.GetInt(nums[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	ans := 0
	for i := 0; i < n; i++ {
		distance := left[i] - right[i]

		if distance < 0 {
			distance *= -1
		}

		ans += distance
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	left := make([]int, 0)
	right := make(map[int]int)

	for _, line := range lines {
		nums := strings.Split(string(line), "   ")

		leftNum := util.GetInt(nums[0])
		left = append(left, leftNum)

		rightNum := util.GetInt(nums[1])
		right[rightNum] += 1
	}

	ans := 0
	for _, num := range left {
		ans += num * right[num]
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day1.test")
	inputLines := util.ReadInput("day1.input")
	fmt.Printf("---- Day 1 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 1 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

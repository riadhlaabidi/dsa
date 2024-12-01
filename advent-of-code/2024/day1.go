package main

import (
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

func solvePartI(lines [][]byte) {
	n := len(lines)
	left := make([]int, n)
	right := make([]int, n)

	for i, line := range lines {
		nums := strings.Split(string(line), "   ")
		left[i] = getInt(nums[0])
		right[i] = getInt(nums[1])
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

		leftNum := getInt(nums[0])
		left = append(left, leftNum)

		rightNum := getInt(nums[1])
		right[rightNum] += 1
	}

	ans := 0
	for _, num := range left {
		ans += num * right[num]
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := ReadInput("input/day1.test")
	inputLines := ReadInput("input/day1.input")
	fmt.Printf("---- Day 1 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 1 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

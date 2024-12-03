package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func solvePartI(lines [][]byte) {
	ans := 0
	for _, line := range lines {
		dims := strings.Split(string(line), "x")
		length := util.GetInt(dims[0])
		width := util.GetInt(dims[1])
		height := util.GetInt(dims[2])

		slack := length * width * height / max(length, width, height)

		ans += (2*length*width + 2*width*height + 2*height*length + slack)
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0
	for _, line := range lines {
		dims := strings.Split(string(line), "x")
		length := util.GetInt(dims[0])
		width := util.GetInt(dims[1])
		height := util.GetInt(dims[2])
		ribbon := 2*(length+width+height) - 2*max(length, width, height)
		bow := length * width * height
		ans += ribbon + bow
	}

	fmt.Printf("Part 2: %d\n", ans)
}

func main() {
	inputLines := util.ReadInput("day2.input")
	fmt.Printf("---- Day 2 ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

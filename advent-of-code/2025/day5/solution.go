package main

import (
	"aoc/util"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type idRange struct {
	start int64
	end   int64
}

func solvePartI(lines [][]byte) {
	ans := 0
	ranges := make([]idRange, 0)
	i := 0

	for _, line := range lines {
		i++
		if len(line) == 0 {
			break
		}
		strs := strings.Split(string(line), "-")
		start := util.GetInt64(strs[0])
		end := util.GetInt64(strs[1])
		ranges = append(ranges, idRange{start, end})
	}

	slices.SortFunc(ranges, func(a idRange, b idRange) int {
		return cmp.Compare(a.start, b.start)
	})

	merged := []idRange{}
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1].end+1 < r.start {
			merged = append(merged, r)
		} else {
			if r.end > merged[len(merged)-1].end {
				merged[len(merged)-1].end = r.end
			}
		}
	}

	for _, numStr := range lines[i:] {
		left := 0
		right := len(merged)
		num := util.GetInt64(string(numStr))

		for left < right {
			middle := (left + right) / 2
			if merged[middle].end < num {
				left = middle + 1
			} else if merged[middle].start > num {
				right = middle
			} else {
				ans++
				break
			}
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := int64(0)
	ranges := make([]idRange, 0)
	i := 0

	for _, line := range lines {
		i++
		if len(line) == 0 {
			break
		}
		strs := strings.Split(string(line), "-")
		start := util.GetInt64(strs[0])
		end := util.GetInt64(strs[1])
		ranges = append(ranges, idRange{start, end})
	}

	slices.SortFunc(ranges, func(a idRange, b idRange) int {
		return cmp.Compare(a.start, b.start)
	})

	merged := []idRange{}
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1].end+1 < r.start {
			merged = append(merged, r)
		} else {
			if r.end > merged[len(merged)-1].end {
				merged[len(merged)-1].end = r.end
			}
		}
	}

	for _, r := range merged {
		ans += (r.end - r.start) + 1
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 5 Test ----\n")
	solvePartI(testLines)  // 3
	solvePartII(testLines) // 14
	fmt.Printf("---- Day 5 Input ----\n")
	solvePartI(inputLines)  // 848
	solvePartII(inputLines) // 334714395325710
}

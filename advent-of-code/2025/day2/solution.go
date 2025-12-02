package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

func solvePartI(lines [][]byte) {
	ans := int64(0)
	ranges := strings.Split(string(lines[0]), ",")

	for _, r := range ranges {
		rangeStr := strings.Split(r, "-")
		startStr := rangeStr[0]
		endStr := rangeStr[1]
		n := len(startStr)
		m := len(endStr)
		start := util.GetInt64(startStr)
		end := util.GetInt64(endStr)

		// If both lengths are odd, no business here
		if n&1 != 0 && m&1 != 0 {
			continue
		}

		halfStr := startStr[:n/2]

		if n&1 != 0 {
			halfStr = fmt.Sprintf("1%s", strings.Repeat("0", n/2))
		}

		half := util.GetInt64(halfStr)

		for {
			invalidIdStr := fmt.Sprintf("%s%s", halfStr, halfStr)
			invalidId := util.GetInt64(invalidIdStr)

			if invalidId > end {
				break
			}

			if invalidId >= start {
				ans += invalidId
			}

			half++
			halfStr = strconv.FormatInt(half, 10)
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := int64(0)
	ranges := strings.Split(string(lines[0]), ",")

	for _, r := range ranges {
		rangeStr := strings.Split(r, "-")
		start := util.GetInt64(rangeStr[0])
		end := util.GetInt64(rangeStr[1])

		for i := start; i <= end; i++ {
			str := strconv.FormatInt(i, 10)
			n := len(str)
			for j := 1; j <= n/2; j++ {
				if n%j == 0 {
					invalid := true
					for k := j; k < n; k += j {
						if str[:j] != str[k:k+j] {
							invalid = false
							break
						}
					}
					if invalid {
						ans += i
						break
					}
				}
			}
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 2 Test ----\n")
	solvePartI(testLines)  // 1227775554
	solvePartII(testLines) // 4174379265
	fmt.Printf("---- Day 2 Input ----\n")
	solvePartI(inputLines)  // 28844599675
	solvePartII(inputLines) // 48778605167
}

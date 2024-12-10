package main

import (
	"aoc/util"
	"fmt"
)

type File struct {
	blocks int
	id     int
}

func solvePartI(lines [][]byte) {
	n := len(lines[0])
	line := make([]int, n)

	for i := range line {
		line[i] = int(lines[0][i] - '0')
	}

	left := 0
	right := n - 1
	pos := 0
	ans := 0

	for left < right {
		for range line[left] {
			fileId := left >> 1
			if (left & 1) == 1 {
				if line[right] == 0 {
					if left >= right-2 {
						break
					}
					right -= 2
				}
				fileId = right >> 1
				line[right] -= 1
			}
			ans += (fileId * pos)
			pos++
		}
		left++
	}

	// count remaining blocks from right pointer if they exist.
	// In my case, test data needs this when it breaks from
	// the previous loop, but actual input doesn't
	for range line[right] {
		ans += ((right >> 1) * pos)
		pos++
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	n := len(lines[0])
	line := make([]int, n)

	for i := range line {
		line[i] = int(lines[0][i] - '0')
	}

	right := n - 1
	left := 1
	moved := make(map[int][]File)

	for left < right {
		for i := left; i < right; i += 2 {
			if line[i] >= line[right] {
				moved[i] = append(moved[i], File{line[right], right >> 1})
				line[i] -= line[right]

				// add removed blocks amount to the previous adjacent
				// free space (right-1) to keep their freed space
				line[right-1] += line[right]

				// free blocks
				line[right] = 0

				// move left if there is no more free spaces
				if i == left && line[i] == 0 {
					left += 2
				}
				break
			}
		}
		right -= 2
	}

	pos := 0
	ans := 0

	for i := range line {
		if i&1 == 1 { /* free space (which may be filled) */
			if files, exists := moved[i]; exists {
				for _, file := range files {
					for range file.blocks {
						ans += file.id * pos
						pos++
					}
				}
			}
			// In case there is free space added after moving a file
			// example: 0099.111777(2)44 => 00992111[777(.)]44, now 7's has one
			// free space
			pos += line[i]
		} else { /* filled space */
			for range line[i] {
				ans += ((i >> 1) * pos)
				pos++
			}
		}

	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day9.test")
	inputLines := util.ReadInput("day9.input")
	fmt.Printf("---- Day 9 Test ----\n")
	solvePartI(testLines)  /* 1928 */
	solvePartII(testLines) /* 2858 */
	fmt.Printf("---- Day 9 Input ----\n")
	solvePartI(inputLines)  /* 6367087064415 */
	solvePartII(inputLines) /* 6390781891880 */
}

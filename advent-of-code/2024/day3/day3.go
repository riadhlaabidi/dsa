package main

import (
	"aoc/util"
	"fmt"
)

func solvePartI(lines [][]byte) {
	ans := 0
	mul := "mul("

	for _, line := range lines {
		j := 0
		for i := 0; i < len(line); i++ {
			if line[i] != mul[j] {
				j = 0
				continue
			}
			j++

			if j == len(mul) {
				i++
				lOperandIdx := i
				rOperandIdx := 0
				x := 0
				y := 0
				for i < len(line) {
					ch := line[i]
					if ch == ',' && i > lOperandIdx {
						x = util.GetInt(string(line[lOperandIdx:i]))
						rOperandIdx = i + 1
					} else if ch == ')' {
						if rOperandIdx > lOperandIdx+1 {
							y = util.GetInt(string(line[rOperandIdx:i]))
							ans += x * y
						}
						break
					} else if ch < '0' || ch > '9' {
						break
					}
					i++
				}
				j = 0
			}
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func updateAdd(line []byte, idx *int, add *bool) {
	i := *idx
	k := 0
	do := "do()"
	dont := "don't()"
	for ; i < len(line); i++ {
		if k == len(dont) {
			*add = false
			break
		}
		if line[i] != dont[k] {
			if k < len(do) && line[i] != do[k] {
				break
			}
			if k == len(do) {
				*add = true
				break
			}
		}
		k++
	}
}

func solvePartII(lines [][]byte) {
	ans := 0
	mul := "mul("
	add := true

	for _, line := range lines {
		j := 0
		for i := 0; i < len(line); i++ {
			updateAdd(line, &i, &add)

			if !add {
				continue
			}

			if line[i] != mul[j] {
				j = 0
				continue
			}
			j++

			if j == len(mul) {
				i++
				lOperandIdx := i
				rOperandIdx := 0
				x := 0
				y := 0
				for i < len(line) {
					ch := line[i]
					if ch == ',' && i > lOperandIdx {
						x = util.GetInt(string(line[lOperandIdx:i]))
						rOperandIdx = i + 1
					} else if ch == ')' {
						if rOperandIdx > lOperandIdx+1 {
							y = util.GetInt(string(line[rOperandIdx:i]))
							ans += x * y
						}
						break
					} else if ch < '0' || ch > '9' {
						break
					}
					i++
				}
				j = 0
			}
		}
	}
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day3.test")
	inputLines := util.ReadInput("day3.input")
	fmt.Printf("---- Day 3 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 3 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

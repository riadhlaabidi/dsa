package main

import (
	"bufio"
	"fmt"
	"os"
)

func solvePartI() {
	file, err := os.Open("input/day1.input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	floor := 0
	for _, c := range line {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}

	fmt.Printf("Part 1: %d\n", floor)
}

func solvePartII() {
	file, err := os.Open("input/day1.input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	floor := 0
	pos := -1
	for i, c := range line[:len(line)-1] {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			pos = i + 1
			break
		}
	}

	fmt.Printf("Part 2: %d\n", pos)
}

func main() {
	fmt.Printf("---- Day 1 ----\n")
	solvePartI()
	solvePartII()
}

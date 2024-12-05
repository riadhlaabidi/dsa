package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func parseRules(lines [][]byte) (map[int]map[int]bool, int) {
	i := 0
	rules := make(map[int]map[int]bool)

	// parse rules
	for ; len(lines[i]) > 0; i++ {
		pages := strings.Split(string(lines[i]), "|")
		left := util.GetInt(pages[0])
		right := util.GetInt(pages[1])

		if _, exists := rules[left]; !exists {
			rules[left] = make(map[int]bool)
		}

		rules[left][right] = true
	}

	return rules, i
}

func isValidUpdate(rules map[int]map[int]bool, update []string) bool {
	for i := 0; i < len(update)-1; i++ {
		left := util.GetInt(update[i])
		right := util.GetInt(update[i+1])

		if _, exists := rules[left]; !exists {
			return false
		}

		if _, exists := rules[left][right]; !exists {
			return false
		}
	}

	return true
}

func fixUpdate(rules map[int]map[int]bool, update []string) []int {
	indegrees := make(map[int]int)

	for _, page := range update {
		p := util.GetInt(page)
		indegrees[p] = 0
	}

	for _, pageNumber := range update {
		p := util.GetInt(pageNumber)
		if pagesAfterP, exists := rules[p]; exists {
			for page := range pagesAfterP {
				if _, exists := indegrees[page]; exists {
					indegrees[page]++
				}
			}
		}
	}

	q := make([]int, 0)
	for page, indegree := range indegrees {
		if indegree == 0 {
			q = append(q, page)
			delete(indegrees, page)
		}
	}

	newUpdate := make([]int, 0)

	for len(q) != 0 {
		page := q[0]
		q = q[1:]
		newUpdate = append(newUpdate, page)

		for pageAfter := range rules[page] {
			if _, exists := indegrees[pageAfter]; exists {
				indegrees[pageAfter]--
				if indegrees[pageAfter] == 0 {
					q = append(q, pageAfter)
					delete(indegrees, page)
				}
			}
		}
	}

	return newUpdate
}

func solvePartI(lines [][]byte) {
	ans := 0

	rules, i := parseRules(lines)

	for _, line := range lines[i+1:] {
		update := strings.Split(string(line), ",")
		n := len(update)

		if isValidUpdate(rules, update) {
			ans += util.GetInt(update[n/2])
		}
	}

	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	ans := 0

	rules, i := parseRules(lines)

	for _, line := range lines[i+1:] {
		update := strings.Split(string(line), ",")
		n := len(update)

		if !isValidUpdate(rules, update) {
			fixedUpdate := fixUpdate(rules, update)
			ans += fixedUpdate[n/2]
		}
	}

	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day5.test")
	inputLines := util.ReadInput("day5.input")
	fmt.Printf("---- Day 5 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 5 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

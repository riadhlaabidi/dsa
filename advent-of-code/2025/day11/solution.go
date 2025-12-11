package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type path struct {
	start string
	fft   bool
	dac   bool
}

func dfs(edges map[string][]string, memo map[string]int, node string) int {
	numPaths, exists := memo[node]
	if exists {
		return numPaths
	}

	if node == "out" {
		memo[node] = 1
		return 1
	}

	for _, dest := range edges[node] {
		numPaths += dfs(edges, memo, dest)
	}
	memo[node] = numPaths
	return numPaths
}

func dfs2(edges map[string][]string, memo map[path]int, node string, fft bool, dac bool) int {
	key := path{node, fft, dac}
	numPaths, exists := memo[key]
	if exists {
		return numPaths
	}

	if node == "out" {
		if fft && dac {
			numPaths = 1
		}
		memo[key] = numPaths
		return numPaths
	}

	for _, dest := range edges[node] {
		nfft := fft || (dest == "fft")
		ndac := dac || (dest == "dac")
		numPaths += dfs2(edges, memo, dest, nfft, ndac)
	}
	memo[key] = numPaths
	return numPaths
}

func solvePartI(lines [][]byte) {
	edges := make(map[string][]string)
	for _, line := range lines {
		words := strings.Fields(string(line))
		src := words[0][:len(words[0])-1]
		edges[src] = words[1:]
	}
	memo := make(map[string]int)
	ans := dfs(edges, memo, "you")
	fmt.Printf("Part 1: %v\n", ans)
}

func solvePartII(lines [][]byte) {
	edges := make(map[string][]string)
	for _, line := range lines {
		words := strings.Fields(string(line))
		src := words[0][:len(words[0])-1]
		edges[src] = words[1:]
	}
	memo := make(map[path]int)
	ans := dfs2(edges, memo, "svr", false, false)
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("test_input")
	testLines2 := util.ReadInput("test_input2")
	inputLines := util.ReadInput("input")
	fmt.Printf("---- Day 11 Test ----\n")
	solvePartI(testLines)   // 5
	solvePartII(testLines2) // 2
	fmt.Printf("---- Day 11 Input ----\n")
	solvePartI(inputLines)  // 788
	solvePartII(inputLines) // 316291887968000
}

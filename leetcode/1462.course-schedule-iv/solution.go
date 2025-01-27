package main

import (
	"fmt"
	"slices"
)

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adjList := make(map[int][]int)
	for _, pre := range prerequisites {
		if _, exists := adjList[pre[0]]; !exists {
			adjList[pre[0]] = make([]int, 0)
		}
		adjList[pre[0]] = append(adjList[pre[0]], pre[1])
	}

	isReachable := make([][]bool, numCourses)
	for i := range numCourses {
		isReachable[i] = make([]bool, numCourses)
	}

	bfs(numCourses, adjList, isReachable)

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = isReachable[q[0]][q[1]]
	}

	return ans
}

func bfs(numCourses int, adjList map[int][]int, isReachable [][]bool) {
	for i := range numCourses {
		q := make([]int, 0)
		q = append(q, i)

		for len(q) > 0 {
			node := q[0]
			q = q[1:]

			if adj, exists := adjList[node]; exists {
				for _, n := range adj {
					if !isReachable[i][n] {
						isReachable[i][n] = true
						q = append(q, n)
					}
				}
			}
		}
	}
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	queries := [][]int{{0, 1}, {1, 0}}
	expected := []bool{false, true}
	actual := checkIfPrerequisite(numCourses, prerequisites, queries)

	if !slices.Equal(actual, expected) {
		fmt.Printf("Expected %v, but got %v\n", expected, actual)
		return
	}

	fmt.Printf("Correct %v\n", actual)
}

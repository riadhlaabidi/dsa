package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := make([][]int, numCourses)
	indegrees := make([]int, numCourses)

	for _, e := range prerequisites {
		nodes[e[1]] = append(nodes[e[1]], e[0])
		indegrees[e[0]]++
	}

	q := make([]int, 0)
	for node, ind := range indegrees {
		if ind == 0 {
			q = append(q, node)
		}
	}

	i := 0
	for len(q) != 0 {
		popped := q[0]
		q = q[1:]
		i++

		for _, node := range nodes[popped] {
			indegrees[node]--

			if indegrees[node] == 0 {
				q = append(q, node)
			}
		}
	}

	return i == numCourses
}

func main() {
	numCourses := 5
	prerequisites := [][]int{
		{1, 4},
		{2, 4},
		{3, 1},
		{3, 2},
	}

	expected := true
	actual := canFinish(numCourses, prerequisites)

	if actual != expected {
		fmt.Printf("Expected %t, but got %t\n", expected, actual)
		return
	}

	fmt.Printf("Correct %t\n", actual)
}

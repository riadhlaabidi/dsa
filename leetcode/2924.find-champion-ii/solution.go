package main

import "fmt"

func findChampion(n int, edges [][]int) int {
	teamIsWeak := make([]bool, n)

	for i := 0; i < len(edges); i++ {
		idx := edges[i][1]
		teamIsWeak[idx] = true
	}

	champions := 0
	championIdx := 0

	for i := 0; i < n; i++ {
		if !teamIsWeak[i] {
			champions++
			championIdx = i
		}

		if champions > 1 {
			return -1
		}
	}

	return championIdx
}

func main() {
	edges := [][]int{
		{0, 2},
		{1, 3},
		{1, 2},
	}

	expected := -1
	actual := findChampion(4, edges)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d\n", expected, actual)
		return
	}

	fmt.Printf("Correct %d\n", actual)
}

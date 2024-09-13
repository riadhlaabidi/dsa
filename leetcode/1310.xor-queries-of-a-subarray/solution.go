package main

import (
	"fmt"
	"slices"
)

func xorQueries(arr []int, queries [][]int) []int {

	for i := 1; i < len(arr); i++ {
		arr[i] ^= arr[i-1]
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		start := queries[i][0]
		end := queries[i][1]
		if start > 0 {
			ans[i] = arr[start-1] ^ arr[end]
		} else {
			ans[i] = arr[end]
		}
	}

	return ans
}

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}

	expected := []int{2, 7, 14, 8}
	actual := xorQueries(arr, queries)

	if !slices.Equal(actual, expected) {
		fmt.Printf("Expexted %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

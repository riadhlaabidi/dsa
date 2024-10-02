package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

type Pair struct {
	value int
	pos   int
}

func arrayRankTransform(arr []int) []int {
	sorted := make([]Pair, len(arr))

	for i, v := range arr {
		sorted[i] = Pair{v, i}
	}

	pairCmp := func(p1, p2 Pair) int {
		return cmp.Compare(p1.value, p2.value)
	}

	slices.SortFunc(sorted, pairCmp)

	ans := make([]int, len(arr))
	prev := math.MinInt
	rank := 0

	for _, item := range sorted {
		if item.value != prev {
			rank++
		}
		ans[item.pos] = rank
		prev = item.value
	}

	return ans
}

func main() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}

	expected := []int{5, 3, 4, 2, 8, 6, 7, 1, 3}
	actual := arrayRankTransform(arr)

	if !slices.Equal(actual, expected) {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

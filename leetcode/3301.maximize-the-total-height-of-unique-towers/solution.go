package main

import (
	"fmt"
	"slices"
)

func maximumTotalSum(maximumHeight []int) int64 {
	slices.Sort(maximumHeight)

	n := len(maximumHeight)
	prevMax := maximumHeight[n-1]
	ans := int64(prevMax)

	for i := n - 2; i >= 0; i-- {
		prevMax = min(prevMax-1, maximumHeight[i])

		if prevMax == 0 {
			return -1
		}

		ans += int64(prevMax)
	}

	return ans
}

func main() {
	maximumHeight := []int{2, 3, 4, 3}
	expected := int64(10)
	actual := maximumTotalSum(maximumHeight)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d", expected, actual)
		return
	}

	fmt.Printf("Correct %d", actual)
}

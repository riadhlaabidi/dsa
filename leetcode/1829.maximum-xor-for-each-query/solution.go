package main

import (
	"fmt"
	"slices"
)

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	ans := make([]int, n)

	numsXored := 0

	for _, num := range nums {
		numsXored ^= num
	}

	mask := (1 << maximumBit) - 1

	for i := range nums {
		ans[i] = numsXored ^ mask
		numsXored ^= nums[n-i-1]
	}

	return ans
}

func main() {
	nums := []int{0, 1, 1, 3}
	maximumBit := 2

	expected := []int{0, 3, 2, 3}
	actual := getMaximumXor(nums, maximumBit)

	if !slices.Equal(actual, expected) {
		fmt.Printf("Expected %v, but got %v\n", expected, actual)
		return
	}

	fmt.Printf("Correct %v\n", actual)
}

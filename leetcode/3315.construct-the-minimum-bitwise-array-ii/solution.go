package main

import (
	"fmt"
	"slices"
)

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, num := range nums {
		if num&1 == 0 {
			ans[i] = -1
			continue
		}

		for j := 0; j < 30; j++ {
			toUnset := 1 << j
			unsetMask := ^toUnset

			// num with j-th bit unset
			n := num & unsetMask

			if (n | (n + 1)) == num {
				ans[i] = n
			}

			if toUnset > num {
				break
			}
		}
	}

	return ans
}

func main() {
	input := []int{2, 3, 5, 7}

	expected := []int{-1, 1, 4, 3}
	actual := minBitwiseArray(input)

	if !slices.Equal(actual, expected) {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

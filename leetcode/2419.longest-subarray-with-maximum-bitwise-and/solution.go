package main

import "fmt"

func longestSubarray(nums []int) int {
	max := 0
	maxCount := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxCount = 1
			count = 1
		} else if nums[i] == max {
			count++
		} else {
			count = 0
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func main() {
	nums := []int{1, 2, 3, 3, 2, 2}
	expected := 2
	actual := longestSubarray(nums)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d", expected, actual)
	}

	fmt.Printf("Correct %d", actual)
}

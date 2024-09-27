package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	goal := n - 1

	for i := n - 2; i >= 0; i-- {
		if nums[i]+i >= goal {
			goal = i
		}
	}

	return goal == 0
}

func main() {
	nums := []int{2, 3, 1, 1, 4}

	expected := true
	actual := canJump(nums)

	if actual != expected {
		fmt.Printf("Expected %t, but got %t", expected, actual)
		return
	}

	fmt.Printf("Correct %t\n", actual)

	nums = []int{3, 2, 1, 0, 4}
	expected = false
	actual = canJump(nums)

	if actual != expected {
		fmt.Printf("Expected %t, but got %t", expected, actual)
		return
	}

	fmt.Printf("Correct %t", actual)
}

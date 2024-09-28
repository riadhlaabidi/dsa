package main

import (
	"fmt"
	"math"
)

func minElement(nums []int) int {
	min := math.MaxInt

	for _, num := range nums {
		sum := 0

		for num > 0 {
			sum += num % 10
			num /= 10
		}

		if sum < min {
			min = sum
		}
	}

	return min
}

func main() {
	nums := []int{10, 12, 13, 14}
	expected := 1
	actual := minElement(nums)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d", expected, actual)
		return
	}

	fmt.Printf("Correct %d", actual)
}

package main

import "fmt"

func minOperations(nums []int, k int) int {
	bits_diff := 0

	for i := 0; i < len(nums); i++ {
		bits_diff ^= nums[i]
	}

	count := 0
	bits_diff ^= k

	for bits_diff != 0 {
		count += (bits_diff & 1)
		bits_diff >>= 1
	}

	return count
}

func main() {
	arr := []int{2, 1, 3, 4}
	k := 1
	expected := 2

	actual := minOperations(arr, k)

	if actual != expected {
		fmt.Printf("Expected %d but found %d", expected, actual)
		return
	}

	fmt.Printf("Correct %d", actual)
}

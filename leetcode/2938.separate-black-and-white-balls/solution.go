package main

import "fmt"

func minimumSteps(s string) int64 {
	var ans int64
	var zeroCount int64

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			ans += zeroCount
		} else {
			zeroCount++
		}
	}

	return ans
}

func main() {
	input := "001001110001"

	expected := int64(14)
	actual := minimumSteps(input)

	if actual != expected {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

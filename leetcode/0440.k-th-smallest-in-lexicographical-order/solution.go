package main

import (
	"fmt"
)

func findKthNumber(n int, k int) int {
	ans := 1
	k--
	for k > 0 {
		steps := countSteps(n, ans, ans+1)
		if steps <= k {
			k -= steps
			ans++
		} else {
			ans *= 10
			k--
		}

	}

	return ans
}

func countSteps(n, prefix1, prefix2 int) (steps int) {
	for prefix1 <= n {
		steps += min(n+1, prefix2) - prefix1
		prefix1 *= 10
		prefix2 *= 10
	}
	return
}

func main() {
	n := 10000
	k := 10000

	expected := 9999
	actual := findKthNumber(n, k)

	if actual != expected {
		fmt.Printf("Failed. Expected %d but got %d", expected, actual)
		return
	}

	fmt.Printf("Correct %d", actual)
}

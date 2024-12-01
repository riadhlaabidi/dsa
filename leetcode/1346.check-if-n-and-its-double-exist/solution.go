package main

import "fmt"

func checkIfExist(arr []int) bool {
	seen := make(map[int]bool)

	for _, num := range arr {
		if seen[num*2] || (num%2 == 0 && seen[num>>1]) {
			return true
		}

		seen[num] = true
	}

	return false
}

func main() {
	arr := []int{10, 2, 5, 3}

	expected := true
	actual := checkIfExist(arr)

	if actual != expected {
		fmt.Printf("Expected %v, but got %v\n", expected, actual)
		return
	}

	fmt.Printf("Correct %v\n", actual)
}

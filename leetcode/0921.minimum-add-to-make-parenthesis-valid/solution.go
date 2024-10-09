package main

import "fmt"

func minAddToMakeValid(s string) int {
	opened := 0
	closed := 0

	for _, c := range s {
		if c == '(' {
			opened++
		} else {
			if opened > 0 {
				opened--
			} else {
				closed++
			}
		}
	}

	return opened + closed
}

func main() {
	s := "()))(("
	expected := 4
	actual := minAddToMakeValid(s)

	if actual != expected {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

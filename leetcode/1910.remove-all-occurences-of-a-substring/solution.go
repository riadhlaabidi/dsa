package main

import "fmt"

func removeOccurrences(s string, part string) string {
	st := make([]byte, 0)
	n := len(part)

	for i := 0; i < len(s); i++ {
		st = append(st, s[i])

		if len(st) >= n {
			if string(st[len(st)-n:]) == part {
				st = st[:len(st)-n]
			}
		}
	}

	return string(st)
}

func main() {
	s := "daabcbaabcbc"
	part := "abc"
	actual := removeOccurrences(s, part)
	expected := "dab"

	if actual != expected {
		fmt.Printf("Expected %s, but got %s\n", expected, actual)
		return
	}

	fmt.Printf("Correct %s\n", actual)
}

package main

import "fmt"

func addBinary(a string, b string) string {
	var ans string
	i := len(a) - 1
	j := len(b) - 1
	carry := 0

	for i >= 0 || j >= 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		ans = fmt.Sprintf("%c%s", (sum%2)+'0', ans)
		carry = sum >> 1
	}

	if carry == 1 {
		ans = fmt.Sprintf("1%s", ans)
	}

	return ans
}

func main() {
	actual := addBinary("1010", "1011")
	expected := "10101"

	if actual != expected {
		fmt.Printf("Expected %s, but got %s", expected, actual)
		return
	}

	fmt.Printf("Correct %s", actual)
}

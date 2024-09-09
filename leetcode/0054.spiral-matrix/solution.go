package main

import (
	"fmt"
	"slices"
)

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	ans := make([]int, m*n)

	top := 0
	bottom := m
	left := 0
	right := n
	k := 0

	for top < bottom && left < right {
		for i := left; i < right; i++ {
			ans[k] = matrix[top][i]
			k++
		}
		top++

		for i := top; i < bottom; i++ {
			ans[k] = matrix[i][right-1]
			k++
		}
		right--

		if left >= right || top >= bottom {
			break
		}

		for i := right - 1; i >= left; i-- {
			ans[k] = matrix[bottom-1][i]
			k++
		}
		bottom--

		for i := bottom - 1; i >= top; i-- {
			ans[k] = matrix[i][left]
			k++
		}
		left++
	}

	return ans
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	actual := spiralOrder(matrix)

	fmt.Printf("Result: %v\nCorrect: %t", actual, slices.Equal(actual, expected))

}

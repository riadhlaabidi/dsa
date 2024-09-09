package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
		for j := range n {
			ans[i][j] = -1
		}
	}

	row := 0
	col := 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0

	for head != nil {
		ans[row][col] = head.Val

		if dr, dc := dirs[d][0], dirs[d][1]; row+dr >= m || col+dc >= n || col+dc < 0 || ans[row+dr][col+dc] != -1 {
			d = (d + 1) % 4
		}

		row += dirs[d][0]
		col += dirs[d][1]

		head = head.Next
	}

	return ans
}

func main() {
	arr := []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}
	const m = 3
	const n = 5

	head := ListNode{arr[0], nil}
	tmp := &head

	for i := 1; i < len(arr); i++ {
		newNode := ListNode{arr[i], nil}
		tmp.Next = &newNode
		tmp = &newNode
	}

	ans := spiralMatrix(m, n, &head)

	fmt.Println(ans)
}

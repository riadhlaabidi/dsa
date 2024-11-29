package main

import (
	"container/heap"
	"fmt"
)

type Cell struct {
	time int /* time to reach the cell */
	r    int /* row */
	c    int /* column */
}

type MinHeap []Cell

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].time < mh[j].time
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(Cell))
}

func (mh *MinHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func minimumTime(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	m := len(grid)
	n := len(grid[0])

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var mh MinHeap
	heap.Init(&mh)
	heap.Push(&mh, Cell{0, 0, 0})

	for mh.Len() != 0 {
		curr := heap.Pop(&mh).(Cell)
		time := curr.time
		r := curr.r
		c := curr.c

		if r == m-1 && c == n-1 {
			return time
		}

		if visited[r][c] {
			continue
		}

		visited[r][c] = true

		for _, dir := range dirs {
			newR := r + dir[0]
			newC := c + dir[1]

			if !isInBounds(m, n, newR, newC) || visited[newR][newC] {
				continue
			}

			waste := 0
			if (grid[newR][newC]-time)%2 == 0 {
				waste = 1
			}

			newTime := max(grid[newR][newC]+waste, time+1)
			heap.Push(&mh, Cell{newTime, newR, newC})
		}
	}

	return -1
}

func isInBounds(m, n, newR, newC int) bool {
	return newR >= 0 && newR < m && newC >= 0 && newC < n
}

func main() {
	grid := [][]int{
		{0, 1, 3, 2},
		{5, 1, 2, 5},
		{4, 3, 8, 6},
	}

	expected := 7
	actual := minimumTime(grid)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d\n", expected, actual)
		return
	}

	fmt.Printf("Correct %d\n", actual)
}

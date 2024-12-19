package main

import (
	"aoc/util"
	"container/heap"
	"fmt"
	"math"
)

type Cell struct {
	direction int
	distance  int
	r         int
	c         int
}

type MinHeap []Cell

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].distance < mh[j].distance
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(Cell))
}

func (mh *MinHeap) Pop() any {
	old := *mh
	l := len(old)
	x := old[l-1]
	*mh = old[0 : l-1]
	return x
}

func getAdjacentThree(grid [][]byte, cell Cell) []Cell {
	adjs := make([]Cell, 0)

	dirs := []int{
		cell.direction,           // same direction
		(cell.direction + 1) % 4, // perpendicular direction
		(cell.direction + 3) % 4, // reverse perpendicular direction
	}

	for _, dir := range dirs {
		nextR := cell.r + util.Directions[dir][0]
		nextC := cell.c + util.Directions[dir][1]
		if grid[nextR][nextC] != '#' {
			adjs = append(adjs, Cell{r: nextR, c: nextC, direction: dir})
		}
	}

	return adjs
}

func solvePartI(lines [][]byte) {
	m := len(lines)
	n := len(lines[0])
	sr, sc := len(lines)-2, 1
	er, ec := 1, len(lines[0])-2
	distances := make([][]int, m)

	for i := range distances {
		distances[i] = make([]int, n)
		for j := range distances[i] {
			distances[i][j] = math.MaxInt
		}
	}
	distances[sr][sc] = 0

	var mh MinHeap
	heap.Init(&mh)
	heap.Push(&mh, Cell{0, 0, sr, sc})

	for len(mh) != 0 {
		src := heap.Pop(&mh).(Cell)
		if src.r == er && src.c == ec {
			break
		}
		for _, adj := range getAdjacentThree(lines, src) {
			lines[adj.r][adj.c] = 'V'
			newDistance := src.distance + 1
			if adj.direction != src.direction {
				newDistance += 1000
			}
			if newDistance < distances[adj.r][adj.c] {
				distances[adj.r][adj.c] = newDistance
				adj.distance = newDistance
				heap.Push(&mh, adj)
			}
		}
	}

	fmt.Printf("Part 1: %v\n", distances[er][ec])
}

func solvePartII(lines [][]byte) {
	ans := 0
	fmt.Printf("Part 2: %v\n", ans)
}

func main() {
	testLines := util.ReadInput("day16.test")
	inputLines := util.ReadInput("day16.input")
	fmt.Printf("---- Day 16 Test ----\n")
	solvePartI(testLines)
	solvePartII(testLines)
	fmt.Printf("---- Day 16 Input ----\n")
	solvePartI(inputLines)
	solvePartII(inputLines)
}

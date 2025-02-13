package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *MinHeap) Pop() any {
	old := *mh
	n := len(old) - 1
	x := old[n]
	*mh = old[:n]
	return x
}

func minOperations(nums []int, k int) int {
	pq := make(MinHeap, len(nums))

	for i, num := range nums {
		pq[i] = num
	}
	heap.Init(&pq)

	ops := 0
	for pq[0] < k {
		ops++
		a := heap.Pop(&pq).(int)
		b := heap.Pop(&pq).(int)
		newValue := min(a, b)*2 + max(a, b)
		heap.Push(&pq, newValue)
	}

	return ops
}

func main() {
	nums := []int{2, 11, 10, 1, 3}
	k := 10
	expected := 2
	actual := minOperations(nums, k)

	if actual != expected {
		fmt.Printf("Expected %d, but got %d\n", expected, actual)
		return
	}

	fmt.Printf("Correct %d\n", actual)
}

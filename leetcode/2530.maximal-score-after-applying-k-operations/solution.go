package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MaxHeap []int

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MaxHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *MaxHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	mh := MaxHeap(nums)
	heap.Init(&mh)

	var score int64

	for k > 0 {
		max := heap.Pop(&mh).(int)
		score += int64(max)
		max = int(math.Ceil(float64(max) / 3))
		heap.Push(&mh, max)
		k--
	}

	return score
}

func main() {
	nums := []int{1, 10, 3, 3, 3}
	k := 3
	expected := int64(17)
	actual := maxKelements(nums, k)

	if actual != expected {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

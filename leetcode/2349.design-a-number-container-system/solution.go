package main

import (
	"container/heap"
	"fmt"
	"log"
	"runtime"
)

type MinHeap []int

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

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

type NumberContainers struct {
	numberIndex map[int]*MinHeap
	indexNumber map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		numberIndex: make(map[int]*MinHeap),
		indexNumber: make(map[int]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if _, exists := this.numberIndex[number]; !exists {
		mh := make(MinHeap, 0)
		this.numberIndex[number] = &mh
	}
	heap.Push(this.numberIndex[number], index)
	this.indexNumber[index] = number
}

func (this *NumberContainers) Find(number int) int {
	if indexes, exists := this.numberIndex[number]; exists {
		for indexes.Len() > 0 {
			idx := (*indexes)[0]
			if this.indexNumber[idx] == number {
				return idx
			}
			heap.Pop(indexes)
		}
	}
	return -1
}

func assertEquals(actual, expected int) {
	if actual == expected {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("%s:%d: Expected %d, but got %d\n", file, line, expected, actual)
	log.Fatal(msg)
}

func main() {
	container := Constructor()
	assertEquals(container.Find(10), -1)
	container.Change(2, 10)
	container.Change(1, 10)
	container.Change(3, 10)
	container.Change(5, 10)
	assertEquals(container.Find(10), 1)
	container.Change(1, 20)
	assertEquals(container.Find(10), 2)
}

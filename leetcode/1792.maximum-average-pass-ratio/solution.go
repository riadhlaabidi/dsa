package main

import "container/heap"

type Class struct {
	pass              int
	total             int
	passRatioIncrease float64
}

type MaxHeap []Class

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(x, y int) bool {
	return mh[x].passRatioIncrease > mh[y].passRatioIncrease
}

func (mh MaxHeap) Swap(x, y int) {
	mh[x], mh[y] = mh[y], mh[x]
}

func (mh *MaxHeap) Push(x any) {
	*mh = append(*mh, x.(Class))
}

func (mh *MaxHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, exrtaStudents int) float64 {
	n := len(classes)
	mh := make(MaxHeap, n)

	passRatioSum := float64(0)

	for i, class := range classes {
		pass := class[0]
		total := class[1]
		passRatio := float64(pass) / float64(total)
		passRatioSum += passRatio
		pass++
		total++
		increase := float64(pass)/float64(total) - passRatio
		mh[i] = Class{pass: pass, total: total, passRatioIncrease: increase}
	}

	heap.Init(&mh)

	for range exrtaStudents {
		maxIncrease := heap.Pop(&mh).(Class)
		passRatioSum += maxIncrease.passRatioIncrease
		maxIncrease.passRatioIncrease = -float64(maxIncrease.pass) / float64(maxIncrease.total)
		maxIncrease.pass++
		maxIncrease.total++
		maxIncrease.passRatioIncrease += float64(maxIncrease.pass) / float64(maxIncrease.total)
		heap.Push(&mh, maxIncrease)
	}

	return passRatioSum / float64(n)
}

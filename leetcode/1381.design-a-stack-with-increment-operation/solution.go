package main

import (
	"fmt"
	"slices"
)

type CustomStack struct {
	values   []int
	capacity int
}

func NewCustomStack(maxSize int) CustomStack {
	return CustomStack{capacity: maxSize}
}

func (s *CustomStack) Push(val int) {
	if len(s.values) >= s.capacity {
		return
	}

	s.values = append(s.values, val)
}

func (s *CustomStack) Pop() int {
	n := len(s.values)

	if n == 0 {
		return -1
	}

	val := s.values[n-1]
	s.values = s.values[:n-1]

	return val
}

func (s *CustomStack) Increment(k int, val int) {
	for i := range s.values {
		if i >= k {
			break
		}
		s.values[i] += val
	}
}

func main() {
	actual := NewCustomStack(3)
	actual.Push(1)
	actual.Push(2)
	actual.Pop()
	actual.Push(2)
	actual.Push(3)
	actual.Push(4)
	actual.Increment(5, 100)
	actual.Increment(2, 100)

	expected := CustomStack{capacity: 3, values: []int{201, 202, 103}}

	if !slices.Equal(expected.values, actual.values) || actual.capacity != expected.capacity {
		fmt.Printf("Increment: Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Increment: Correct %v\n", actual)

	actual.Pop()
	actual.Pop()
	actual.Pop()
	actual.Pop()

	if len(actual.values) > 0 {
		fmt.Printf("Popping: Expected %v, but got %v", []int{}, actual)
	}

	fmt.Printf("Popping: Correct %v", actual)
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	tmp := head

	for tmp != nil {
		next := tmp.Next
		tmp.Next = newHead
		newHead = tmp
		tmp = next
	}

	return newHead
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := ListNode{arr[0], nil}
	tmp := &head

	for i := 1; i < 5; i++ {
		newNode := ListNode{arr[i], nil}
		tmp.Next = &newNode
		tmp = tmp.Next
	}

	expected := []int{5, 4, 3, 2, 1}
	actual := reverseList(&head)

	i := 0
	for actual != nil {
		if actual.Val != expected[i] {
			fmt.Printf("Expected %d in position %d but found %d", expected[i], i, actual.Val)
			return
		}
		i++
		actual = actual.Next
	}

	fmt.Println("Correct")
}

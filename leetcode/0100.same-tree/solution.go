package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{1, nil, nil}
	p.Left = &TreeNode{2, nil, nil}
	p.Right = &TreeNode{3, nil, nil}

	q := &TreeNode{1, nil, nil}
	q.Left = &TreeNode{2, nil, nil}
	q.Right = &TreeNode{3, nil, nil}

	expected := true
	actual := isSameTree(p, q)

	if actual != expected {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

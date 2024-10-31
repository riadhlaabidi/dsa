package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return walk(root.Left, root.Right)
}

func walk(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return walk(root1.Left, root2.Right) && walk(root1.Right, root2.Left)
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{4, nil, nil}
	root.Right.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{4, nil, nil}

	expected := true
	actual := isSymmetric(root)

	if actual != expected {
		fmt.Printf("Expected %t, but got %t\n", expected, actual)
		return
	}

	fmt.Printf("Correct %t\n", actual)
}

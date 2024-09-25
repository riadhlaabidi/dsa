package main

import (
	"fmt"
	"slices"
)

type TrieNode struct {
	next  []*TrieNode
	count int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{next: make([]*TrieNode, 26)}
}

func sumPrefixScores(words []string) []int {
	ans := make([]int, len(words))
	root := NewTrieNode()

	for _, s := range words {
		insertWord(root, s)
	}

	for i, w := range words {
		ans[i] = countScore(root, w)
	}

	return ans
}

func insertWord(root *TrieNode, w string) {
	curr := root

	for _, char := range w {
		idx := char - 'a'

		if curr.next[idx] == nil {
			curr.next[idx] = NewTrieNode()
		}

		curr.next[idx].count++
		curr = curr.next[idx]
	}
}

func countScore(root *TrieNode, word string) int {
	curr := root
	count := 0

	for _, char := range word {
		idx := char - 'a'

		if curr.next[idx] == nil {
			break
		}

		count += curr.next[idx].count
		curr = curr.next[idx]
	}

	return count
}

func main() {
	words := []string{"abc", "ab", "bc", "b"}

	expected := []int{5, 4, 3, 2}
	actual := sumPrefixScores(words)

	if !slices.Equal(actual, expected) {
		fmt.Printf("Expected %v, but got %v", expected, actual)
	}

	fmt.Printf("Correct %v", actual)
}

package main

import (
	"fmt"
	"slices"
	"strings"
)

type TrieNode struct {
	isEndOfPath bool
	children    map[string]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{false, make(map[string]*TrieNode)}
}

func removeSubFolders(folder []string) []string {
	root := NewTrieNode()

	for _, path := range folder {
		current := root

		for _, c := range strings.Split(path, "/")[1:] {
			if _, exists := current.children[c]; !exists {
				current.children[c] = NewTrieNode()
			}
			current = current.children[c]
		}

		current.isEndOfPath = true
	}

	res := []string{}

	for _, path := range folder {
		isSubFolder := false
		current := root
		folderNames := strings.Split(path, "/")[1:]

		for i, c := range folderNames {
			next := current.children[c]
			if next.isEndOfPath && i != len(folderNames)-1 {
				isSubFolder = true
				break
			}
			current = next
		}

		if !isSubFolder {
			res = append(res, path)
		}
	}

	return res
}

func main() {
	folders := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	expected := []string{"/a", "/c/d", "/c/f"}
	actual := removeSubFolders(folders)

	if !slices.Equal(actual, expected) {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}

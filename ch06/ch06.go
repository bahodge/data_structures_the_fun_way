package ch06

import "fmt"

type TrieNode struct {
	IsEntry  bool
	Children []*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Run() {
	fmt.Println("----------------ch06 start-----------------")

	fmt.Println("hello world")
	fmt.Println("----------------ch06 end-----------------")
}

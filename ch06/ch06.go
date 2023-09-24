package ch06

import (
	"fmt"
	"strings"
)

type Node struct {
	Char     string
	Children [26]*Node
}

func NewNode(char string) *Node {
	node := &Node{Char: char}
	for i := 0; i < 26; i++ {
		node.Children[i] = nil
	}
	return node
}

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	root := NewNode("\000")
	return &Trie{RootNode: root}
}

func (t *Trie) Insert(word string) error {
	current := t.RootNode
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(strippedWord[i]))
		}
		current = current.Children[index]
	}
	return nil
}

func (t *Trie) Search(word string) bool {
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	current := t.RootNode
	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'
		if current == nil || current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	return true
}

func Run() {
	fmt.Println("----------------ch06 start-----------------")
	tr := NewTrie()

	words := []string{"ben", "steven", "sardyboy", "samantha", "sardy", "bob", "sam"}
	allWords := []string{"not  ", "in ", "the", "trie"}

	for _, word := range words {
		tr.Insert(word)
		allWords = append(allWords, word)
	}

	for _, word := range allWords {
		fmt.Printf("search result for %s: %t\n", word, tr.Search(word))
	}

	fmt.Println("----------------ch06 end-----------------")
}

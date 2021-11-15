package trie

import (
	"strings"
)

type trieNode struct {
	letter      byte
	children    map[byte]*trieNode
	isEndOfWord bool
}

type Trie struct {
	rootNode *trieNode
}

func New() *Trie {
	return &Trie{rootNode: &trieNode{letter: 0, children: make(map[byte]*trieNode)}}
}

// AddWord : Add a new word to the trie
func (t Trie) AddWord(word string) {
	currentNode := t.rootNode

	for _, letter := range strings.ToLower(word) {
		childNode := currentNode.children[byte(letter)]
		if childNode == nil {
			childNode = &trieNode{letter: byte(letter), children: make(map[byte]*trieNode)}
			currentNode.children[byte(letter)] = childNode
		}
		currentNode = childNode
	}
	currentNode.isEndOfWord = true
}

// FindWord : Traverse the trie to search for a given word
func (t Trie) FindWord(word string) bool {
	currentNode := t.rootNode

	for _, letter := range strings.ToLower(word) {
		currentNode = currentNode.children[byte(letter)]
		if currentNode == nil {
			return false
		}
	}
	return currentNode.isEndOfWord
}

// FindWordsByPrefix : Find all words with a common specified prefix
func (t Trie) FindWordsByPrefix(prefix string) []string {
	return []string{""}
}

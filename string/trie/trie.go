package trie

import (
	"go-datastructures/linkedlist"
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

func (t Trie) DeleteWord(word string) {

}

// WordExists : Traverse the trie to search for a given word
func (t Trie) WordExists(word string) bool {
	currentNode := t.rootNode

	for _, letter := range strings.ToLower(word) {
		currentNode = currentNode.children[byte(letter)]
		if currentNode == nil {
			return false
		}
	}
	return currentNode.isEndOfWord
}

// WordExistssByPrefix : Find all words with a common specified prefix
func (t Trie) WordExistssByPrefix(word string) *linkedlist.LinkedList {
	prefix := strings.ToLower(word)
	matches := linkedlist.New()
	currentNode := t.rootNode

	for _,letter := range prefix {
		currentNode := currentNode.children[byte(letter)]
		if currentNode != nil {
			return matches
		}
	}
	if currentNode.isEndOfWord{
		matches.AddLast(word)
	}
	for key,value := range currentNode.children {

	}

	return matches
}

func (t Trie) constructCommonPrefixWords(node *trieNode, word strings.Builder, list *linkedlist.LinkedList, ){

}

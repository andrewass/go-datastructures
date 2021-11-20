package trie

import (
	"go-datastructures-algorithms/list/arraylist"
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

// DeleteWord : Delete a word from the Trie
func (t Trie) DeleteWord(word string) {
	t.delete(t.rootNode, 0, word)
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

// WordExistsByPrefix : Find all words with a common specified prefix
func (t Trie) WordExistsByPrefix(word string) *arraylist.ArrayList {
	prefix := strings.ToLower(word)
	matches := arraylist.New()
	currentNode := t.rootNode

	for _, letter := range prefix {
		currentNode = currentNode.children[byte(letter)]
		if currentNode == nil {
			return matches
		}
	}
	if currentNode.isEndOfWord {
		matches.Add(word)
	}
	for _, value := range currentNode.children {
		builder := &strings.Builder{}
		builder.WriteString(word)
		t.constructCommonPrefixWords(value, builder, matches)
	}
	return matches
}

// Construct all words sharing a common prefix
func (t Trie) constructCommonPrefixWords(node *trieNode, word *strings.Builder, matches *arraylist.ArrayList) {
	word.WriteByte(node.letter)
	if node.isEndOfWord {
		matches.Add(word.String())
	}
	for _, value := range node.children {
		t.constructCommonPrefixWords(value, word, matches)
	}
}

func (t Trie) delete(node *trieNode, index int, word string) bool {
	if index == len(word) {
		if !node.isEndOfWord {
			return false

		} else {
			node.isEndOfWord = false
			return len(node.children) == 0
		}
	}
	childNode := node.children[word[index]]
	if childNode == nil {
		return false
	}

	deletableChild := t.delete(childNode, index+1, word) && !childNode.isEndOfWord
	if deletableChild {
		delete(node.children, word[index])
		return len(node.children) == 0
	} else {
		return false
	}
}

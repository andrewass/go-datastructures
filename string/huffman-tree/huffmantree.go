package huffman_tree

import (
	"bytes"
	priorityqueue "go-datastructures/queue/priority-queue"
)

type HuffmanTree struct {
	word           string
	queue          *priorityqueue.PriorityQueue
	encodingLength int
	encodeMap      map[byte]string
	root           *node
}

type node struct {
	char  byte
	freq  int
	left  *node
	right *node
}

func nodeComparator(a, b interface{}) int {
	taskA := a.(*node)
	taskB := b.(*node)
	if taskA.freq < taskB.freq {
		return -1
	} else if taskA.freq > taskB.freq {
		return 1
	} else {
		return 0
	}
}

func New(word string) *HuffmanTree {
	huffmanTree := &HuffmanTree{
		word:      word,
		queue:     priorityqueue.New(nodeComparator),
		encodeMap: make(map[byte]string),
	}
	huffmanTree.populatePriorityQueue()
	huffmanTree.buildTree()
	huffmanTree.encodeWord(huffmanTree.root, &bytes.Buffer{})
	return huffmanTree
}

func (ht *HuffmanTree) buildTree() {
	if len(ht.word) == 1 {
		ht.queue.Insert(&node{freq: 0})
	}
	for ht.queue.Size() > 1 {
		firstNode := ht.queue.PollMin().(*node)
		secondNode := ht.queue.PollMin().(*node)
		newNode := &node{freq: firstNode.freq + secondNode.freq}
		newNode.left = firstNode
		newNode.right = secondNode
		ht.queue.Insert(newNode)
	}
	ht.root = ht.queue.PollMin().(*node)
}

func (ht *HuffmanTree) populatePriorityQueue() {
	freqMap := make(map[byte]int)
	for _, char := range ht.word {
		freqMap[byte(char)]++
	}
	for key, val := range freqMap {
		ht.queue.Insert(&node{char: key, freq: val})
	}
}

func (ht *HuffmanTree) encodeWord(node *node, code *bytes.Buffer) {
	if node != nil {
		if node.char != 0 {
			ht.encodingLength += node.freq + code.Len()
			ht.encodeMap[node.char] = code.String()
		} else {
			code.WriteString("0")
			ht.encodeWord(node.left, code)
			code.Truncate(code.Len() - 1)

			code.WriteString("1")
			ht.encodeWord(node.right, code)
			code.Truncate(code.Len() - 1)
		}
	}
}

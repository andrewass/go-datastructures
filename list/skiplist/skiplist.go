package skiplist

import (
	"go-datastructures-algorithms/util/comparator"
	"math/rand"
)

type node struct {
	key       interface{}
	item      interface{}
	nodeLevel int
	isHead    bool
	next      *node
	above     *node
	under     *node
	column    []*node
}

type SkipList struct {
	head        *node
	size        int
	maxLevel    int
	probability float32
	comparator  comparator.Comparator
}

func New() *SkipList {
	SkipList := &SkipList{
		maxLevel:    32,
		probability: 0.5,
	}
	headNodes := make([]*node, SkipList.maxLevel)
	for i := 0; i < SkipList.maxLevel; i++ {
		headNodes[i] = &node{key: nil, item: nil, isHead: true}
	}
	SkipList.head = headNodes[0]
	SkipList.head.column = headNodes

	return SkipList
}

func (sl *SkipList) Insert(key, item interface{}) {
	//nodeLevel := sl.generateNodeLevel()
	 //:= &node{item: item, key: key, nodeLevel: nodeLevel}

	//start with the
	currentNode := sl.head
	for i := 0; i < sl.maxLevel; i++ {
		currentNode = currentNode.column[i]
		for currentNode.next != nil && sl.comparator(key, currentNode.next.key) < 0 {
			currentNode = currentNode.next
		}
	}
}

func (sl *SkipList) generateNodeLevel() int {
	level := 1
	for level < sl.maxLevel && rand.Float32() <= sl.probability {
		level++
	}
	return level
}

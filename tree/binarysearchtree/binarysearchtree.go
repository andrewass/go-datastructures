package binarysearchtree

import "go-datastructures/util"

type node struct {
	item       interface{}
	key        interface{}
	parent     *node
	leftChild  *node
	rightChild *node
}

type BinarySearchTree struct {
	root       *node
	comparator util.Comparator
	size       int
}

// New : create a new BinarySearchTree, by specifying its comparator
func New(comparator util.Comparator) *BinarySearchTree  {
	return &BinarySearchTree{comparator: comparator}
}


//Insert : insert an item to the tree, wrapped in a new node
func (bst *BinarySearchTree) Insert(key interface{}, item interface{}) {
	newNode := &node{key: key, item: item}
	currentNode := bst.root
	var leafParent *node

	for currentNode != nil {
		leafParent = currentNode
		if bst.comparator(newNode.item, currentNode.item) < 0 {
			currentNode = currentNode.leftChild
		} else {
			currentNode = currentNode.rightChild
		}
	}
	newNode.parent = leafParent

	if leafParent == nil {
		bst.root = newNode
	} else if bst.comparator(newNode.item, leafParent.item) < 0 {
		leafParent.leftChild = newNode
	} else {
		leafParent.rightChild = newNode
	}
	bst.size++
}

//Delete : remove an item from the tree, along with the wrapping node
func (bst *BinarySearchTree) Delete(key interface{}) {
}

func (bst *BinarySearchTree) Find(key interface{}) interface{} {
	return nil
}

// Minimum : find the item with the lowest comparable value in the tree
func (bst *BinarySearchTree) Minimum() interface{} {
	currentNode := bst.root
	for currentNode != nil && currentNode.leftChild != nil {
		currentNode = currentNode.leftChild
	}
	return currentNode.item
}

// Maximum : find the item with the greatest comparable value in the tree
func (bst *BinarySearchTree) Maximum() interface{} {
	currentNode := bst.root
	for currentNode != nil && currentNode.rightChild != nil {
		currentNode = currentNode.rightChild
	}
	return currentNode.item
}

func (bst *BinarySearchTree) InOrderList() []interface{} {
	return nil
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

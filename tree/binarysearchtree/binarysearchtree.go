package binarysearchtree

import (
	"go-datastructures-algorithms/list/linkedlist"
	"go-datastructures-algorithms/util/comparator"
)

type node struct {
	item       interface{}
	key        interface{}
	parent     *node
	leftChild  *node
	rightChild *node
}

type BinarySearchTree struct {
	root       *node
	comparator comparator.Comparator
	size       int
}

// New : create a new BinarySearchTree, by specifying its comparator
func New(comparator comparator.Comparator) *BinarySearchTree {
	return &BinarySearchTree{comparator: comparator}
}

// Insert : insert an item to the tree, where its position in the tree is based on its key value
func (bst *BinarySearchTree) Insert(key interface{}, item interface{}) {
	newNode := &node{key: key, item: item}
	currentNode := bst.root
	var leafParent *node

	for currentNode != nil {
		leafParent = currentNode
		if bst.comparator(newNode.key, currentNode.key) < 0 {
			currentNode = currentNode.leftChild
		} else {
			currentNode = currentNode.rightChild
		}
	}
	newNode.parent = leafParent

	if leafParent == nil {
		bst.root = newNode
	} else if bst.comparator(newNode.key, leafParent.key) < 0 {
		leafParent.leftChild = newNode
	} else {
		leafParent.rightChild = newNode
	}
	bst.size++
}

// Delete : delete an item from the tree, based in its key value
func (bst *BinarySearchTree) Delete(key interface{}) {
	deleteNode := bst.findNode(key)
	if deleteNode != nil {
		if deleteNode.leftChild == nil {
			bst.transplant(deleteNode, deleteNode.rightChild)
		} else if deleteNode.rightChild == nil {
			bst.transplant(deleteNode, deleteNode.leftChild)
		} else {
			bst.deleteNodeWithTwoChildren(deleteNode)
		}
		bst.size--
	}
}

func (bst *BinarySearchTree) deleteNodeWithTwoChildren(deleteNode *node) {
	minRightNode := bst.findMinimum(deleteNode.rightChild)
	if minRightNode.parent != deleteNode {
		bst.transplant(minRightNode, minRightNode.rightChild)
		minRightNode.rightChild = deleteNode.rightChild
		minRightNode.rightChild.parent = minRightNode
	}
	bst.transplant(deleteNode, minRightNode)
	minRightNode.leftChild = deleteNode.leftChild
	minRightNode.leftChild.parent = minRightNode
}

// Removes the linking between a parent and child, and creates a direct
// link between the child and its previous grandparent
func (bst *BinarySearchTree) transplant(parent *node, child *node) {
	if parent.parent == nil {
		parent = child
	} else if parent == parent.parent.leftChild {
		parent.parent.leftChild = child
	} else {
		parent.parent.rightChild = child
	}

	if child != nil {
		child.parent = parent.parent
	}
}

// Find : search the tree for any node containing the given key.
// If exists return its item, else return nil
func (bst *BinarySearchTree) Find(key interface{}) interface{} {
	foundNode := bst.findNode(key)
	if foundNode != nil {
		return foundNode.item
	}
	return nil
}

func (bst *BinarySearchTree) findNode(key interface{}) *node {
	currentNode := bst.root
	for currentNode != nil {
		compVal := bst.comparator(key, currentNode.key)
		if compVal == 0 {
			return currentNode
		} else if compVal < 0 {
			currentNode = currentNode.leftChild
		} else {
			currentNode = currentNode.rightChild
		}
	}
	return nil
}

func (bst *BinarySearchTree) findMinimum(node *node) *node {
	currentNode := node
	for currentNode != nil && currentNode.leftChild != nil {
		currentNode = currentNode.leftChild
	}
	return currentNode
}

func (bst *BinarySearchTree) findMaximum(node *node) *node {
	currentNode := node
	for currentNode != nil && currentNode.rightChild != nil {
		currentNode = currentNode.rightChild
	}
	return currentNode
}

// Minimum : find the item with the lowest comparable value in the tree
func (bst *BinarySearchTree) Minimum() interface{} {
	minimumNode := bst.findMinimum(bst.root)
	if minimumNode != nil {
		return minimumNode.item
	}
	return nil
}

// Maximum : find the item with the greatest comparable value in the tree
func (bst *BinarySearchTree) Maximum() interface{} {
	maximumNode := bst.findMaximum(bst.root)
	if maximumNode != nil {
		return maximumNode.item
	}
	return nil
}

// InOrderList : returns a sorted list of all items in the tree
func (bst *BinarySearchTree) InOrderList() *linkedlist.LinkedList {
	inOrderList := linkedlist.New()
	bst.fillInOrderList(bst.root, inOrderList)

	return inOrderList
}

func (bst *BinarySearchTree) fillInOrderList(node *node, list *linkedlist.LinkedList) {
	if node != nil {
		bst.fillInOrderList(node.leftChild, list)
		list.AddLast(node.item)
		bst.fillInOrderList(node.rightChild, list)
	}
}

// Size : returns the number of items in the tree
func (bst *BinarySearchTree) Size() int {
	return bst.size
}

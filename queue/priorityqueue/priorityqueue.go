package priorityqueue

import (
	"go-datastructures/util/comparator"
)

type PriorityQueue struct {
	heap        []interface{}
	comparator  comparator.Comparator
	currentSize int
}

func New(comparator comparator.Comparator) *PriorityQueue {
	return &PriorityQueue{
		heap:        make([]interface{},1),
		currentSize: 0,
		comparator:  comparator}
}

func (pq *PriorityQueue) PeekMin() interface{} {
	if pq.currentSize == 0 {
		return nil
	}
	return pq.heap[1]
}

func (pq *PriorityQueue) PollMin() interface{} {
	if pq.currentSize == 0 {
		return nil
	}
	response := pq.heap[1]
	pq.deleteMin()
	return response
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.currentSize == 0
}

func (pq *PriorityQueue) Size() int {
	return pq.currentSize
}

func (pq *PriorityQueue) Insert(node interface{}) {
	if pq.currentSize == len(pq.heap)-1 {
		pq.heap = append(pq.heap, nil)
	}
	pq.currentSize++
	pq.heap[pq.currentSize] = node
	slot := pq.currentSize

	for slot > 1 {
		parent := slot/2
		if pq.comparator(pq.heap[parent], pq.heap[slot]) <= 0 {
			break
		}
		pq.heap[slot], pq.heap[parent] = pq.heap[parent], pq.heap[slot]
		slot = parent
	}
}


func (pq *PriorityQueue) deleteMin() {
	pq.heap[1] = pq.heap[pq.currentSize]
	pq.heap[pq.currentSize] = nil
	pq.currentSize--
	slot := 1

	for slot < pq.currentSize {
		minChild := pq.getMinChild(slot)
		if minChild != slot {
			pq.heap[slot], pq.heap[minChild] = pq.heap[minChild], pq.heap[slot]
			slot = minChild
		} else {
			break
		}
	}
}

func (pq *PriorityQueue) getMinChild(slot int) int {
	leftChild := slot * 2
	rightChild := leftChild + 1
	smallest := slot

	if leftChild <= pq.currentSize && pq.comparator(pq.heap[leftChild], pq.heap[smallest]) < 0 {
		smallest = leftChild
	}
	if rightChild <= pq.currentSize && pq.comparator(pq.heap[rightChild], pq.heap[smallest]) < 0 {
		smallest = rightChild
	}
	return smallest
}

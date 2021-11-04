package priority_queue

import (
	"go-datastructures/util"
)

type PriorityQueue struct {
	heap       []interface{}
	comparator util.Comparator
	size       int
}

func New(comparator util.Comparator) *PriorityQueue {
	return &PriorityQueue{
		heap:       make([]interface{}, 2),
		size:       0,
		comparator: comparator}
}

func (pq *PriorityQueue) Peek() interface{} {
	if pq.size == 0 {
		return nil
	}
	return pq.heap[0]
}

func (pq *PriorityQueue) Poll() interface{} {
	if pq.size == 0 {
		return nil
	}
	response := pq.heap[1]

	return response
}

func (pq *PriorityQueue) Insert(node interface{}) {
}

func (pq *PriorityQueue) heapify(index int) {
	leftChild := index * 2
	rightChild := index*2 + 1
	smallest := index

	if leftChild <= pq.size && pq.comparator(pq.heap[leftChild], pq.heap[index]) < 0 {
		smallest = leftChild
	}
	if rightChild <= pq.size && pq.comparator(pq.heap[rightChild], pq.heap[smallest]) < 0 {
		smallest = rightChild
	}
	if smallest != pq.heap[index] {
	}
}

func (pq *PriorityQueue) swap() {

}

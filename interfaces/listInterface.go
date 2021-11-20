package interfaces

import "go-datastructures-algorithms/util/comparator"

type ListInterface interface {
	Get(index int) interface{}
	Add(interface{})
	Remove(index int)
	Swap(indexA, indexB int)
	Replace(index int, item interface{})
	Size() int
	Sort(comparator comparator.Comparator)
}

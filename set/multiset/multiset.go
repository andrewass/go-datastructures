package multiset

import (
	"go-datastructures-algorithms/set/hashset"
)

type MultiSet struct {
	set          *hashset.HashSet
	countMap     map[interface{}]int
	totalSize    int
	distinctSize int
}

//New create a new MultiSet
func New() *MultiSet {
	return &MultiSet{set: hashset.New(), countMap: make(map[interface{}]int)}
}

//Add inserts an item into the MultiSet. This item may already exist in the set,
//which only results in an increase in its count
func (ms *MultiSet) Add(item interface{}) {
	existingCount := ms.countMap[item]
	if existingCount > 0 {
		ms.countMap[item] = existingCount + 1
	} else {
		ms.set.Add(item)
		ms.countMap[item] = 1
		ms.distinctSize++
	}
	ms.totalSize++
}

//Remove deletes an occurrence of given item from the MultiSet
func (ms *MultiSet) Remove(item interface{}) {
	existingCount := ms.countMap[item]
	if existingCount > 1 {
		ms.countMap[item] = existingCount - 1
		ms.totalSize--
	} else if existingCount == 1 {
		delete(ms.countMap, item)
		ms.set.Remove(item)
		ms.totalSize--
		ms.distinctSize--
	}
}

//RemoveAll deletes all occurrences of given item from the MultiSet
func (ms *MultiSet) RemoveAll(item interface{}) {
	existingCount := ms.countMap[item]
	if existingCount > 1 {
		delete(ms.countMap, item)
		ms.set.Remove(item)
		ms.totalSize -= existingCount
		ms.distinctSize--
	}
}

// Contains returns the number of occurrences of a given item in the MultiSet
func (ms *MultiSet) Contains(item interface{}) int {
	return ms.countMap[item]
}

// TotalSize returns total number of occurrences in the MultiSet
func (ms *MultiSet) TotalSize() int {
	return ms.totalSize
}

// DistinctSize returns the number of distinct items in the MultiSet
func (ms *MultiSet) DistinctSize() int {
	return ms.distinctSize
}

package hashset

type HashSet struct {
	hashMap map[interface{}]bool
	size    int
}

// New returns a new HashSet
func New() *HashSet {
	return &HashSet{hashMap: make(map[interface{}]bool)}
}

// Add insert an item to the set
func (hs *HashSet) Add(item interface{}) {
	hs.hashMap[item] = true
	hs.size = len(hs.hashMap)
}

// Contains returns true if item is contained in the set, else false
func (hs *HashSet) Contains(item interface{}) bool {
	return hs.hashMap[item]
}

// Remove deletes an item from the set
func (hs *HashSet) Remove(item interface{}) {
	delete(hs.hashMap, item)
	hs.size = len(hs.hashMap)
}

// Clear removes all existing items in the set
func (hs *HashSet) Clear() {
	hs.hashMap = make(map[interface{}]bool)
	hs.size = 0
}

// IsEmpty returns true if size of set is 0, else false
func (hs *HashSet) IsEmpty() bool {
	return hs.size == 0
}

// Size returns the number of items in the set
func (hs *HashSet) Size() int {
	return hs.size
}

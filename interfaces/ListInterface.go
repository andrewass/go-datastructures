package interfaces

type ListInterface interface {
	Get(index int) interface{}
	Add(interface{})
	Remove(index int)
	Swap(indexA int, indexB int)
	Size() int
}

package fenwicktree

type FenwickTree struct {
	list []int64
	size int
}

func New(size int) *FenwickTree {
	return &FenwickTree{size: size, list: make([]int64, size+1)}
}

func (fwt *FenwickTree) Add(ind int, value int64) {
	index := ind

	for index <= fwt.size {
		fwt.list[index] += value
		index += index & -index
	}
}

func (fwt *FenwickTree) GetSum(ind int) int64 {
	index := ind
	var sum int64

	for index >= 1 {
		sum += fwt.list[index]
		index -= index & -index
	}
	return sum
}

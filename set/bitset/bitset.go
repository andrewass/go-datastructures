package bitset

type BitSet struct {
	list []bool

}

func New(size int) *BitSet {
	return &BitSet{list: make([]bool, size)}
}

func (bs *BitSet) Size() int {
	return len(bs.list)
}


func(bs *BitSet) Add(isSet bool) {
	bs.list = append(bs.list, isSet)
}

func (bs *BitSet) Get(index int) bool {
	return bs.list[index]
}

func (bs *BitSet) Set(index int) {
	bs.list[index] = true
}

func (bs *BitSet) SetRange(from int, to int) {
	for i := from; i < to; i++ {
		bs.list[i] = true
	}
}

func (bs *BitSet) Clear(index int) {
	bs.list[index] = false
}

func (bs *BitSet) ClearRange(from int, to int) {
	for i := from; i < to; i++ {
		bs.list[i] = false
	}
}

func (bs *BitSet) Flip(index int) {
	bs.list[index] = !bs.list[index]
}

func (bs *BitSet) FlipRange(from int, to int) {
	for i := from; i < to; i++ {
		bs.list[i] = !bs.list[i]
	}
}

func (bs *BitSet) Cardinality() int {
	setBits := 0
	for _, bit := range bs.list {
		if bit {
			setBits++
		}
	}
	return setBits
}

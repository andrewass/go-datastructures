package bitset

import (
	"math/big"
)

type BitSet struct {
	bits big.Int
	length int
}

func New() *BitSet {
	return &BitSet{}
}

func (bs *BitSet) Size() int {
	return bs.length
}

func (bs *BitSet) Add(bit uint) {
	bs.bits.SetBit(&bs.bits, bs.length, bit)
	bs.length++
}

func (bs *BitSet) Get(index int) uint {
	return bs.bits.Bit(index)
}

func (bs *BitSet) Set(index int) {
	bs.bits.SetBit(&bs.bits, index, 1)
	if index >= bs.length {
		bs.length = index+1
	}
}

func (bs *BitSet) SetRange(from int, to int) {
	for i := from; i < to; i++ {
		bs.bits.SetBit(&bs.bits, i, 1)
	}
	if to-1 >= bs.length {
		bs.length = to
	}
}

func (bs *BitSet) Clear(index int) {
	bs.bits.SetBit(&bs.bits, index, 0)
}

func (bs *BitSet) ClearRange(from int, to int) {
	for i := from; i < to; i++ {
		bs.bits.SetBit(&bs.bits, i, 0)
	}
}

func (bs *BitSet) FlipRange(from int, to int) {
	for i := from; i < to; i++ {
		bs.Flip(i)
	}
	if to-1 >= bs.length {
		bs.length = to
	}
}

func (bs *BitSet) Flip(index int) {
	var bit = bs.bits.Bit(index)
	if bit == 0 {
		bit = 1
	} else {
		bit = 0
	}
	bs.bits.SetBit(&bs.bits, index, bit)
	if index >= bs.length {
		bs.length = index+1
	}
}

func (bs *BitSet) Cardinality() int {
	setBits := 0
	for i := 0; i< bs.length; i++ {
		if bs.bits.Bit(i) == 1 {
			setBits++
		}
	}
	return setBits
}

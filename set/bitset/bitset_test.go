package bitset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldSetExpectedBit(t *testing.T) {
	bitSet := New()

	bitSet.Set(4)
	assert.Equal(t, 1, bitSet.Get(4))
}

func TestShouldClearExpectedBit(t *testing.T) {
	bitSet := New()

	bitSet.Set(4)
	bitSet.Clear(4)
	assert.Equal(t, 0, bitSet.Get(4))
}

func TestShouldFlipExpectedBit(t *testing.T) {
	bitSet := New()

	bitSet.Flip(4)
	assert.Equal(t, 1, bitSet.Get(4))

	bitSet.Flip(4)
	assert.Equal(t, 0, bitSet.Get(4))
}

func TestShouldSetRangeOfBits(t *testing.T) {
	bitSet := New()

	bitSet.SetRange(3, 10)
	assert.Equal(t, 7, bitSet.Cardinality())
}

func TestShouldClearRangeOfBits(t *testing.T) {
	bitSet := New()

	bitSet.SetRange(0, 20)
	bitSet.ClearRange(5, 10)
	assert.Equal(t, 15, bitSet.Cardinality())
}

func TestShouldFlipRangeOfBits(t *testing.T) {
	bitSet := New()

	bitSet.FlipRange(5, 10)
	assert.Equal(t, 5, bitSet.Cardinality())

	bitSet.FlipRange(5, 10)
	assert.Equal(t, 0, bitSet.Cardinality())
}

package bitset

import "testing"

func TestShouldSetExpectedBit(t *testing.T)  {
	bitSet := New()
	bitSet.Set(4)

	if bitSet.Get(4) == 0 {
		t.Error("Bit at index 4 should be set")
	}
}

func TestShouldClearExpectedBit(t *testing.T) {
	bitSet := New()
	bitSet.Set(4)
	bitSet.Clear(4)

	if bitSet.Get(4) == 1 {
		t.Error("Bit at index 4 should be cleared")
	}
}

func TestShouldFlipExpectedBit(t *testing.T)  {
	bitSet := New()

	bitSet.Flip(4)
	if bitSet.Get(4) == 0{
		t.Error("Bit at index 4 should be set")
	}

	bitSet.Flip(4)
	if bitSet.Get(4) == 1 {
		t.Error("Bit at index 4 should be cleared")
	}
}

func TestShouldSetRangeOfBits(t *testing.T)  {
	bitSet := New()
	bitSet.SetRange(3, 10)

	setCount := bitSet.Cardinality()
	if setCount != 7 {
		t.Error("Number of set bits should be 7, got ",setCount)
	}
}

func TestShouldClearRangeOfBits(t *testing.T)  {
	bitSet := New()
	bitSet.SetRange(0, 20)
	bitSet.ClearRange(5, 10)

	setCount := bitSet.Cardinality()
	if setCount != 15 {
		t.Error("Number of set bits should be 15, got ",setCount)
	}
}

func TestShouldFlipRangeOfBits(t *testing.T)  {
	bitSet := New()

	bitSet.FlipRange(5, 10)
	setCount := bitSet.Cardinality()
	if setCount != 5 {
		t.Error("Number of set bits should be 5, got ",setCount)
	}

	bitSet.FlipRange(5, 10)
	setCount = bitSet.Cardinality()
	if setCount != 0 {
		t.Error("Number of set bits should be 0, got ",setCount)
	}
}
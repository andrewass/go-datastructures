package range_query

import "testing"

func TestShouldUpdateFenwickTreeAsExpected(t *testing.T)  {
	fwt := New(10)

	fwt.Add(5, 50)
	sum := fwt.GetSum(5)
	if sum != 50 {
		t.Error("Sum should be 50, got ",sum)
	}

	fwt.Add(8, 100)
	sum = fwt.GetSum(10)
	if sum != 150 {
		t.Error("Sum should be 150, got ",sum)
	}

	fwt.Add(8, -200)
	sum = fwt.GetSum(10)
	if sum != -50 {
		t.Error("Sum should be -50, got ",sum)
	}
}

func TestShouldReturnZeroSumFromEmptyRange(t *testing.T)  {
	fwt := New(10)

	fwt.Add(10, 50)
	sum := fwt.GetSum(9)
	if sum != 0 {
		t.Error("Sum should be 0, got ",sum)
	}
}


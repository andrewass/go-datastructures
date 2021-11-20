package fenwicktree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldUpdateFenwickTreeAsExpected(t *testing.T)  {
	fwt := New(10)

	fwt.Add(5, 50)
	assert.Equal(t, int64(50), fwt.GetSum(5))

	fwt.Add(8, 100)
	assert.Equal(t, int64(150), fwt.GetSum(10))

	fwt.Add(8, -200)
	assert.Equal(t, int64(-50), fwt.GetSum(10))
}

func TestShouldReturnZeroSumFromEmptyRange(t *testing.T)  {
	fwt := New(10)

	fwt.Add(10, 50)
	assert.Equal(t, int64(0), fwt.GetSum(9))
}


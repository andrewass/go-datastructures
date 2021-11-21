package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnPrimeFactorsOfNumber(t *testing.T)  {
	primeFactors := GetAllPrimeFactors(5645767435)

	assert.Equal(t, 5, primeFactors.Size())
	assert.Equal(t, int64(5), primeFactors.Get(0))
	assert.Equal(t, int64(7), primeFactors.Get(1))
	assert.Equal(t, int64(11), primeFactors.Get(2))
	assert.Equal(t, int64(11), primeFactors.Get(3))
	assert.Equal(t, int64(1333121), primeFactors.Get(4))
}

func TestShouldReturnDistinctPrimeFactorsOfNumber(t *testing.T)  {
	primeFactors := GetAllDistinctPrimeFactors(5645767435)

	assert.Equal(t, 4, primeFactors.Size())
	assert.Equal(t, int64(5), primeFactors.Get(0))
	assert.Equal(t, int64(7), primeFactors.Get(1))
	assert.Equal(t, int64(11), primeFactors.Get(2))
	assert.Equal(t, int64(1333121), primeFactors.Get(3))
}

func TestShouldReturnSinglePrimeWhenInputIsPrime(t *testing.T)  {
	primeFactors := GetAllPrimeFactors(1333121)

	assert.Equal(t, 1, primeFactors.Size())
	assert.Equal(t, int64(1333121), primeFactors.Get(0))
}

func TestShouldReturnEmptyListWhenNumberContainsNoPrimes(t *testing.T)  {
	primeFactors := GetAllPrimeFactors(1)

	assert.Equal(t, 0, primeFactors.Size())
}

func TestShouldReturnGreatestCommonDivisor(t *testing.T)  {
	assert.Equal(t, int64(12), GreatestCommonDivisor(36, 24))
	assert.Equal(t, int64(1), GreatestCommonDivisor(31, 11))
	assert.Equal(t, int64(56), GreatestCommonDivisor(56, 56))
}

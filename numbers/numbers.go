package numbers

import (
	"go-datastructures-algorithms/list/arraylist"
	"math"
)

// GetAllDistinctPrimeFactors :Get all distinct prime factors of given number
func GetAllDistinctPrimeFactors(number int64) *arraylist.ArrayList {
	factors := GetAllPrimeFactors(number)
	return factors.GetDistinctItems()
}

// GetAllPrimeFactors :Get all prime factors of given number, including duplicates
func GetAllPrimeFactors(number int64) *arraylist.ArrayList {
	factors := arraylist.New()
	remaining := float64(number)
	limit := Sqrt(number)

	for math.Mod(remaining, 2.00) == 0 {
		factors.Add(2)
		remaining /= 2
	}
	divisor := 3.00
	for divisor <= limit {
		if math.Mod(remaining, divisor) == 0 {
			factors.Add(int64(divisor))
			remaining /= divisor
		} else {
			divisor += 2
		}
	}
	if remaining != 1 {
		factors.Add(int64(remaining))
	}
	return factors
}

// GreatestCommonDivisor : Calculate greatest number dividing both int a and b
func GreatestCommonDivisor(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, Mod(a, b))
}

// Sqrt : Calculate the square root of given number
func Sqrt(number int64) float64 {
	return math.Sqrt(float64(number))
}

// Mod : Calculate remainder when dividing int a with b
func Mod(a, b int64) int64 {
	return int64(math.Mod(float64(a), float64(b)))
}

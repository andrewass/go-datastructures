package numbers

import (
	"go-datastructures-algorithms/list/arraylist"
	"math"
)

func GetAllDistinctPrimeFactors(number int64) *arraylist.ArrayList {
	factors := GetAllPrimeFactors(number)
	return factors.GetDistinctItems()
}

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

func Sqrt(number int64) float64 {
	return math.Sqrt(float64(number))
}

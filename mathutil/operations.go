package mathutil

import (
	"errors"
	"math"
)

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func checkOdd(x int) bool {
	return x%2 == 0
}

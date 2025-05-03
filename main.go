package main

import (
	"errors"
	"math"
)

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("sqrt of negative")
	}
	return math.Sqrt(a), nil
}

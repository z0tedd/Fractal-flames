package pkg

import (
	"math"
	"math/rand"
)

// Utility functions.
func RandomBit() int {
	return rand.Intn(2)
}

func Modulus(a, b float64) float64 {
	return math.Mod(a, b)
}

func RandRange(mn, mx float64) float64 {
	return mn + rand.Float64()*(mx-mn)
}

func RandR(mn, mx float64) float64 {
	return mn + rand.Float64()*(mx-mn)
}

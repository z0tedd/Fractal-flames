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

func RandRangeUINT8(mn, mx uint8) uint8 {
	return mn + uint8(rand.Intn(int(mx-mn)))
}

func RandInt(n int) int {
	return rand.Intn(n)
}

func RandFloat() float64 {
	return rand.Float64()
}

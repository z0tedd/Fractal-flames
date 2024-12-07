package pkg

import (
	"crypto/rand"
	"math"
	"math/big"
)

// clamp function to restrict values between 0 and 255.
func Clamp(value uint) uint8 {
	if value > 255 {
		return 255
	}

	return uint8(value)
}

// Utility functions.
func RandomBit() int {
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		panic(err)
	}

	return int(n.Int64())
}

func Modulus(a, b float64) float64 {
	return math.Mod(a, b)
}

func RandRange(mn, mx float64) float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64((mx-mn)*1e9)))
	if err != nil {
		panic(err)
	}

	return mn + float64(n.Int64())/1e9
}

func RandRangeUint8(mn, mx uint8) uint8 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(mx-mn)))
	if err != nil {
		panic(err)
	}

	// А здесь вы можете видеть, как человек борется с линтером, предупреждающим о переполнении uint8
	return mn + Clamp(uint(n.Uint64()))
}

func RandInt(n int) int {
	if n <= 0 {
		panic("RandInt: n must be greater than 0")
	}

	val, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err)
	}

	return int(val.Int64())
}

func RandFloat() float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(1e9))
	if err != nil {
		panic(err)
	}

	return float64(n.Int64()) / 1e9
}

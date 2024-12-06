package application

import (
	"flame/internal/domain"
	"math"
)

// Apply gamma correction and logarithmic normalization.
func GammaLog(fractal *domain.Flame) {
	mx := 0.0

	// Find the maximum logarithmic value
	for row := 0; row < fractal.YRes; row++ {
		for col := 0; col < fractal.XRes; col++ {
			if fractal.Pixels[row][col].Value.Counter != 0 {
				normal := math.Log(float64(fractal.Pixels[row][col].Value.Counter))
				fractal.Pixels[row][col].Value.Normal = normal

				if normal > mx {
					mx = normal
				}
			}
		}
	}

	// Normalize and apply gamma correction
	for row := 0; row < fractal.YRes; row++ {
		for col := 0; col < fractal.XRes; col++ {
			pixel := &fractal.Pixels[row][col]
			pixel.Value.Normal /= mx
			gammaFactor := math.Pow(float64(pixel.Value.Normal), 1.0/GammaCorrection)
			pixel.Color.R = uint8(float64(pixel.Color.R) * gammaFactor)
			pixel.Color.G = uint8(float64(pixel.Color.G) * gammaFactor)
			pixel.Color.B = uint8(float64(pixel.Color.B) * gammaFactor)
		}
	}
}

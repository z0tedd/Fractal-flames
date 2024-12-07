package application

import (
	"flame/internal/domain"
	"math"
)

type Canvas interface {
	Canvas() [][]domain.Pixel
}

// Apply gamma correction and logarithmic normalization.
func GammaLog(fractal Canvas) {
	mx := 0.0
	canvas := fractal.Canvas()
	XRes, YRes := len(canvas[0]), len(canvas)
	// Find the maximum logarithmic value
	for row := 0; row < YRes; row++ {
		for col := 0; col < XRes; col++ {
			if canvas[row][col].Value.Counter != 0 {
				normal := math.Log(float64(canvas[row][col].Value.Counter))
				canvas[row][col].Value.Normal = normal

				if normal > mx {
					mx = normal
				}
			}
		}
	}

	// Normalize and apply gamma correction
	for row := 0; row < YRes; row++ {
		for col := 0; col < XRes; col++ {
			pixel := &canvas[row][col]
			pixel.Value.Normal /= mx
			gammaFactor := math.Pow(float64(pixel.Value.Normal), 1.0/GammaCorrection)
			pixel.Color.R = uint8(float64(pixel.Color.R) * gammaFactor)
			pixel.Color.G = uint8(float64(pixel.Color.G) * gammaFactor)
			pixel.Color.B = uint8(float64(pixel.Color.B) * gammaFactor)
		}
	}
}

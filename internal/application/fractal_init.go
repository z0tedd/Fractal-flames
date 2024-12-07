package application

import (
	"flame/internal/domain"
)

// Initialize default flame structure values.
func NewFractal(config *domain.Config) *domain.Flame {
	var fractal domain.Flame
	fractal.Config = config

	fractal.Pixels = make([][]domain.Pixel, fractal.YRes)
	for y := 0; y < fractal.YRes; y++ {
		fractal.Pixels[y] = make([]domain.Pixel, fractal.XRes)
	}
	CoeffInit(&fractal)
	return &fractal
}

package application

import (
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
)

// Initialize default github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd structure values.
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

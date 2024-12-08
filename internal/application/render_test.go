package application_test

import (
	"flame/internal/application"
	"flame/internal/domain"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	fractal := application.NewFractal(&domain.Config{
		FractalType:   1,
		Samples:       10,
		Iterations:    100,
		Symmetry:      2,
		XMin:          -1.0,
		XMax:          1.0,
		YMin:          -1.0,
		YMax:          1.0,
		XRes:          100,
		YRes:          100,
		RanX:          2.0,
		RanY:          2.0,
		NumV:          16,
		Debug:         false,
		MaxColorRange: color.RGBA{64, 64, 64, 0},
	})

	// Execute Render function
	application.Render(fractal)

	// Verify that the function modifies some pixels
	pixelModified := false

	for row := range fractal.Pixels {
		for col := range row {
			pixel := &fractal.Pixels[row][col]
			if pixel.Value.Counter > 0 {
				pixelModified = true
				break
			}
		}
	}

	assert.True(t, pixelModified, "At least one pixel should be modified after rendering")
}

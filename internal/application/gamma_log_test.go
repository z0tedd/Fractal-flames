package application_test

import (
	"image/color"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/pkg/mocks"
)

func TestGammaLog(t *testing.T) {
	// Mock canvas setup.
	xRes, yRes := 3, 3
	mockPixels := make([][]domain.Pixel, yRes)

	for i := 0; i < yRes; i++ {
		mockPixels[i] = make([]domain.Pixel, xRes)
		for j := 0; j < xRes; j++ {
			mockPixels[i][j] = domain.Pixel{
				Value: domain.HitCounter{Counter: i*j + 1, Normal: 0}, // Simple counter logic
				Color: color.RGBA{R: 100, G: 150, B: 200},
			}
		}
	}

	mockCanvas := new(mocks.Canvas)
	mockCanvas.On("Canvas").Return(mockPixels)

	// Test configuration
	config := &domain.Config{
		GammaCorrection: 2.2,
	}

	// Call GammaLog
	application.GammaLog(mockCanvas, config)

	// Assertions
	canvas := mockCanvas.Canvas()
	// Verify normalization and gamma correction
	for row := 0; row < yRes; row++ {
		for col := 0; col < xRes; col++ {
			pixel := &canvas[row][col]

			// Check color adjustment
			gammaFactor := math.Pow(pixel.Value.Normal, 1.0/config.GammaCorrection)
			assert.Equal(t, uint8(float64(100)*gammaFactor), pixel.Color.R, "Gamma-corrected R value mismatch")
			assert.Equal(t, uint8(float64(150)*gammaFactor), pixel.Color.G, "Gamma-corrected G value mismatch")
			assert.Equal(t, uint8(float64(200)*gammaFactor), pixel.Color.B, "Gamma-corrected B value mismatch")
		}
	}

	mockCanvas.AssertExpectations(t)
}

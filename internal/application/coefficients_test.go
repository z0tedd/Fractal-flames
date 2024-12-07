package application_test

import (
	"flame/internal/application"
	"flame/internal/domain"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoeffInit(t *testing.T) {
	// Prepare test fractal.
	numV := 3
	standartColor := color.RGBA{R: 0, G: 0, B: 0}
	fractal := &domain.Flame{
		Config: &domain.Config{
			NumV:          numV,
			StandartColor: standartColor,
			MinColorRange: color.RGBA{R: 10, G: 20, B: 30},
			MaxColorRange: color.RGBA{R: 100, G: 200, B: 255},
			Debug:         false,
		},
		Coefficients: nil,
	}
	// Call CoeffInit.
	application.CoeffInit(fractal)

	// Assert the length of coefficients.
	assert.Equal(t, numV, len(fractal.Coefficients), "Coefficients length should match NumV")

	for i, coeff := range fractal.Coefficients {
		// Assert coefficient values are initialized.
		assert.NotZero(t, coeff.AC, "AC should be initialized for coefficient %d", i)
		assert.NotZero(t, coeff.BC, "BC should be initialized for coefficient %d", i)
		assert.NotZero(t, coeff.CC, "CC should be initialized for coefficient %d", i)
		assert.NotZero(t, coeff.DC, "DC should be initialized for coefficient %d", i)
		assert.NotZero(t, coeff.EC, "EC should be initialized for coefficient %d", i)
		assert.NotZero(t, coeff.FC, "FC should be initialized for coefficient %d", i)

		// Assert additional parameters are assigned within the correct range.
		assert.GreaterOrEqual(t, coeff.PA1, -2.0, "PA1 should be >= -2")
		assert.LessOrEqual(t, coeff.PA1, 2.0, "PA1 should be <= 2")
		assert.GreaterOrEqual(t, coeff.PA2, -2.0, "PA2 should be >= -2")
		assert.LessOrEqual(t, coeff.PA2, 2.0, "PA2 should be <= 2")
		assert.GreaterOrEqual(t, coeff.PA3, -2.0, "PA3 should be >= -2")
		assert.LessOrEqual(t, coeff.PA3, 2.0, "PA3 should be <= 2")
		assert.GreaterOrEqual(t, coeff.PA4, -2.0, "PA4 should be >= -2")
		assert.LessOrEqual(t, coeff.PA4, 2.0, "PA4 should be <= 2")

		// Assert color ranges are correctly assigned.
		assert.GreaterOrEqual(t, coeff.Color.R, fractal.MinColorRange.R, "Color R should be >= MinColorRange.R for coefficient %d", i)
		assert.LessOrEqual(t, coeff.Color.R, fractal.MaxColorRange.R, "Color R should be <= MaxColorRange.R for coefficient %d", i)
		assert.GreaterOrEqual(t, coeff.Color.G, fractal.MinColorRange.G, "Color G should be >= MinColorRange.G for coefficient %d", i)
		assert.LessOrEqual(t, coeff.Color.G, fractal.MaxColorRange.G, "Color G should be <= MaxColorRange.G for coefficient %d", i)
		assert.GreaterOrEqual(t, coeff.Color.B, fractal.MinColorRange.B, "Color B should be >= MinColorRange.B for coefficient %d", i)
		assert.LessOrEqual(t, coeff.Color.B, fractal.MaxColorRange.B, "Color B should be <= MaxColorRange.B for coefficient %d", i)
	}
}

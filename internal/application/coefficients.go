package application

import (
	"log/slog"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/pkg"
)

// Initialize coefficients for the fractal.
func CoeffInit(fractal *domain.Flame) {
	fractal.Coefficients = make([]domain.Coeff, fractal.NumV)

	for i := 0; i < fractal.NumV; i++ {
		if pkg.RandomBit() == 1 {
			ContractiveMapping(&fractal.Coefficients[i])
		} else {
			fractal.Coefficients[i] = domain.Coeff{
				AC: pkg.RandRange(-1.5, 1.5),
				BC: pkg.RandRange(-1.5, 1.5),
				CC: pkg.RandRange(-1.5, 1.5),
				DC: pkg.RandRange(-1.5, 1.5),
				EC: pkg.RandRange(-1.5, 1.5),
				FC: pkg.RandRange(-1.5, 1.5),
			}
		}

		// Assign additional parameters
		fractal.Coefficients[i].PA1 = pkg.RandRange(-2, 2)
		fractal.Coefficients[i].PA2 = pkg.RandRange(-2, 2)
		fractal.Coefficients[i].PA3 = pkg.RandRange(-2, 2)
		fractal.Coefficients[i].PA4 = pkg.RandRange(-2, 2)

		// Assign colors.
		fractal.Coefficients[i].Color = fractal.StandartColor

		if fractal.StandartColor.R == 0 {
			fractal.Coefficients[i].Color.R = pkg.RandRangeUint8(fractal.MinColorRange.R, fractal.MaxColorRange.R)
		}

		if fractal.StandartColor.G == 0 {
			fractal.Coefficients[i].Color.G = pkg.RandRangeUint8(fractal.MinColorRange.G, fractal.MaxColorRange.G)
		}

		if fractal.StandartColor.B == 0 {
			fractal.Coefficients[i].Color.B = pkg.RandRangeUint8(fractal.MinColorRange.B, fractal.MaxColorRange.B)
		}
	}

	pkg.Logger.Debug("Coefficients")
	// Print coefficients for debugging
	for _, coeff := range fractal.Coefficients {
		pkg.Logger.Debug("ac,bc,cc,dc,ec,fc", slog.Any("values: ", []float64{coeff.AC, coeff.BC, coeff.CC, coeff.DC, coeff.EC, coeff.FC}))
	}

	pkg.Logger.Debug("Colors")

	for _, coeff := range fractal.Coefficients {
		pkg.Logger.Debug("coeff :", slog.Any("color:", coeff.Color))
	}
}

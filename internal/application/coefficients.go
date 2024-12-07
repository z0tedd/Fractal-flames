package application

import (
	"flame/internal/domain"
	"flame/pkg"
	"fmt"
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

	if fractal.Debug {
		fmt.Println("Coefficients")
		// Print coefficients for debugging
		for _, coeff := range fractal.Coefficients {
			fmt.Printf("%f %f %f %f %f %f\n", coeff.AC, coeff.BC, coeff.CC, coeff.DC, coeff.EC, coeff.FC)
		}

		fmt.Println("Colors")

		for _, coeff := range fractal.Coefficients {
			fmt.Printf("%v\n", coeff.Color)
		}
	}
}

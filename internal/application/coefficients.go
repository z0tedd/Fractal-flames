package application

import (
	"flame/internal/domain"
	"flame/pkg"
	"fmt"
)

// Initialize coefficients for the fractal.
func CoeffInit(fractal *domain.Flame) {
	fractal.Coefficients = make([]domain.Coeff, fractal.N)

	for i := 0; i < fractal.N; i++ {
		if pkg.RandomBit() == 1 {
			contractiveMapping(&fractal.Coefficients[i])
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
		fractal.Coefficients[i].Color = fractal.Color

		if fractal.Color.R == 0 {
			fractal.Coefficients[i].Color.R = pkg.RandRangeUINT8(128, 255)
		}

		if fractal.Color.G == 0 {
			fractal.Coefficients[i].Color.G = pkg.RandRangeUINT8(0, 128)
		}

		if fractal.Color.B == 0 {
			fractal.Coefficients[i].Color.B = pkg.RandRangeUINT8(0, 128)
		}
	}

	// Print coefficients for debugging
	for _, coeff := range fractal.Coefficients {
		fmt.Printf("%f %f %f %f %f %f\n", coeff.AC, coeff.BC, coeff.CC, coeff.DC, coeff.EC, coeff.FC)
	}

	fmt.Println("Colors")

	for _, coeff := range fractal.Coefficients {
		fmt.Printf("%v\n", coeff.Color)
	}
}

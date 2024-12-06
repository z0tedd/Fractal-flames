package application

import (
	"flame/internal/domain"
	"flame/pkg"
	"fmt"
	"log"
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

		if fractal.Color.R == 255 {
			fractal.Coefficients[i].Color.R = uint8(pkg.RandRange(64, 255))
		}

		if fractal.Color.G == 255 {
			fractal.Coefficients[i].Color.G = uint8(pkg.RandRange(64, 255))
		}

		if fractal.Color.B == 255 {
			fractal.Coefficients[i].Color.B = uint8(pkg.RandRange(64, 255))
		}
	}

	// Load palette file if provided
	if fractal.PaletteFile != nil {
		i := 0

		var r, g, b uint8
		for _, err := fmt.Fscanf(fractal.PaletteFile, "%d %d %d\n", &r, &g, &b); err == nil && i < fractal.N; {
			fractal.Coefficients[i].Color.R = r
			fractal.Coefficients[i].Color.G = g
			fractal.Coefficients[i].Color.B = b
			fmt.Printf("Setting index %d to %d,%d,%d\n", i, r, g, b)

			i++
		}

		err := fractal.PaletteFile.Close()
		if err != nil {
			log.Println("Failed to close palette file:", err)
		}
	}

	// Load coefficients file if provided
	if fractal.CoeffFile != nil {
		i := 0
		var ac, bc, cc, dc, ec, fc float64
		for _, err := fmt.Fscanf(fractal.CoeffFile, "%f %f %f %f %f %f\n", &ac, &bc, &cc, &dc, &ec, &fc); err == nil && i < fractal.N; {
			fractal.Coefficients[i] = domain.Coeff{AC: ac, BC: bc, CC: cc, DC: dc, EC: ec, FC: fc}
			fmt.Printf("Setting index coeffs at index %d\n", i)
			i++
		}

		err := fractal.CoeffFile.Close()
		if err != nil {
			log.Println("Failed to close coefficients file:", err)
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

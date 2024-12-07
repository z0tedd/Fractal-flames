package application

import (
	"flame/internal/domain"
	"flame/pkg"
)

// Contractive mapping function.
func contractiveMapping(coeff *domain.Coeff) {
	var a, b, d, e float64

	for {
		// Generate valid values for a, d
		for {
			a = pkg.RandFloat()
			d = pkg.RandRange(a*a, 1)

			if pkg.RandomBit() == 1 {
				d = -d
			}

			if a*a+d*d <= 1 {
				break
			}
		}

		// Generate valid values for b, e
		for {
			b = pkg.RandFloat()
			e = pkg.RandRange(b*b, 1)

			if pkg.RandomBit() == 1 {
				e = -e
			}

			if b*b+e*e <= 1 {
				break
			}
		}

		// Check final condition
		if a*a+b*b+d*d+e*e <= 1+(a*e-d*b)*(a*e-d*b) {
			break
		}
	}

	coeff.AC = a
	coeff.BC = b
	coeff.CC = pkg.RandRange(-2, 2)
	coeff.DC = d
	coeff.EC = e
	coeff.FC = pkg.RandRange(-2, 2)
}

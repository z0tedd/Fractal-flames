package application

import (
	"math"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
)

func applyTransformations(k int, x, y float64, Coefficient domain.Coeff) (float64, float64) { //nolint
	var newX, newY float64

	switch k {
	case 0: // Linear
		newX, newY = x, y
	case 1: // Sinusoidal
		newX, newY = math.Sin(x), math.Sin(y)
	case 2: // Spherical
		r := 1.0 / (x*x + y*y)
		newX, newY = r*x, r*y
	case 3: // Swirl
		r := x*x + y*y
		newX, newY = x*math.Sin(r)-y*math.Cos(r), x*math.Cos(r)+y*math.Sin(r)
	case 4: // Horseshoe
		r := 1.0 / math.Sqrt(x*x+y*y)
		newX, newY = r*(x-y)*(x+y), r*2.0*x*y
	case 5: // Polar
		newX = math.Atan2(y, x) / math.Pi
		newY = math.Sqrt(x*x+y*y) - 1.0
	case 6: // Handkerchief
		r := math.Sqrt(x*x + y*y)
		theta := math.Atan2(y, x)
		newX = r * math.Sin(theta+r)
		newY = r * math.Cos(theta-r)

	case 7: // Heart
		r := math.Sqrt(x*x + y*y)
		theta := math.Atan2(y, x)
		newX = r * math.Sin(theta*r)
		newY = -r * math.Cos(theta*r)

	case 8: // Disk
		r := math.Sqrt(x*x+y*y) * math.Pi
		theta := math.Atan2(y, x) / math.Pi
		newX = theta * math.Sin(r)
		newY = theta * math.Cos(r)

	case 9: // Spiral
		r := math.Sqrt(x*x + y*y)
		theta := math.Atan2(y, x)
		newX = (1.0 / r) * (math.Cos(theta) + math.Sin(r))
		newY = (1.0 / r) * (math.Sin(theta) - math.Cos(r))

	case 10: // Hyperbolic
		r := math.Sqrt(x*x + y*y)
		theta := math.Atan2(y, x)
		newX = math.Sin(theta) / r
		newY = r * math.Cos(theta)

	case 11: // Diamond
		r := math.Sqrt(x*x + y*y)
		theta := math.Atan2(y, x)
		newX = math.Sin(theta) * math.Cos(r)
		newY = math.Cos(theta) * math.Sin(r)

	case 12: // Ex
		r := math.Sqrt(x*x + y*y)
		theta := math.Atan2(y, x)
		P0 := math.Sin(theta + r)
		P0 = P0 * P0 * P0
		P1 := math.Cos(theta - r)
		P1 = P1 * P1 * P1
		newX = r * (P0 + P1)
		newY = r * (P0 - P1)

	case 13: // Waves
		newX = x + Coefficient.PA1*math.Sin(y/(Coefficient.PA2*Coefficient.PA2))
		newY = y + Coefficient.PA3*math.Sin(x/(Coefficient.PA4*Coefficient.PA4))

	case 14: // Fisheye
		r := 2.0 / (1. + math.Sqrt(x*x+y*y))
		newX = r * y
		newY = r * x

	case 15: // Popcorn
		newX = x + Coefficient.CC*math.Sin(math.Tan(3.0*y))
		newY = y + Coefficient.FC*math.Sin(math.Tan(3.0*x))

	case 16: // math.exponential
		newX = math.Exp(x-1.0) * math.Cos(math.Pi*y)
		newY = math.Exp(x-1.0) * math.Sin(math.Pi*y)
	case 17: // math.power
		r := math.Sqrt(x*x + y*y)
		theta := math.Atan2(y, x)
		newX = math.Pow(r, math.Sin(theta)) * math.Cos(theta)
		newY = math.Pow(r, math.Sin(theta)) * math.Sin(theta)

	case 18: // Comath.Sine
		newX = math.Cos(math.Pi*x) * math.Cosh(y)
		newY = -math.Sin(math.Pi*x) * math.Sinh(y)
	case 19: // Eyefish
		r := 2.0 / (1. + math.Sqrt(x*x+y*y))
		newX = r * x
		newY = r * y

	case 20: // Bubble
		r := 4 + x*x + y*y
		newX = (4.0 * x) / r
		newY = (4.0 * y) / r

	case 21: // Cylinder
		newX = math.Sin(x)
		newY = y

	case 22: // math.Tangent
		newX = math.Sin(x) / math.Cos(y)
		newY = math.Tan(y)

	case 23: // Cross
		r := math.Sqrt(1.0 / ((x*x - y*y) * (x*x - y*y)))
		newX = x * r
		newY = y * r

	case 24: // Collatz
		newX = .25 * (1.0 + 4.0*x - (1.0+2.0*x)*math.Cos(math.Pi*x))
		newY = .25 * (1.0 + 4.0*y - (1.0+2.0*y)*math.Cos(math.Pi*y))
	default:
		newX, newY = 0, 0
	}

	return newX, newY
}

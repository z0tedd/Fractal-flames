package domain

import "image/color"

type Coeff struct {
	AC, BC, CC, DC, EC, FC float64 // Transformation coefficients
	PA1, PA2, PA3, PA4     float64
	Color                  color.RGBA
}

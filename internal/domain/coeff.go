package domain

import "image/color"

type Coeff struct {
	AC, BC, CC, DC, EC, FC float64 // Transformation coefficients
	PA1, PA2, PA3, PA4     float64
	// R, G, B                uint8 // RGB color values
	Color color.RGBA
}

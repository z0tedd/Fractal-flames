package domain

import (
	"image/color"
	"sync"
)

type Coeff struct {
	AC, BC, CC, DC, EC, FC float64 // Transformation coefficients
	PA1, PA2, PA3, PA4     float64
	Color                  color.RGBA
}

type HitCounter struct {
	Counter uint32  // Number of times a pixel has been hit
	Normal  float64 // Normalized pixel value
}

type Pixel struct {
	Value HitCounter
	Color color.RGBA // RGB color for a pixel
	Lock  sync.Mutex
}

type Flame struct {
	*Config
	Coefficients []Coeff
	Pixels       [][]Pixel
}

func (f Flame) Canvas() [][]Pixel {
	return f.Pixels
}

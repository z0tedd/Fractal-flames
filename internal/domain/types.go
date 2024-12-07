package domain

import (
	"image/color"
	"sync"
)

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
	File string

	Coefficients []Coeff
	Locks        []sync.Mutex
	Choice       []int
	Pixels       [][]Pixel

	XMin, YMin      float64
	XMax, YMax      float64
	RanX, RanY      float64
	GammaCorrection float64
	Iterations      int64
	Seed            int64

	XRes, YRes int

	N          int // Number of equations
	Samples    int
	NumThreads int
	Symmetry   int
	Count      int
	Invert     bool
	Color      color.RGBA
}

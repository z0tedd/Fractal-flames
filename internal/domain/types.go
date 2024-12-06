package domain

import (
	"image/color"
	"os"
	"sync"
)

type HitCounter struct {
	Counter uint32  // Number of times a pixel has been hit
	Normal  float64 // Normalized pixel value
}

type Pixel struct {
	Value HitCounter
	Color color.RGBA // RGB color for a pixel
}

type Flame struct {
	File string

	CoeffFile   *os.File
	PaletteFile *os.File

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

	// R, G, B       uint8
	N             int // Number of equations
	SuperSampling int
	Samples       int
	NumThreads    int
	Symmetry      int
	Count         int
	Invert        bool
	Color         color.RGBA
}

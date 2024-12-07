package domain

import "image/color"

type Config struct {
	OutputPath      string
	XRes, YRes      int
	XMin, YMin      float64
	XMax, YMax      float64
	RanX, RanY      float64
	GammaCorrection float64
	Iterations      int
	NumThreads      int
	NumV            int
	Samples         int
	Symmetry        int
	Count           int
	FractalType     int
	StandartColor   color.RGBA
	MaxColorRange   color.RGBA
	MinColorRange   color.RGBA

	Debug  bool
	Invert bool
}

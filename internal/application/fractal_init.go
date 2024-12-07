package application

import (
	"flame/internal/domain"
	"image/color"
	"sync"
)

// Initialize default flame structure values.
func FractalInit(fractal *domain.Flame) {
	fractal.XRes = XRes
	fractal.YRes = YRes
	fractal.XMin = min(XMax, XMin)
	fractal.YMin = min(YMin, YMax)
	fractal.XMax = max(XMax, XMax)
	fractal.YMax = max(YMax, YMin)
	fractal.RanX = fractal.XMax - fractal.XMin
	fractal.RanY = fractal.YMax - fractal.YMin
	fractal.N = NumV
	fractal.Seed = int64(Seed)
	fractal.Samples = Samples
	fractal.Symmetry = Symmetry
	fractal.Invert = Invert
	fractal.File = OutputPath
	fractal.Iterations = int64(Iterations)
	fractal.NumThreads = ThreadGroupSize
	fractal.Color = color.RGBA{0, 0, 0, 255}

	fractal.Count = 1
	fractal.Choice = Choice
	fractal.Choice[0] = 0

	if fractal.XRes <= 0 {
		fractal.XRes = XRes
	}

	if fractal.YRes <= 0 {
		fractal.YRes = YRes
	}

	fractal.Pixels = make([][]domain.Pixel, fractal.YRes)
	fractal.Locks = make([]sync.Mutex, fractal.YRes)

	for y := 0; y < fractal.YRes; y++ {
		fractal.Pixels[y] = make([]domain.Pixel, fractal.XRes)
		// Mutex for row-level locking
		fractal.Locks[y] = sync.Mutex{}
	}
}

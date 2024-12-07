package application_test

import (
	"flame/internal/application"
	"flame/internal/domain"
	"image/color"
	"runtime"
	"testing"
)

func BenchmarkDefaultRender(b *testing.B) {
	// Setup a dummy flame object
	fractal := application.NewFractal(&domain.Config{
		FractalType:   10, // Testing with Sinusoidal transformation
		Samples:       2000,
		Iterations:    10000,
		Symmetry:      2,
		XMin:          -1.0,
		XMax:          1.0,
		YMin:          -1.0,
		YMax:          1.0,
		XRes:          1920,
		YRes:          1080,
		RanX:          2.0,
		RanY:          2.0,
		NumV:          16,
		Debug:         false,
		MaxColorRange: color.RGBA{64, 64, 64, 0},
	})
	// Run benchmark
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		application.Render(fractal)
	}
}

func BenchmarkMultithreadRender(b *testing.B) {
	// Setup a dummy flame object
	fractal := application.NewFractal(&domain.Config{
		FractalType:   10, // Testing with Sinusoidal transformation
		Samples:       2000,
		Iterations:    10000,
		Symmetry:      2,
		XMin:          -1.0,
		XMax:          1.0,
		YMin:          -1.0,
		YMax:          1.0,
		XRes:          1920,
		YRes:          1080,
		RanX:          2.0,
		RanY:          2.0,
		NumV:          16,
		Debug:         false,
		MaxColorRange: color.RGBA{64, 64, 64, 0},
		NumThreads:    runtime.NumCPU(),
	})
	// Run benchmark
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		application.RenderMultithreading(fractal)
	}

	b.Log("Number of threads", runtime.NumCPU())
}

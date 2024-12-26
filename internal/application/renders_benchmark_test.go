package application_test

import (
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
	"image/color"
	"runtime"
	"testing"
)

func BenchmarkDefaultRender(b *testing.B) {
	fractal := application.NewFractal(&domain.Config{
		FractalType:   10,
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
	fractal := application.NewFractal(&domain.Config{
		FractalType:   10,
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

	b.Log("Number of threads", fractal.NumThreads)
}

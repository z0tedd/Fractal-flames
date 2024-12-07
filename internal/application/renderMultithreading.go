package application

import (
	"flame/internal/domain"
	"flame/pkg"
	"fmt"
	"math"
	"sync"
)

func RenderMultithreading(fractal *domain.Flame) {
	var wg sync.WaitGroup

	// Channel to distribute samples to worker goroutines
	sampleChan := make(chan int, fractal.Samples)

	// Initialize worker pool
	numWorkers := fractal.NumThreads // Use the number of CPU cores
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := range sampleChan {
				renderSample(fractal)
				fmt.Println(i, "_____", w)
			}
		}()
	}

	// Send samples to workers
	for num := 0; num < fractal.Samples; num++ {
		sampleChan <- num
	}

	close(sampleChan) // Close the channel to signal workers

	// Wait for all workers to finish
	wg.Wait()
}

func renderSample(fractal *domain.Flame) {
	newX := pkg.RandR(fractal.XMin, fractal.XMax)
	newY := pkg.RandR(fractal.YMin, fractal.YMax)

	for step := -20; step < int(fractal.Iterations); step++ {
		i := pkg.RandInt(fractal.N)

		c, f, b, e := fractal.Coefficients[i].CC, fractal.Coefficients[i].FC, fractal.Coefficients[i].BC, fractal.Coefficients[i].EC
		ac, dc := fractal.Coefficients[i].AC, fractal.Coefficients[i].DC

		x := ac*newX + b*newY + c
		y := dc*newX + e*newY + f

		newX, newY = applyTransformations(FractalType, x, y, fractal.Coefficients[i])

		if step <= 0 {
			continue
		}

		for s := 0; s < fractal.Symmetry; s++ {
			theta := float64(s) * (2 * math.Pi / float64(fractal.Symmetry))
			xRot := newX*math.Cos(theta) - newY*math.Sin(theta)
			yRot := newX*math.Sin(theta) + newY*math.Cos(theta)

			if xRot >= fractal.XMin && xRot <= fractal.XMax && yRot >= fractal.YMin && yRot <= fractal.YMax {
				x1 := fractal.XRes - int((fractal.XMax-xRot)/fractal.RanX*float64(fractal.XRes))
				y1 := fractal.YRes - int((fractal.YMax-yRot)/fractal.RanY*float64(fractal.YRes))

				if x1 >= 0 && x1 < fractal.XRes && y1 >= 0 && y1 < fractal.YRes {
					updatePixel(&fractal.Pixels[y1][x1], &fractal.Coefficients[i])
				}
			}
		}
	}
}

func updatePixel(point *domain.Pixel, coeff *domain.Coeff) {
	// Lock the pixel to ensure thread-safe updates
	point.Lock.Lock()
	defer point.Lock.Unlock()

	if point.Value.Counter == 0 {
		point.Color = coeff.Color
	} else {
		point.Color.R = (point.Color.R + coeff.Color.R) / 2
		point.Color.G = (point.Color.G + coeff.Color.G) / 2
		point.Color.B = (point.Color.B + coeff.Color.B) / 2
	}

	point.Value.Counter++
}

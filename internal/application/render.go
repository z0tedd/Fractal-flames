package application

import (
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/pkg"
	"math"
)

func Render(fractal *domain.Flame) {
	k := fractal.FractalType

	for num := 0; num < fractal.Samples; num++ {
		newX := pkg.RandRange(fractal.XMin, fractal.XMax)
		newY := pkg.RandRange(fractal.YMin, fractal.YMax)

		for step := -20; step < fractal.Iterations; step++ {
			i := pkg.RandInt(fractal.NumV)

			c, f, b, e := fractal.Coefficients[i].CC, fractal.Coefficients[i].FC, fractal.Coefficients[i].BC, fractal.Coefficients[i].EC
			ac, dc := fractal.Coefficients[i].AC, fractal.Coefficients[i].DC

			x := ac*newX + b*newY + c
			y := dc*newX + e*newY + f

			newX, newY = applyTransformations(k, x, y, fractal.Coefficients[i])

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
						point := &fractal.Pixels[y1][x1]

						if point.Value.Counter == 0 {
							point.Color = fractal.Coefficients[i].Color
						} else {
							point.Color.R = (point.Color.R + fractal.Coefficients[i].Color.R) / 2
							point.Color.G = (point.Color.G + fractal.Coefficients[i].Color.G) / 2
							point.Color.B = (point.Color.B + fractal.Coefficients[i].Color.B) / 2
						}

						point.Value.Counter++
					}
				}
			}
		}
	}
}

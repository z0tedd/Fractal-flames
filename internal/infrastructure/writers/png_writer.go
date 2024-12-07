package writers

import (
	"flame/internal/application"
	"flame/internal/domain"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type PngWriter struct {
	config *domain.Config
	canvas [][]domain.Pixel
}

func NewPngWriter(canvas application.Canvas, config *domain.Config) *PngWriter {
	return &PngWriter{canvas: canvas.Canvas(), config: config}
}

func (w PngWriter) Write() error {
	// Create the output file
	outputFile, err := os.Create(w.config.OutputPath)
	if err != nil {
		return fmt.Errorf("writing image: %w", err)
	}
	defer outputFile.Close()

	XRes, YRes := len(w.canvas[0]), len(w.canvas)
	// Create an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, XRes, YRes))

	// Populate the image pixels
	for row := 0; row < YRes; row++ {
		for col := 0; col < XRes; col++ {
			pixel := &w.canvas[row][col]
			r := pixel.Color.R
			g := pixel.Color.G
			b := pixel.Color.B

			if w.config.Invert {
				r = ^r
				g = ^g
				b = ^b
			}

			img.Set(col, row, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	// Encode the image to BMP format and save it
	err = png.Encode(outputFile, img)
	if err != nil {
		return fmt.Errorf("writing image: %w", err)
	}
	return nil
}

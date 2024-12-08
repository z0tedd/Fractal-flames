package writers

import (
	"flame/internal/application"
	"flame/internal/domain"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

type DefaultWriter struct {
	config  *domain.Config
	canvas  [][]domain.Pixel
	quality int
}

func NewDefaultWriter(canvas application.Canvas, config *domain.Config, quality int) *DefaultWriter {
	return &DefaultWriter{canvas: canvas.Canvas(), config: config, quality: quality}
}

func (w DefaultWriter) Write() error {
	outputFile, err := os.Create(w.config.OutputPath)
	if err != nil {
		return fmt.Errorf("writing image: %w", err)
	}
	defer outputFile.Close()

	XRes, YRes := len(w.canvas[0]), len(w.canvas)
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

	switch {
	case strings.Contains(w.config.OutputPath, "png"):
		err = png.Encode(outputFile, img)
	case strings.Contains(w.config.OutputPath, "jpeg"):
		err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: w.quality})
	case strings.Contains(w.config.OutputPath, "bmp"):
		err = bmp.Encode(outputFile, img)
	}

	if err != nil {
		return fmt.Errorf("writing image: %w", err)
	}

	return nil
}

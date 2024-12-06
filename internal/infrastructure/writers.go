package infrastructure

import (
	"flame/internal/domain"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

func WriteToPng(fractal *domain.Flame, fileName string) {
	// Create a new RGBA image with the size of the fractal
	img := image.NewRGBA(image.Rect(0, 0, fractal.XRes, fractal.YRes))

	// Populate the image with pixel data from the fractal
	for y := 0; y < fractal.YRes; y++ {
		for x := 0; x < fractal.XRes; x++ {
			pixel := fractal.Pixels[y][x]
			col := color.RGBA{R: pixel.Color.R, G: pixel.Color.G, B: pixel.Color.B, A: 255} // Use full alpha
			img.Set(x, y, col)
		}
	}

	// Create the output file
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create PNG file: %v", err)
	}
	defer outputFile.Close()

	// Encode the image to PNG format and save it to the file
	if err := png.Encode(outputFile, img); err != nil {
		log.Fatalf("Failed to encode PNG image: %v", err)
	}

	log.Printf("PNG file successfully Written: %s\n", fileName)
}

func WriteToJPEG(fractal *domain.Flame, quality int) {
	// Create the output file
	outputFile, err := os.Create("/tmp/fractal.jpeg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open outgoing image: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Create an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, fractal.XRes, fractal.YRes))

	// Populate the image pixels
	for row := 0; row < fractal.YRes; row++ {
		for col := 0; col < fractal.XRes; col++ {
			pixel := fractal.Pixels[row][col]
			r := pixel.Color.R
			g := pixel.Color.G
			b := pixel.Color.B

			if fractal.Invert {
				r = ^r
				g = ^g
				b = ^b
			}

			img.Set(col, row, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	// Encode the image to JPEG format and save it
	err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: quality})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not Write image: %v\n", err)
		os.Exit(1)
	}
}

// WriteToBMP generates a BMP image based on the domain.Flame struct
func WriteToBMP(fractal *domain.Flame) {
	// Create the output file
	outputFile, err := os.Create("/tmp/fractal.bmp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open outgoing image: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Create an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, fractal.XRes, fractal.YRes))

	// Populate the image pixels
	for row := 0; row < fractal.YRes; row++ {
		for col := 0; col < fractal.XRes; col++ {
			pixel := fractal.Pixels[row][col]
			r := pixel.Color.R
			g := pixel.Color.G
			b := pixel.Color.B

			if fractal.Invert {
				r = ^r
				g = ^g
				b = ^b
			}

			img.Set(col, row, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	// Encode the image to BMP format and save it
	err = bmp.Encode(outputFile, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not Write image: %v\n", err)
		os.Exit(1)
	}
}

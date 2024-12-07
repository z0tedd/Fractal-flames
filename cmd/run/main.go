package main

import (
	"flame/internal/application"
	"flame/internal/infrastructure/writers"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type ImageWriter interface {
	Write() error
}

func main() {
	// Seed the randomizer
	rand.Seed(time.Now().UnixNano())
	// Parse arguments from the command line
	config := application.ParseArgs()
	// Initialize our Flame Fractal
	fractal := application.NewFractal(config)

	// Correct for threads
	if fractal.NumThreads <= 1 {
		application.Render(fractal)
	} else {
		application.RenderMultithreading(fractal)
	}

	// Gamma and log correct
	fmt.Println("Finalizing and writing out...")
	application.GammaLog(fractal)
	// Write out the file
	var w ImageWriter
	switch {
	case strings.Contains(config.OutputPath, "png"):
		w = writers.NewPngWriter(fractal, config)
	case strings.Contains(config.OutputPath, "jpeg"):
		w = writers.NewJpegWriter(fractal, config, 90)
	case strings.Contains(config.OutputPath, "bmp"):
		w = writers.NewBmpWriter(fractal, config)
	}

	w.Write()
	fmt.Println("Done!")
}

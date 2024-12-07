package main

import (
	"flame/internal/application"
	"flame/internal/domain"
	"flame/internal/infrastructure"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var fractal domain.Flame
	// Parse arguments from the command line
	fmt.Println("Parsing user arguments...")
	application.ParseArgs()
	fmt.Println("Done!")

	// Initialize our Flame Fractal
	fmt.Println("Initializing...")
	application.FractalInit(&fractal)
	fmt.Println("Initialized!")

	// Seed the randomizer
	rand.Seed(time.Now().UnixNano())

	fmt.Println(fractal.Seed)

	// Initialize the random coefficients
	fmt.Println("Initializing Coefficients and Colors...")
	application.CoeffInit(&fractal)
	fmt.Println("Done!")

	// Correct for threads
	if fractal.NumThreads <= 1 {
		application.Render(&fractal)
	} else {
		application.RenderMultithreading(&fractal)
	}

	// Gamma and log correct
	fmt.Println("Finalizing and writing out...")
	application.GammaLog(&fractal)

	// Write out the file
	infrastructure.WriteToPng(&fractal, application.OutputPath)
	infrastructure.WriteToBMP(&fractal)
	infrastructure.WriteToJPEG(&fractal, 90)
	// Clean up
	fmt.Println("Cleaning up...")
	// freeResources(&fractal)
	fmt.Println("Done!")
}

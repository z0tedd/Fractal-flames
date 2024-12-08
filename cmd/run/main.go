package main

import (
	"flame/internal/application"
	"flame/internal/infrastructure/writers"
	"fmt"
	"os"
)

func main() {
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
	application.GammaLog(fractal, fractal.Config)
	// Write out the file
	w := writers.NewDefaultWriter(fractal, config, 100)

	err := w.Write()
	if err != nil {
		fmt.Printf("Program doesn't write file, error: %s\n", err.Error())
		os.Exit(2)
	}

	fmt.Println("Done!")
}

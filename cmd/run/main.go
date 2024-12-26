package main

import (
	"fmt"
	"os"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/infrastructure/writers"
)

func main() {
	// Parse arguments from the command line
	config := application.ParseArgs()
	// Initialize our Flame Fractal
	fractal := application.NewFractal(config)

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

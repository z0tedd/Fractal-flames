package main

import (
	"log/slog"
	"os"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/infrastructure/writers"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/pkg"
)

func main() {
	// Parse arguments from the command line
	config := application.ParseArgs()
	// Initialize our Flame Fractal
	fractal := application.NewFractal(config)

	// Gamma and log correct
	pkg.Logger.Info("Finalizing and writing out...")
	application.GammaLog(fractal, fractal.Config)
	// Write out the file
	w := writers.NewDefaultWriter(fractal, config, 100)

	err := w.Write()
	if err != nil {
		pkg.Logger.Info("Program doesn't write file:", slog.String("error: ", err.Error()))
		os.Exit(2)
	}

	pkg.Logger.Info("Done!")
}

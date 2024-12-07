package application

import (
	"flag"
	"flame/internal/domain"
	"flame/pkg"
	"fmt"
	"image/color"
	"runtime"
)

func ParseArgs() *domain.Config {
	var (
		config        domain.Config
		R, G, B       uint
		mnR, mnG, mnB uint
		mxR, mxG, mxB uint
	)
	// Define flags
	flag.StringVar(&config.OutputPath, "f", "/tmp/fractal.png", "File to write")
	flag.IntVar(&config.Symmetry, "S", 1, "Enable rotational symmetry axis")
	flag.BoolVar(&config.Invert, "I", false, "Invert colors in final image")
	flag.BoolVar(&config.Debug, "D", false, "Allow debug info")
	flag.IntVar(&config.XRes, "x", 1920, "Image x resolution")
	flag.IntVar(&config.YRes, "y", 1080, "Image y resolution")
	flag.Float64Var(&config.XMin, "m", -1.777, "Graph x minimum")
	flag.Float64Var(&config.XMax, "M", 1.777, "Graph x maximum")
	flag.Float64Var(&config.YMin, "l", -1, "Graph y minimum")
	flag.Float64Var(&config.YMax, "L", 1, "Graph y maximum")
	flag.IntVar(&config.NumV, "n", 16, "Number of random vectors to use")
	flag.IntVar(&config.Samples, "s", 20000, "Number of image samples")
	flag.IntVar(&config.Iterations, "i", 10000, "Number of iterations per sample")
	flag.Float64Var(&config.GammaCorrection, "Gc", 2.2, "Correctional gamma factor")
	flag.IntVar(&config.NumThreads, "T", runtime.NumCPU(), "Number of threads to run")
	flag.IntVar(&config.FractalType, "t", 1, "Fractal type, use equation by number")
	flag.UintVar(&R, "SR", 0, "Standart R color")
	flag.UintVar(&G, "SG", 0, "Standart G color")
	flag.UintVar(&B, "SB", 0, "Standart B color")
	flag.UintVar(&mnR, "r", 64, "min R color")
	flag.UintVar(&mnG, "g", 0, "min G color")
	flag.UintVar(&mnB, "b", 0, "min B color")
	flag.UintVar(&mxR, "R", 255, "max R color")
	flag.UintVar(&mxG, "G", 255, "max G color")
	flag.UintVar(&mxB, "B", 255, "max B color")

	// Custom usage message
	flag.Usage = func() {
		fmt.Println("Usage: fractal [options]")
		fmt.Println("\nOptions:")
		fmt.Println("  -h\tPrint this help message")
		flag.PrintDefaults()
		fmt.Println("\nTransformation values for -t:")
		fmt.Println("\t0: Linear, 1: Sinusoidal, 2: Spherical, 3: Swirl, 4: Horseshoe, ...")
	}
	// Parse command-line arguments
	flag.Parse()

	config.StandartColor = color.RGBA{pkg.Clamp(R), pkg.Clamp(G), pkg.Clamp(B), 255}
	config.MaxColorRange = color.RGBA{pkg.Clamp(mxR), pkg.Clamp(mxG), pkg.Clamp(mxB), 255}
	config.MinColorRange = color.RGBA{pkg.Clamp(mnR), pkg.Clamp(mnG), pkg.Clamp(mnB), 255}
	config.RanX, config.RanY = config.XMax-config.XMin, config.YMax-config.YMin

	return &config
}

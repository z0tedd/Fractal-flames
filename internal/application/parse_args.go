package application

import (
	"flag"
	"flame/internal/domain"
	"fmt"
	"image/color"
	"runtime"
)

var (
	XRes            = 1920
	YRes            = 1080
	XMin            = -1.777
	XMax            = 1.777
	YMin            = -1.0
	YMax            = 1.0
	NumV            = 16
	Samples         = 20000
	Iterations      = 10000
	GammaCorrection = 2.2
	OutputPath      = "/tmp/fractal.png"
	ThreadGroupSize = runtime.NumCPU()
	Symmetry        = 1
	Invert          = false
	FractalType     = 1
)

type Config struct {
	OutputPath      string
	XRes, YRes      int
	XMin, YMin      float64
	XMax, YMax      float64
	RanX, RanY      float64
	GammaCorrection float64
	Iterations      int
	NumThreads      int
	NumV            int
	Samples         int
	Symmetry        int
	Count           int
	Invert          bool
}

func ParseArgs() *domain.Config {
	var (
		config        domain.Config
		R, G, B       uint
		mnR, mnG, mnB uint
		mxR, mxG, mxB uint
	)
	// Define flags
	flag.StringVar(&config.OutputPath, "f", OutputPath, "File to write")
	flag.IntVar(&config.Symmetry, "S", 1, "Enable rotational symmetry axis")
	flag.BoolVar(&config.Invert, "I", false, "Invert colors in final image")
	flag.BoolVar(&config.Debug, "D", false, "Allow debug info")
	flag.IntVar(&config.XRes, "x", XRes, "Image x resolution")
	flag.IntVar(&config.YRes, "y", YRes, "Image y resolution")
	flag.Float64Var(&config.XMin, "m", XMin, "Graph x minimum")
	flag.Float64Var(&config.XMax, "M", XMax, "Graph x maximum")
	flag.Float64Var(&config.YMin, "l", YMin, "Graph y minimum")
	flag.Float64Var(&config.YMax, "L", YMax, "Graph y maximum")
	flag.IntVar(&config.NumV, "n", NumV, "Number of random vectors to use")
	flag.IntVar(&config.Samples, "s", Samples, "Number of image samples")
	flag.IntVar(&config.Iterations, "i", Iterations, "Number of iterations per sample")
	flag.Float64Var(&config.GammaCorrection, "Gc", GammaCorrection, "Correctional gamma factor")
	flag.IntVar(&config.NumThreads, "T", ThreadGroupSize, "Number of threads to run")
	flag.IntVar(&config.FractalType, "t", FractalType, "Fractal type, use equation by number")
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
	config.StandartColor = color.RGBA{uint8(R), uint8(G), uint8(B), 255}
	config.MaxColorRange = color.RGBA{uint8(mxR), uint8(mxG), uint8(mxB), 255}
	config.MinColorRange = color.RGBA{uint8(mnR), uint8(mnG), uint8(mnB), 255}
	config.RanX, config.RanY = config.XMax-config.XMin, config.YMax-config.YMin
	return &config
}

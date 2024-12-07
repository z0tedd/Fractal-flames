package application

import (
	"flag"
	"fmt"
	"runtime"
	"time"
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
	Seed            = 1
	OutputPath      = "/tmp/fractal.png"
	ThreadGroupSize = runtime.NumCPU()
	Symmetry        = 1
	R               = -1
	G               = -1
	B               = -1
	Invert          = false
	Choice          = []int{0}
	FractalType     = 1
)

func ParseArgs() {
	// Define flags
	flag.IntVar(&Seed, "R", int(time.Now().UTC().UnixNano()), "Seed randomizer")
	flag.StringVar(&OutputPath, "f", OutputPath, "File to write")
	flag.IntVar(&Symmetry, "S", 1, "Enable rotational symmetry axis")
	flag.BoolVar(&Invert, "I", false, "Invert colors in final image")
	flag.IntVar(&XRes, "x", XRes, "Image x resolution")
	flag.IntVar(&YRes, "y", YRes, "Image y resolution")
	flag.Float64Var(&XMin, "m", XMin, "Graph x minimum")
	flag.Float64Var(&XMax, "M", XMax, "Graph x maximum")
	flag.Float64Var(&YMin, "l", YMin, "Graph y minimum")
	flag.Float64Var(&YMax, "L", YMax, "Graph y maximum")
	flag.IntVar(&NumV, "n", NumV, "Number of random vectors to use")
	flag.IntVar(&Samples, "s", Samples, "Number of image samples")
	flag.IntVar(&Iterations, "i", Iterations, "Number of iterations per sample")
	flag.IntVar(&R, "r", -1, "Static RED channel value")
	flag.IntVar(&G, "g", -1, "Static GREEN channel value")
	flag.IntVar(&B, "b", -1, "Static BLUE channel value")
	flag.Float64Var(&GammaCorrection, "G", GammaCorrection, "Correctional gamma factor")
	flag.IntVar(&ThreadGroupSize, "T", ThreadGroupSize, "Number of threads to run")
	flag.IntVar(&FractalType, "t", FractalType, "Fractal type, use equation by number")

	// Custom usage message
	flag.Usage = func() {
		printUsage()
	}

	// Parse command-line arguments
	flag.Parse()
}

func printUsage() {
	fmt.Println("Usage: fractal [options]")
	fmt.Println("\nOptions:")
	fmt.Println("\t-h\t\t\tPrint this help message")
	fmt.Printf("\t-R NUM\t\t\tSeed randomizer with NUM (default: %d)\n", Seed)
	fmt.Printf("\t-f NAME\t\t\tFile to write (default: %s)\n", OutputPath)
	fmt.Printf("\t-S NUM>1\t\tEnable rotational symmetry axis (default: 1)\n")
	fmt.Println("\t-I\t\t\tInvert colors in final image")
	fmt.Printf("\t-x XRES\t\t\tImage x resolution (default: %d)\n", XRes)
	fmt.Printf("\t-y YRES\t\t\tImage y resolution (default: %d)\n", YRes)
	fmt.Printf("\t-m XMIN\t\t\tGraph x minimum (default: %f)\n", XMin)
	fmt.Printf("\t-M XMAX\t\t\tGraph x maximum (default: %f)\n", XMax)
	fmt.Printf("\t-l YMIN\t\t\tGraph y minimum (default: %f)\n", YMin)
	fmt.Printf("\t-L YMAX\t\t\tGraph y maximum (default: %f)\n", YMax)
	fmt.Printf("\t-n NUMV\t\t\tNumber of random vectors to use (default: %d)\n", NumV)
	fmt.Printf("\t-s SAMPLES\t\tNumber of image samples (default: %d)\n", Samples)
	fmt.Printf("\t-i NUM>20\t\tNumber of iterations per sample (default: %d)\n", Iterations)
	fmt.Println("\t-r NUM\t\t\tSet static RED channel value (0-255)")
	fmt.Println("\t-g NUM\t\t\tSet static GREEN channel value (0-255)")
	fmt.Println("\t-b NUM\t\t\tSet static BLUE channel value (0-255)")
	fmt.Printf("\t-G NUM\t\t\tCorrectional gamma factor (default: %f)\n", GammaCorrection)
	fmt.Printf("\t-T THREADS\t\tNumber of threads to run (default: %d)\n", ThreadGroupSize)
	fmt.Println("\t-v NUM\t\t\tUse equation by number (see below for options)")
	fmt.Println("\nTransformation values for -t:")
	fmt.Println("\t0: Linear, 1: Sinusoidal, 2: Spherical, 3: Swirl, 4: Horseshoe, ...")
}

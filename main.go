package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/image/bmp"
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
	Iterations      = 1000
	SuperSampling   = 1
	GammaCorrection = 2.2
	Seed            = 1
	OutputPath      = "/tmp/fractal.png"
	ThreadGroupSize = 9
	PaletteFile     = ""
	CoeffFile       = ""
	Symmetry        = 1
	R               = -1
	G               = -1
	B               = -1
	Invert          = false
	Choice          = []int{0}
)

//	type Coeff struct {
//		A, B, C, D, E, F   float64 // Transformation coefficients
//		PA1, PA2, PA3, PA4 float64
//		R, G, B            uint8 // RGB color values
//	}
// type Pixel struct {
// 	R, G, B uint8
// 	Value   uint32
// 	Normal  float64
// }

type HitCounter struct {
	Counter uint32  // Number of times a pixel has been hit
	Normal  float64 // Normalized pixel value
}

type Pixel struct {
	Value HitCounter
	Color color.RGBA // RGB color for a pixel
}

type Flame struct {
	XRes, YRes      int
	XMin, YMin      float64
	XMax, YMax      float64
	RanX, RanY      float64
	R, G, B         int
	N               int // Number of equations
	SuperSampling   int
	Samples         int
	Iterations      int64
	Invert          bool
	Symmetry        int
	Seed            int64
	NumThreads      int
	GammaCorrection float64
	File            string
	Coefficients    []Coeff
	Pixels          [][]Pixel
	Choice          []int
	Count           int
	PaletteFile     *os.File
	CoeffFile       *os.File
	Locks           []sync.Mutex
}

// Apply gamma correction and logarithmic normalization
func gammaLog(fractal *Flame) {
	var max1 float64 = 0.0

	// Find the maximum logarithmic value
	for row := 0; row < fractal.YRes; row++ {
		for col := 0; col < fractal.XRes; col++ {
			if fractal.Pixels[row][col].Value.Counter != 0 {

				normal := math.Log(float64(fractal.Pixels[row][col].Value.Counter))
				fractal.Pixels[row][col].Value.Normal = normal
				if normal > max1 {
					max1 = normal
				}
			}
		}
	}

	// Normalize and apply gamma correction
	for row := 0; row < fractal.YRes; row++ {
		for col := 0; col < fractal.XRes; col++ {
			pixel := &fractal.Pixels[row][col]
			pixel.Value.Normal /= max1
			// if pixel.Value.Normal != 0 {
			gammaFactor := math.Pow(float64(pixel.Value.Normal), 1.0/GammaCorrection)
			pixel.Color.R = uint8(float64(pixel.Color.R) * gammaFactor)
			pixel.Color.G = uint8(float64(pixel.Color.G) * gammaFactor)
			pixel.Color.B = uint8(float64(pixel.Color.B) * gammaFactor)
			// }
		}
	}
}

// Reduce a high-resolution fractal image to a lower resolution
func reduce(fractal *Flame) {
	sample := fractal.SuperSampling
	// _ := fractal.YRes
	newXRes := fractal.XRes / sample
	newYRes := fractal.YRes / sample

	// Allocate new reduced pixel array
	reduction := make([][]Pixel, newYRes)
	for y := range reduction {
		reduction[y] = make([]Pixel, newXRes)
	}

	// Perform reduction with anti-aliasing
	for y := 0; y < newYRes; y++ {
		for x := 0; x < newXRes; x++ {
			var rSum, gSum, bSum uint32
			var count uint32

			for sy := 0; sy < sample; sy++ {
				for sx := 0; sx < sample; sx++ {
					srcPixel := fractal.Pixels[y*sample+sy][x*sample+sx]
					rSum += uint32(srcPixel.Color.R)
					gSum += uint32(srcPixel.Color.G)
					bSum += uint32(srcPixel.Color.B)
					count += srcPixel.Value.Counter
				}
			}

			// Calculate the averaged pixel values
			reduction[y][x] = Pixel{
				Value: HitCounter{Counter: count, Normal: fractal.Pixels[y][x].Value.Normal},
				Color: color.RGBA{
					R: uint8(rSum / uint32(sample*sample)),
					G: uint8(gSum / uint32(sample*sample)),
					B: uint8(bSum / uint32(sample*sample)),
				},
			}
			// reduction[y][x] = Pixel{
			// 	R:     uint8(rSum / uint32(sample*sample)),
			// 	G:     uint8(gSum / uint32(sample*sample)),
			// 	B:     uint8(bSum / uint32(sample*sample)),
			// 	Value: count,
			// }
		}
	}

	// Free the old pixel data and assign the reduced data
	fractal.Pixels = reduction
	fractal.XRes = newXRes
	fractal.YRes = newYRes
}

// type Pixel struct {
// 	R, G, B    uint8 // RGB values
// 	Value      uint32
// 	Normalized float64
// }

// type Coeff struct {
// 	A, B, C, D, E, F   float64
// 	PA1, PA2, PA3, PA4 float64
// 	R, G, B            uint8
// }

// func randomBit() bool {
// 	return rand.Intn(2) == 1
// }

func randomRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func applyTransformation(k int, x, y float64, coeff Coeff) (float64, float64) {
	var newx, newy float64
	var r, theta, _, _ float64
	// pa1, pa2, pa3, _ := coeff.PA1, coeff.PA2, coeff.PA3, coeff.PA4
	// _, f, _, _ := coeff.CC, coeff.FC, coeff.BC, coeff.EC

	switch k {
	case 0: // Linear
		newx = x
		newy = y
	case 1: // Sinusoidal
		newx = math.Sin(x)
		newy = math.Sin(y)
	case 2: // Spherical
		r = 1.0 / (x*x + y*y)
		newx = r * x
		newy = r * y
	case 3: // Swirl
		r = x*x + y*y
		newx = x*math.Sin(r) - y*math.Cos(r)
		newy = x*math.Cos(r) + y*math.Sin(r)
	case 4: // Horseshoe
		r = 1.0 / math.Sqrt(x*x+y*y)
		newx = r * (x - y) * (x + y)
		newy = r * 2.0 * x * y
	case 5: // Polar
		newx = math.Atan2(y, x) / math.Pi
		newy = math.Sqrt(x*x+y*y) - 1.0
	case 6: // Handkerchief
		r = math.Sqrt(x*x + y*y)
		theta = math.Atan2(y, x)
		newx = r * math.Sin(theta+r)
		newy = r * math.Cos(theta-r)
	// Add more cases as needed for all transformations...
	default:
		newx = x
		newy = y
	}

	return newx, newy
}

func renderWorker(flame *Flame, wg *sync.WaitGroup) {
	defer wg.Done()

	var newx, newy float64

	for sample := 0; sample < flame.Samples; sample++ {
		newx = randomRange(flame.XMin, flame.XMax)
		newy = randomRange(flame.YMin, flame.YMax)

		for step := -20; step < int(flame.Iterations); step++ {
			// Select random coefficient
			// tran := rand.Intn(len(flame.Choice)) flame.Choice[tran]
			i := rand.Intn(flame.N)

			// Apply transformation
			coeff := flame.Coefficients[i]
			newx, newy = applyTransformation(4, newx, newy, coeff)

			// Skip the first few steps
			if step <= 0 {
				continue
			}

			// Apply symmetry
			for s := 0; s < flame.Symmetry; s++ {
				theta := 2 * math.Pi * float64(s) / float64(flame.Symmetry)
				xRot := newx*math.Cos(theta) - newy*math.Sin(theta)
				yRot := newx*math.Sin(theta) + newy*math.Cos(theta)

				if xRot < flame.XMin || xRot > flame.XMax || yRot < flame.YMin || yRot > flame.YMax {
					continue
				}

				x1 := flame.XRes - int(((flame.XMax-xRot)/flame.RanX)*float64(flame.XRes))
				y1 := flame.YRes - int(((flame.YMax-yRot)/flame.RanY)*float64(flame.YRes))

				if x1 >= 0 && x1 < flame.XRes && y1 >= 0 && y1 < flame.YRes {
					// Safely update the pixel
					flame.Locks[y1].Lock()
					pixel := &flame.Pixels[y1][x1]
					if pixel.Value.Counter == 0 {
						pixel.Color.R = coeff.R
						pixel.Color.G = coeff.G
						pixel.Color.B = coeff.B
					} else {
						pixel.Color.R = (pixel.Color.R + coeff.R) / 2
						pixel.Color.G = (pixel.Color.G + coeff.G) / 2
						pixel.Color.B = (pixel.Color.B + coeff.B) / 2
					}
					pixel.Value.Counter++
					flame.Locks[y1].Unlock()
				}
			}
		}
	}
}

func render(flame *Flame) {
	var wg sync.WaitGroup
	for t := 0; t < flame.NumThreads; t++ {
		wg.Add(1)
		go renderWorker(flame, &wg)
	}
	wg.Wait()
}

func randR(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
} // writeToJPEG generates a JPEG image based on the Flame struct
func writeToJPEG(fractal *Flame, quality int) {
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
		fmt.Fprintf(os.Stderr, "Could not write image: %v\n", err)
		os.Exit(1)
	}
}

// writeToBMP generates a BMP image based on the Flame struct
func writeToBMP(fractal *Flame) {
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
		fmt.Fprintf(os.Stderr, "Could not write image: %v\n", err)
		os.Exit(1)
	}
}

func render2(fractal *Flame) {
	var tran int

	for num := 0; num < fractal.Samples; num++ {
		newX := randR(fractal.XMin, fractal.XMax)
		newY := randR(fractal.YMin, fractal.YMax)

		for step := -20; step < int(fractal.Iterations); step++ {
			k := 2 // fractal.Choice[tran%fractal.Count]
			tran++
			i := rand.Intn(fractal.N)

			// pa1, pa2, pa3, pa4 := fractal.Coefficients[i].PA1, fractal.Coefficients[i].PA2, fractal.Coefficients[i].PA3, fractal.Coefficients[i].PA4
			c, f, b, e := fractal.Coefficients[i].CC, fractal.Coefficients[i].FC, fractal.Coefficients[i].BC, fractal.Coefficients[i].EC
			ac, dc := fractal.Coefficients[i].AC, fractal.Coefficients[i].DC

			x := ac*newX + b*newY + c
			y := dc*newX + e*newY + f

			switch k {
			case 0: // Linear
				newX, newY = x, y
			case 1: // Sinusoidal
				newX, newY = math.Sin(x), math.Sin(y)
			case 2: // Spherical
				r := 1.0 / (x*x + y*y)
				newX, newY = r*x, r*y
			case 3: // Swirl
				r := x*x + y*y
				newX, newY = x*math.Sin(r)-y*math.Cos(r), x*math.Cos(r)+y*math.Sin(r)
			case 4: // Horseshoe
				r := 1.0 / math.Sqrt(x*x+y*y)
				newX, newY = r*(x-y)*(x+y), r*2.0*x*y
				// Add other transformations similarly...
			}

			if step > 0 {
				for s := 0; s < fractal.Symmetry; s++ {
					theta := float64(s) * (2 * math.Pi / float64(fractal.Symmetry))
					xRot := newX*math.Cos(theta) - newY*math.Sin(theta)
					yRot := newX*math.Sin(theta) + newY*math.Cos(theta)

					if xRot >= fractal.XMin && xRot <= fractal.XMax && yRot >= fractal.YMin && yRot <= fractal.YMax {
						x1 := fractal.XRes - int((fractal.XMax-xRot)/fractal.RanX*float64(fractal.XRes))
						y1 := fractal.YRes - int((fractal.YMax-yRot)/fractal.RanY*float64(fractal.YRes))

						if x1 >= 0 && x1 < fractal.XRes && y1 >= 0 && y1 < fractal.YRes {
							// fractal.Lock[y1].Lock()
							point := &fractal.Pixels[y1][x1]

							if point.Value.Counter == 0 {
								point.Color.R = fractal.Coefficients[i].R
								point.Color.G = fractal.Coefficients[i].G
								point.Color.B = fractal.Coefficients[i].B
							} else {
								point.Color.R = (point.Color.R + fractal.Coefficients[i].R) / 2
								point.Color.G = (point.Color.G + fractal.Coefficients[i].G) / 2
								point.Color.B = (point.Color.B + fractal.Coefficients[i].B) / 2
							}
							point.Value.Counter++
							// fractal.Lock[y1].Unlock()
						}
					}
				}
			}
		}
	}
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
	fmt.Printf("\t-sup NUM\t\tSuper sample NUM^2 buckets (default: %d)\n", SuperSampling)
	fmt.Printf("\t-G NUM\t\t\tCorrectional gamma factor (default: %f)\n", GammaCorrection)
	fmt.Println("\t-p FILE\t\t\tUse input color palette file")
	fmt.Println("\t-c FILE\t\t\tUse coefficient input file")
	fmt.Printf("\t-T THREADS\t\tNumber of threads to run (default: %d)\n", ThreadGroupSize)
	fmt.Println("\t-v NUM\t\t\tUse equation by number (see below for options)")
	fmt.Println("\nTransformation values for -v:")
	fmt.Println("\t0: Linear, 1: Sinusoidal, 2: Spherical, 3: Swirl, 4: Horseshoe, ...")
}

func parseArgs() {
	// Define flags
	flag.IntVar(&Seed, "R", Seed, "Seed randomizer")
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
	flag.IntVar(&SuperSampling, "sup", SuperSampling, "Super sample NUM^2 buckets")
	flag.Float64Var(&GammaCorrection, "G", GammaCorrection, "Correctional gamma factor")
	flag.StringVar(&PaletteFile, "p", "", "Input color palette file")
	flag.StringVar(&CoeffFile, "c", "", "Coefficient input file")
	flag.IntVar(&ThreadGroupSize, "T", ThreadGroupSize, "Number of threads to run")

	// Handle non-standard flags
	var choiceFlag string
	flag.StringVar(&choiceFlag, "v", "0", "Use equation by number")

	// Custom usage message
	flag.Usage = func() {
		printUsage()
	}

	// Parse command-line arguments
	flag.Parse()

	// Apply parsed flags

	// Parse -v (choice of equations)
	if choiceFlag != "" {
		for _, v := range flag.Args() {
			num, err := strconv.Atoi(v)
			fmt.Print(num, " ")
			if err == nil {
				Choice = append(Choice, num)
			}
		}
	}
}

// Initialize default flame structure values
func fractalInit(fractal *Flame) {
	fractal.XRes = XRes
	fractal.YRes = YRes
	fractal.XMin = XMin
	fractal.YMin = YMin
	fractal.XMax = XMax
	fractal.YMax = YMax
	fractal.N = NumV
	fractal.Seed = int64(Seed)
	fractal.Samples = Samples
	fractal.Symmetry = Symmetry
	fractal.Invert = Invert
	fractal.File = OutputPath
	fractal.SuperSampling = SuperSampling
	fractal.Iterations = int64(Iterations)
	fractal.NumThreads = ThreadGroupSize
	fractal.R = -1
	fractal.G = -1
	fractal.B = -1
	fractal.Count = 1
	fractal.Choice = Choice
	fractal.Choice[0] = 0
}

// Allocate memory for the image buffer
func bufferInit(fractal *Flame) {
	// Ensure resolution is valid
	if fractal.XRes <= 0 {
		fractal.XRes = XRes
	}
	if fractal.YRes <= 0 {
		fractal.YRes = YRes
	}

	// Ensure proper min/max values
	if fractal.XMin > fractal.XMax {
		fractal.XMin, fractal.XMax = fractal.XMax, fractal.XMin
	}
	if fractal.YMin > fractal.YMax {
		fractal.YMin, fractal.YMax = fractal.YMax, fractal.YMin
	}

	// Calculate ranges
	fractal.RanX = fractal.XMax - fractal.XMin
	fractal.RanY = fractal.YMax - fractal.YMin

	// Apply super-sampling
	fractal.XRes *= fractal.SuperSampling
	fractal.YRes *= fractal.SuperSampling

	// Memory allocation for pixels
	// fmt.Printf("Attempting to allocate %.2f MiB of RAM.\n",
	// 	float64(fractal.XRes*fractal.YRes*int(math.Ceil(float64(len(Pixel{nil,nil})))))/(1024.0*1024.0))

	fractal.Pixels = make([][]Pixel, fractal.YRes)
	fractal.Locks = make([]sync.Mutex, fractal.YRes)

	for y := 0; y < fractal.YRes; y++ {
		fractal.Pixels[y] = make([]Pixel, fractal.XRes)
		// Mutex for row-level locking
		fractal.Locks[y] = sync.Mutex{}
	}
	fmt.Println("Buffer initialization complete!")
}

type Coeff struct {
	AC, BC, CC, DC, EC, FC float64 // Transformation coefficients
	PA1, PA2, PA3, PA4     float64
	R, G, B                uint8 // RGB color values
}

// Utility functions
func randomBit() int {
	return rand.Intn(2)
}

func modulus(a, b float64) float64 {
	return math.Mod(a, b)
}

func randRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Contractive mapping function
func contractiveMapping(coeff *Coeff) {
	var a, b, d, e float64

	for {
		// Generate valid values for a, d
		for {
			a = rand.Float64()
			d = randRange(a*a, 1)
			if randomBit() == 1 {
				d = -d
			}
			if a*a+d*d <= 1 {
				break
			}
		}

		// Generate valid values for b, e
		for {
			b = rand.Float64()
			e = randRange(b*b, 1)
			if randomBit() == 1 {
				e = -e
			}
			if b*b+e*e <= 1 {
				break
			}
		}

		// Check final condition
		if a*a+b*b+d*d+e*e <= 1+(a*e-d*b)*(a*e-d*b) {
			break
		}
	}

	coeff.AC = a
	coeff.BC = b
	coeff.CC = randRange(-2, 2)
	coeff.DC = d
	coeff.EC = e
	coeff.FC = randRange(-2, 2)
}

// Initialize coefficients for the fractal
func coeffInit(fractal *Flame) {
	fractal.Coefficients = make([]Coeff, fractal.N)

	for i := 0; i < fractal.N; i++ {
		if randomBit() == 1 {
			contractiveMapping(&fractal.Coefficients[i])
		} else {
			fractal.Coefficients[i] = Coeff{
				AC: randRange(-1.5, 1.5),
				BC: randRange(-1.5, 1.5),
				CC: randRange(-1.5, 1.5),
				DC: randRange(-1.5, 1.5),
				EC: randRange(-1.5, 1.5),
				FC: randRange(-1.5, 1.5),
			}
		}

		// Assign additional parameters
		fractal.Coefficients[i].PA1 = randRange(-2, 2)
		fractal.Coefficients[i].PA2 = randRange(-2, 2)
		fractal.Coefficients[i].PA3 = randRange(-2, 2)
		fractal.Coefficients[i].PA4 = randRange(-2, 2)

		// Assign colors
		fractal.Coefficients[i].R = uint8(fractal.R)
		fractal.Coefficients[i].G = uint8(fractal.G)
		fractal.Coefficients[i].B = uint8(fractal.B)

		if fractal.R == -1 {
			fractal.Coefficients[i].R = uint8(randRange(64, 256))
		}
		if fractal.G == -1 {
			fractal.Coefficients[i].G = uint8(randRange(64, 256))
		}
		if fractal.B == -1 {
			fractal.Coefficients[i].B = uint8(randRange(64, 256))
		}
	}

	// Load palette file if provided
	if fractal.PaletteFile != nil {
		i := 0
		var r, g, b int
		for _, err := fmt.Fscanf(fractal.PaletteFile, "%d %d %d\n", &r, &g, &b); err == nil && i < fractal.N; {
			fractal.Coefficients[i].R = uint8(r)
			fractal.Coefficients[i].G = uint8(g)
			fractal.Coefficients[i].B = uint8(b)
			fmt.Printf("Setting index %d to %d,%d,%d\n", i, r, g, b)
			i++
		}
		err := fractal.PaletteFile.Close()
		if err != nil {
			log.Println("Failed to close palette file:", err)
		}
	}

	// Load coefficients file if provided
	if fractal.CoeffFile != nil {
		i := 0
		var ac, bc, cc, dc, ec, fc float64
		for _, err := fmt.Fscanf(fractal.CoeffFile, "%f %f %f %f %f %f\n", &ac, &bc, &cc, &dc, &ec, &fc); err == nil && i < fractal.N; {
			fractal.Coefficients[i] = Coeff{AC: ac, BC: bc, CC: cc, DC: dc, EC: ec, FC: fc}
			fmt.Printf("Setting index coeffs at index %d\n", i)
			i++
		}
		err := fractal.CoeffFile.Close()
		if err != nil {
			log.Println("Failed to close coefficients file:", err)
		}
	}

	// Print coefficients for debugging
	for _, coeff := range fractal.Coefficients {
		fmt.Printf("%f %f %f %f %f %f\n", coeff.AC, coeff.BC, coeff.CC, coeff.DC, coeff.EC, coeff.FC)
	}
	fmt.Println("Colors")
	for _, coeff := range fractal.Coefficients {
		fmt.Printf("%d %d %d\n", coeff.R, coeff.G, coeff.B)
	}
}

func writeToPng(fractal *Flame, fileName string) {
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

	log.Printf("PNG file successfully written: %s\n", fileName)
}

func main() {
	var fractal Flame
	// var wg sync.WaitGroup

	// Parse arguments from the command line
	fmt.Println("Parsing user arguments...")
	parseArgs()
	fmt.Println("Done!")

	// Initialize our Flame Fractal
	fmt.Println("Initializing...")
	fractalInit(&fractal)
	fmt.Println("Initialized!")

	// Seed the randomizer
	rand.Seed(time.Now().UnixNano())

	fmt.Println(fractal.Seed)

	// Initialize the random coefficients
	fmt.Println("Initializing Coefficients and Colors...")
	coeffInit(&fractal)
	fmt.Println("Done!")

	// Allocate our memory buffer
	fmt.Println("Allocating memory...")
	bufferInit(&fractal)
	fmt.Println("Done!")

	// Correct for threads
	if fractal.NumThreads <= 0 {
		fractal.NumThreads = 1
	}

	fractal.Samples /= fractal.NumThreads

	// render(&fractal)
	render2(&fractal)
	// Gamma and log correct
	fmt.Println("Finalizing and writing out...")
	gammaLog(&fractal)
	// if fractal.SuperSampling > 1 {
	// 	reduce(&fractal)
	// }

	// Write out the file
	writeToPng(&fractal, OutputPath)
	writeToBMP(&fractal)
	writeToJPEG(&fractal, 90)
	// Clean up
	fmt.Println("Cleaning up...")
	// freeResources(&fractal)
	fmt.Println("Done!")
}

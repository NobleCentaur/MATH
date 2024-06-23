package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"strings"
	"time"
)

const pi float64 = 3.141592653589793115997963468544185161590576171875

// Image parameters
// imageSize is the width and height in pixels, min 256
// imageSharpness is the complexity of the math operations
// ranges from 1 to the imageSize
// lower values take longer to calculate
// 5-10 is usually an acceptable balance between precision and speed
var imageSize int
var imageSharpness int = 1

// other important values
var valueRange float64 = 2.5
var maxIteration int
var startingX float64 = -2
var startingY float64 = 1.25
var complexNum complex128
var step float64

// coloring variables
var clr1 float64
var clr2 float64
var clr3 float64
var escapeTime float64
var gradientScale float64 = 0.1
var varTempFloat64 float64

// simple, unoptimized escape time algorithm
func escapeTimeAlgorithm(c complex128) uint64 {
	var z complex128
	var n uint64
	for i := 0; cmplx.Abs(z) < 2 && i < maxIteration; i++ {
		z = z*z + c
		n++
	}
	return (n)
}

// runs the escape time algorith for a given row and returns to a given channel
func escapeTimeAlgorithmByRow(rowNum int, array [][]uint64, channel chan bool) {
	for j := 0; j < imageSize; j++ {
		complexNum = complex((startingX + (float64(j) * step)), (startingY - (float64(rowNum) * step)))
		array[j][rowNum] = escapeTimeAlgorithm(complexNum)
	}
	channel <- true
}

// give it a percent and it will show a progress bar displaying that percent
func progressBar(percent int) {
	fmt.Print("\r[")
	fmt.Print(strings.Repeat("#", percent/2))
	fmt.Print(strings.Repeat("-", 50-(percent/2)))
	fmt.Print("] ")
	fmt.Print(fmt.Sprint(percent) + "%")
}

func render(adv bool) {
	//ensures that imageSize is even
	if imageSize%2 != 0 {
		imageSize += 1
	}

	maxIteration = imageSize / imageSharpness
	//option to change default parameters manually
	if adv {
		fmt.Print("gradientScale     :")
		fmt.Scan(&gradientScale)
	}

	step = valueRange / (float64(imageSize) - 1)

	ch := make(chan bool, imageSize/2)

	//creates blank 2d array with width and heighth of imageSize
	escapeTimeTable := make([][]uint64, imageSize)
	for i := range escapeTimeTable {
		escapeTimeTable[i] = make([]uint64, imageSize)
	}

	//time check
	start := time.Now()

	// multiprocessed version of the math
	fmt.Println("spawning processes")
	for j := 0; j < imageSize/2; j++ {
		go escapeTimeAlgorithmByRow(j, escapeTimeTable, ch)
	}
	fmt.Println("calculating...")
	fmt.Println("")
	// joins all processes
	var percent float64
	print("[--------------------------------------------------] 0%")
	for j := 0; j < imageSize/2; j++ {
		percent = ((float64(j+1) / (float64(imageSize) / 2)) * 100)
		progressBar(int(percent))
		<-ch
	}
	fmt.Println("")
	fmt.Println("")

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println("end math, start interpretation")

	//copies top half to bottom half
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize/2; k++ {
			escapeTimeTable[j][imageSize/2+k] = escapeTimeTable[j][imageSize/2-k]
		}
	}

	//escape time color normalization
	escapeHistogram := make([]uint64, maxIteration)
	var varTempUint64 uint64
	for j := 0; j < maxIteration; j++ {
		for k := 0; k < maxIteration/2; k++ {
			varTempUint64 = escapeTimeTable[j][k]
			escapeHistogram[varTempUint64-1]++
		}
	}

	// MUCH improved coloration algorithm
	// https://www.desmos.com/calculator/gmbe5ekk3z
	// desmos project shows the math behind it so you
	// don't have to read this gross spaghetti code
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize; k++ {
			escapeTime = float64(escapeTimeTable[j][k])
			if escapeTime == float64(maxIteration) {
				img.Set(j, k, color.RGBA{0, 0, 0, 255})
			} else {
				varTempFloat64 = (gradientScale * escapeTime)
				clr1 = 255 * ((math.Sin(varTempFloat64) + 1) / 2)
				varTempFloat64 = gradientScale * (escapeTime + ((2 * pi) / (3 * gradientScale)))
				clr2 = 255 * ((math.Sin(varTempFloat64) + 1) / 2)
				varTempFloat64 = gradientScale * (escapeTime + ((4 * pi) / (3 * gradientScale)))
				clr3 = 255 * ((math.Sin(varTempFloat64) + 1) / 2)
				img.Set(j, k, color.RGBA{uint8(clr2), uint8(clr1), uint8(clr3), 255})
			}
		}
	}

	//renders the image
	f1, err := os.OpenFile("renders/output.png", os.O_WRONLY|os.O_CREATE, 0600)
	errHandler(err)
	defer f1.Close()
	png.Encode(f1, img)
	//time check
	duration = time.Since(start)
	fmt.Println(duration)
	//
	fmt.Println("done")
	fmt.Println("")

	// Exports escape histogram to CSV
	f2, err := os.Create("data/escapeHistogram.csv")
	errHandler(err)
	defer f2.Close()
	csvWrite1 := csv.NewWriter(f2)
	escapeHistogramString := make([]string, len(escapeHistogram))
	for i := 0; i < len(escapeHistogram); i++ {
		escapeHistogramString[i] = fmt.Sprint(escapeHistogram[i])
	}
	csvWrite1.Write(escapeHistogramString)

	//exports escape time table to csv
	f3, _ := os.Create("data/escapeTimeTable.csv")
	errHandler(err)
	defer f3.Close()
	csvWrite2 := csv.NewWriter(f3)
	escapeTimeTableString := make([][]string, imageSize)
	for i := range escapeTimeTable {
		escapeTimeTableString[i] = make([]string, imageSize)
	}
	for i := 0; i < imageSize; i++ {
		for j := 0; j < imageSize; j++ {
			escapeTimeTableString[i][j] = fmt.Sprint(escapeTimeTable[i][j])
		}
	}
	csvWrite2.WriteAll(escapeTimeTableString)

}

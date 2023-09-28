package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"time"
)

// Image parameters
// imageSize is the width and height in pixels, min 256
// imageSharpness is the complexity of the math operations
// ranges from 1 to the imageSize
// lower values take longer to calculate
// 5-10 is usually an acceptable balance between precision and speed
var imageSize int = 5000
var imageSharpness int = 1

// other important values
var valueRange float64 = 2.5
var maxIteration int = imageSize / imageSharpness
var startingX float64 = -2
var startingY float64 = 1.25
var complexNum complex128
var step float64 = valueRange / (float64(imageSize) - 1)

// gradient starting and ending rgb values
var startingRed float64 = 0
var startingGreen float64 = 0
var startingBlue float64 = 0
var endingRed float64 = 255
var endingGreen float64 = 0
var endingBlue float64 = 0
var clrRed float64
var clrBlue float64
var clrGreen float64

// simple, unoptimized escape time algorithm
func escapeTimeAlgorithm(c complex128) uint8 {
	var z complex128
	var n uint8
	for i := 0; cmplx.Abs(z) < 2 && i < maxIteration; i++ {
		z = z*z + c
		n++
	}
	return (n)
}

// runs the escape time algorith for a given row and returns to a given channel
func escapeTimeAlgorithmByRow(rowNum int, array [][]uint8) {
	for j := 0; j < imageSize; j++ {
		complexNum = complex((startingX + (float64(j) * step)), (startingY - (float64(rowNum) * step)))
		array[j][rowNum] = escapeTimeAlgorithm(complexNum)
	}
}

func main() {
	//ensures that imageSize is even
	if imageSize%2 != 0 {
		imageSize += 1
	}

	//creates blank 2d array with width and heighth of imageSize
	escapeTimeTable := make([][]uint8, imageSize)
	for i := range escapeTimeTable {
		escapeTimeTable[i] = make([]uint8, imageSize)
	}

	//time check
	start := time.Now()

	// multiprocessed version of the math
	fmt.Println("start math")
	for j := 0; j < imageSize/2; j++ {
		go escapeTimeAlgorithmByRow(j, escapeTimeTable)
	}

	//copies top half to bottom half
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize/2; k++ {
			escapeTimeTable[j][imageSize/2+k] = escapeTimeTable[j][imageSize/2-k]
		}
	}

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println("end math, start interpretation")

	//escape time color normalization
	escapeHistogram := make([]uint8, maxIteration)
	var varTemp uint8
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize; k++ {
			varTemp = escapeTimeTable[j][k]
			escapeHistogram[varTemp-1]++
		}
	}

	//gradient adjustments
	var dividendAdjusted float64
	var count int = bytes.Count(escapeHistogram, []byte{0})
	dividendAdjusted = float64(maxIteration) - float64(count)

	//interprets the escape time table to an image
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize; k++ {
			var escapeTime = escapeTimeTable[j][k]
			if escapeTime == uint8(maxIteration) {
				img.Set(j, k, color.RGBA{0, 0, 0, 255})
			} else {
				clrRed = (startingRed + ((endingRed-startingRed)/dividendAdjusted)*float64(escapeTime))
				clrBlue = (startingBlue + ((endingBlue-startingBlue)/dividendAdjusted)*float64(escapeTime))
				clrGreen = (startingGreen + ((endingGreen-startingGreen)/dividendAdjusted)*float64(escapeTime))
				img.Set(j, k, color.RGBA{uint8(clrRed), uint8(clrBlue), uint8(clrGreen), 255})
			}
		}
	}

	//renders the image
	f, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	//time check
	duration = time.Since(start)
	fmt.Println(duration)
	//
	fmt.Println("done")
}

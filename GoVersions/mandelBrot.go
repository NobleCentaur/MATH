package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"time"
)

// Image parameters
// imageSize is the width and height in pixels
// imageSharpness is the complexity of the math operations
// ranges from 1 to the imageSize
// lower values take longer to calculate
// 5-10 is usually an acceptable balance between precision and speed
var imageSize int = 1000
var imageSharpness int = 1

// other important values
var valueRange float64 = 2.5
var maxIteration int = imageSize / imageSharpness
var startingX float64 = -2
var startingY float64 = 1.25
var complexNum complex128
var step float64 = valueRange / (float64(imageSize) - 1)

// simple, unoptimized escape time algorithm
func escapeTimeAlgorithm(c complex128) uint8 {
	var z complex128
	var n uint8
	for i := 0; i < maxIteration && cmplx.Abs(z) < 2; i++ {
		z = z*z + c
		n++
	}
	return (n)
}

func main() {
	//creates blank 2d array with width and heighth of imageSize
	escapeTimeTable := make([][]uint8, imageSize)
	for i := range escapeTimeTable {
		escapeTimeTable[i] = make([]uint8, imageSize)
	}
	//time check
	start := time.Now()
	fmt.Println("start math")
	//does the mandelbrot test for every pixel
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize; k++ {
			complexNum = complex((startingX + (float64(j) * step)), (startingY - (float64(k) * step)))
			escapeTimeTable[j][k] = escapeTimeAlgorithm(complexNum)
		}
	}
	//gradient starting and ending rgb values
	var startingRed float32 = 0
	var startingGreen float32 = 0
	var startingBlue float32 = 0

	var endingRed float32 = 0
	var endingGreen float32 = 255
	var endingBlue float32 = 0

	var clrRed float32
	var clrBlue float32
	var clrGreen float32

	escapeHistogram := make([]uint8, maxIteration)
	var varTemp uint8
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize; k++ {
			varTemp = escapeTimeTable[j][k]
			escapeHistogram[varTemp-1]++
		}
	}
	fmt.Println(escapeHistogram)

	//interprets the escape time table to an image
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize; k++ {
			var escapeTime = escapeTimeTable[j][k]
			if escapeTime == uint8(maxIteration) {
				img.Set(j, k, color.RGBA{0, 0, 0, 255})
			} else {
				clrRed = (startingRed + ((endingRed-startingRed)/float32(maxIteration))*float32(escapeTime))
				clrBlue = (startingBlue + ((endingBlue-startingBlue)/float32(maxIteration))*float32(escapeTime))
				clrGreen = (startingGreen + ((endingGreen-startingGreen)/float32(maxIteration))*float32(escapeTime))
				img.Set(j, k, color.RGBA{uint8(clrRed), uint8(clrBlue), uint8(clrGreen), 255})
			}
		}
	}

	//renders the image
	f, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	//time check
	duration := time.Since(start)
	fmt.Println(duration)
	//
	fmt.Println("done")
}

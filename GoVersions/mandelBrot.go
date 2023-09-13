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

var valueRange float64 = 2.5
var imageSize int = 250
var imageSharpness int = 1
var maxIteration int = imageSize / imageSharpness
var startingX float64 = -2
var startingY float64 = 1.25
var complexNum complex128
var step float64 = valueRange / (float64(imageSize) - 1)

func escapeTimeAlgorithm(c complex128) uint8 {
	//uses the escape time algorithm
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
	//interprets the escape time table to an image
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	for j := 0; j < imageSize; j++ {
		for k := 0; k < imageSize; k++ {
			var escapeTime = escapeTimeTable[j][k]
			if escapeTime == uint8(maxIteration) {
				img.Set(j, k, color.RGBA{0, 0, 0, 255})
			} else {
				clr := (5 + (uint8((5-255)/maxIteration))*escapeTime)
				img.Set(j, k, color.RGBA{clr, 0, 0, 255})
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

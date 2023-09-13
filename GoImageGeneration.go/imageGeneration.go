package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var imageWidth int = 10
var imageHeight int = 10

func main() {
	//create img with {imageWidth} width and {imageHeight} height
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	img.Set(2, 3, color.RGBA{255, 0, 0, 255})
	f, _ := os.OpenFile("output.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

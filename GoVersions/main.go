package main

import (
	"fmt"
)

func main() {

	// COMMANDS LIST
	// render - creates new render
	// exit - exits program

	var input string

	for stop := false; !stop; {

		fmt.Print(">")
		fmt.Scan(&input)

		if input == "render" {
			fmt.Println("")
			fmt.Print("image Size (min 256): ")
			fmt.Scan(&imageSize)
			render(false)
		} else if input == "renderAdv" {
			fmt.Println("")
			fmt.Print("image Size (min 256): ")
			fmt.Scan(&imageSize)
			render(true)
		} else if input == "exit" {
			stop = true
		} else if input == "networkMain" {
			networkMain()
		} else if input == "networkWorker" {
			networkWorker()
		} else if input == "writetest" {
			arrayTest := []int{2, 4, 6, 7}
			writeTest("escapeTable/file.csv", arrayTest)
		} else {
			fmt.Println("Not a recognized command.")
		}

	}

}

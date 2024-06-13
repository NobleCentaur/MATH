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
		} else if input == "main" {
			networkMain()
		} else if input == "worker" {
			networkWorker()
		} else {
			fmt.Println("Not a recognized command.")
		}

	}

}

package main

import (
	"fmt"
	"net"
)

var key = []byte{124, 44, 32, 124, 33, 44, 32, 124, 124, 44, 32, 124, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var buf = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var workerList []string
var input string

var startingX_local float64
var startingY_local float64
var step_local float64
var imageSize_local float64

func addWorker() {
	fmt.Print("worker node ip: ")
	fmt.Scan(&input)
	conn, err := net.Dial("tcp", input+":8080")
	errHandler(err)
	_, err = conn.Write(key)
	errHandler(err)
	_, err = conn.Read(buf)
	errHandler(err)
	if Float64frombytes(buf) == Float64frombytes(key) {
		workerList = append(workerList, input)
		fmt.Println("Connection successful, [" + input + "] added to worker list")
	}
}

func errHandler(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func networkMain() {
	for stop := false; !stop; {
		fmt.Print("Main>")
		fmt.Scan(&input)
		if input == "addWorker" {
			addWorker()
		} else if input == "exit" {
			stop = true
		} else if input == "workerList" {
			for i := 0; i < len(workerList); i++ {
				fmt.Println(workerList[i])
			}
			if len(workerList) < 1 {
				fmt.Println("There are no workers currently added.")
			}
		} else {
			fmt.Println("Not a recognized command.")
		}
	}
}

func networkWorker() {
	//listen for connectiosn
	fmt.Println("listening for connection from main")
	ln, err := net.Listen("tcp", ":8080")
	errHandler(err)
	conn, err := ln.Accept()
	errHandler(err)
	fmt.Println("connection received from main")

	_, err = conn.Read(buf)
	errHandler(err)
	if Float64frombytes(buf) == Float64frombytes(key) {
		conn.Write(key)
	}

	conn.Close()

	for stop := false; !stop; {
		ln, err = net.Listen("tcp", ":8080")
		errHandler(err)
		conn, err = ln.Accept()
		errHandler(err)
		defer conn.Close()

		_, err = conn.Read(buf)
		errHandler(err)
		imageSize_local = Float64frombytes(buf)
		conn.Write([]byte{1})

		_, err = conn.Read(buf)
		errHandler(err)
		startingX_local = Float64frombytes(buf)
		conn.Write([]byte{1})

		_, err = conn.Read(buf)
		errHandler(err)
		startingY_local = Float64frombytes(buf)
		conn.Write([]byte{1})

		_, err = conn.Read(buf)
		errHandler(err)
		step_local = Float64frombytes(buf) / (float64(imageSize_local) - 1)
		// receive step

		fmt.Println(imageSize_local)
		fmt.Print(startingX_local)
		fmt.Println(startingY_local)
		fmt.Println(step_local)

		for stopNested := false; !stopNested; {
			stopNested = true
		}
	}
}

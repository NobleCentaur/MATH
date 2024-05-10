package main

import (
	"fmt"
	"net"
	"time"
)

var key = []byte{124, 44, 32, 124, 33, 44, 32, 124, 124, 44, 32, 124, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var buf = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var workerList []net.Conn
var input string
var memCapacity int

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
		workerList = append(workerList, conn)
		fmt.Println("Connection successful, [" + input + "] added to worker list")
	} else {
		networkMain()
	}
}

func errHandler(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

// only used for waitingRead. Allows waitingRead to run the below in a goroutine.
func readContinuous(ch chan []byte, eCh chan error, connection net.Conn) {
	for {
		data := make([]byte, 64)
		_, err := connection.Read(data)
		if err != nil {
			eCh <- err
			return
		}
		if Float64frombytes(data) != 0 {
			ch <- data
			break
		}
	}
}

// waits until conn.read actually includes data then sends it to ch channel
func waitingRead(connection net.Conn, timeout int) (float64, error) {
	ch := make(chan []byte)
	eCh := make(chan error)

	go readContinuous(ch, eCh, connection)

	ticker := time.Tick(time.Second * time.Duration(timeout))
	for {
		select {
		case data := <-ch:
			return Float64frombytes(data), nil
		case err := <-eCh:
			fmt.Println(err)
		case <-ticker:
			return 0, net.ErrClosed
		}
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
		} else if input == "Render" {
			networkRender()
		} else {
			fmt.Println("Not a recognized command.")
		}
	}
}

func networkWorker() {
	//listen for connections
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

	fmt.Println(waitingRead(conn, 1800))

	time.Sleep(60 * time.Second)

	//imageSize_local
	//startingX_local
	//startingY_local
	//step_local
	fmt.Println(imageSize_local)
	fmt.Print(startingX_local)
	fmt.Println(startingY_local)
	fmt.Println(step_local)
	for stopNested := false; !stopNested; {

	}
}

func networkRender() {
	fmt.Println("")
	fmt.Print("image Size (min 256): ")
	fmt.Scan(&imageSize)
	if imageSize%2 != 0 {
		imageSize += 1
	}
	areYouSure := "n"
	for areYouSure != "y" {
		fmt.Print("Allowed memory usage (MB): ")
		fmt.Scan(&memCapacity)
		fmt.Println("Using " + fmt.Sprint(memCapacity) + " MB of memory")
		fmt.Println("[ATTENTION] Using more ram than you device has will crash it.")
		fmt.Print("Are you sure (y/n)")
		fmt.Scan(&areYouSure)
	}

	maxIteration = imageSize / imageSharpness
	step = valueRange / (float64(imageSize) - 1)

	for i := 0; i < len(workerList); i++ {
		go workerHandler(workerList[i])
		fmt.Println("Process spawned for Worker" + string(i))
	}

}

func workerHandler(worker net.Conn) {
	worker.Write(Float64bytes(float64(imageSize)))

}

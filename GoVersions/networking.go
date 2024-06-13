package main

import (
	"fmt"
	"math"
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
var imageSize_local uint64

var payload []byte

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
func waitingRead(connection net.Conn, timeout int) ([]byte, error) {
	ch := make(chan []byte)
	eCh := make(chan error)

	go readContinuous(ch, eCh, connection)

	ticker := time.Tick(time.Second * time.Duration(timeout))
	for {
		select {
		case data := <-ch:
			return data, nil
		case err := <-eCh:
			fmt.Println(err)
		case <-ticker:
			return nil, net.ErrClosed
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
		} else if input == "render" {
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

	payload, err = waitingRead(conn, 1800)
	errHandler(err)
	fmt.Println(payload)

	payload, err = waitingRead(conn, 3)
	errHandler(err)
	imageSize_local = uint64frombytes(payload)
	fmt.Println(imageSize_local)
	fmt.Println("---------")

	payload, err = waitingRead(conn, 3)
	errHandler(err)
	startingX_local = Float64frombytes(payload)
	fmt.Println(startingX_local)
	fmt.Println("---------")

	payload, err = waitingRead(conn, 3)
	errHandler(err)
	startingY_local = Float64frombytes(payload)
	fmt.Println(startingY_local)
	fmt.Println("---------")

	payload, err = waitingRead(conn, 3)
	errHandler(err)
	step_local = Float64frombytes(payload)
	fmt.Println(step_local)
	fmt.Println("---------")

	payload, err = waitingRead(conn, 3)
	// row := uint64frombytes(payload)
	errHandler(err)
	fmt.Println("---------")
}

func networkRender() {
	fmt.Println("")
	fmt.Print("image size [Multiple of 500]: ")
	fmt.Scan(&imageSize)

	// Makes imageSize a multiple of 500
	varTemp := imageSize / 500
	imageSize = int((math.Round(float64(varTemp))) * 500)

	areYouSure := "n"
	for areYouSure != "y" {
		fmt.Print("Allowed memory usage (MB): ")
		fmt.Scan(&memCapacity)
		fmt.Println("Using " + fmt.Sprint(memCapacity) + " MB of memory")
		fmt.Println("[ATTENTION] Using more ram than you device has will crash it.")
		fmt.Print("Are you sure (y/n) ")
		fmt.Scan(&areYouSure)
	}

	maxIteration = imageSize / imageSharpness
	step = valueRange / (float64(imageSize) - 1)

	rowList := make([]bool, imageSize)
	for i := 0; i < len(workerList); i++ {
		go workerHandler(workerList[i], rowList)
		fmt.Println("Process spawned for Worker" + fmt.Sprint(i))
	}

}

func workerHandler(worker net.Conn, rowList []bool) {
	worker.Write(key)

	worker.Write(uint64bytes(uint64(imageSize)))
	time.Sleep(100 * time.Millisecond)
	worker.Write(Float64bytes(startingX))
	time.Sleep(100 * time.Millisecond)
	worker.Write(Float64bytes(startingY))
	time.Sleep(100 * time.Millisecond)
	worker.Write(Float64bytes(step))
	time.Sleep(100 * time.Millisecond)

	rowToAssign := 0
	for i := 0; rowList[i]; i++ {
		rowToAssign = i + 1
	}
	rowList[rowToAssign] = true
	fmt.Println("8===========D")
	worker.Write(uint64bytes(uint64(rowToAssign)))
	time.Sleep(100 * time.Millisecond)
	worker.Write(uint64bytes(uint64(rowToAssign)))
	time.Sleep(100 * time.Millisecond)
	worker.Write(uint64bytes(uint64(rowToAssign)))
	time.Sleep(100 * time.Millisecond)
}

package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

func Float64frombytes(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes
}

func uint64frombytes(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}

func uint64bytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, value)
	return bytes
}

func runtimeBuster() {
	fmt.Println("8=========D")
	time.Sleep(1 * time.Second)
}

// func encodingTest() {
// 	file, err := os.Create("escapeTable/test.bin")
// 	defer file.Close()
// 	errHandler(err)

// 	bs := make([]byte, 8)
// 	for i := 0; i < 6; i++ {
// 		randNum := rand.Uint64()
// 		binary.BigEndian.PutUint64(bs, randNum)
// 		fmt.Println(randNum)
// 		_, err = file.Write(bs)
// 		errHandler(err)
// 	}

// 	file.Close()
// 	fmt.Println("-----------------------------------")

// 	file, err = os.Open("escapeTable/test.bin")
// 	defer file.Close()
// 	errHandler(err)
// 	for i := 0; i < 6; i++ {
// 		_, err = file.Read(bs)
// 		errHandler(err)
// 		fmt.Println(binary.BigEndian.Uint64(bs))
// 	}
// }

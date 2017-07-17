package main

import (
	"encoding/binary"
)

func main() {
	ascBytes := []byte("34")

	data_int16 := binary.LittleEndian.Uint16(ascBytes)

	println(data_int16)

	var revBytes [2]byte

	binary.LittleEndian.PutUint16(revBytes[:], data_int16)

	println(string(revBytes[:]))

	var value uint16 = 32769

	println(float64(int16(value)) / 100)
}

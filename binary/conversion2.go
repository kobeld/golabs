package main

import (
	"encoding/binary"
)

func main() {

	var (
		num      = uint16(45823)
		revBytes [2]byte
	)

	binary.BigEndian.PutUint16(revBytes[:], num)

	// binary.BigEndian.Uint16(revBytes[:])

	println(binary.BigEndian.Uint16(revBytes[:]))

	// binary.LittleEndian.

	newBytes := []byte{revBytes[0], revBytes[1]}

	result := binary.LittleEndian.Uint16(newBytes)

	println(int16(result))

	// ascBytes := []byte("34")

	// data_int16 := binary.LittleEndian.Uint16(ascBytes)

	// println(data_int16)

	// var revBytes [2]byte

	// binary.LittleEndian.PutUint16(revBytes[:], data_int16)

	// println(string(revBytes[:]))

	// var value uint16 = 32769

	// println(float64(int16(value)) / 100)
}

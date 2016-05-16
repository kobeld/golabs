package main

import (
	"github.com/jmcvetta/randutil"
)

func main() {
	str, _ := randutil.AlphaString(10)
	println(str)

	str, _ = randutil.String(4, randutil.Numerals)
	println(str)

}

/*
  Applying the << operator and the use of iota in constants,
  the following type definition neatly defines memory constants
*/
type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

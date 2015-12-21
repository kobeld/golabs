package main

import "github.com/kobeld/goutils"

func main() {
	var (
		milliseconds = []int64{1446174176682, 1446174181681, 1446174186681}
	)

	for _, millisecond := range milliseconds {
		println(goutils.MillisecondToTime(millisecond).String())
	}
}

package main

import (
	"time"
)

const TIME_FORMAT = "01-02"

func main() {
	weekDay := (int)(time.Now().Weekday())

	weekStartTime := time.Now().AddDate(0, 0, -weekDay)

	y, m, d := weekStartTime.Date()

	startTime := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	println(startTime.String())
}

package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat  float64
	Lon  float64
	Date time.Time
}

type Training struct {
	Bootcamp
	Type string
}

func main() {
	go_bootcamp := Training{}
	go_bootcamp.Lat = 34.012836
	go_bootcamp.Lon = -118.495338
	go_bootcamp.Date = time.Now()
	go_bootcamp.Type = "Go"
	fmt.Printf("%+v", go_bootcamp)
}

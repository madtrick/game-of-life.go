package main

import (
	"flag"
	"fmt"
	"time"
)

func ticker(sleep time.Duration, ticks chan<- bool) {
	time.Sleep(sleep)
	ticks <- true
	ticker(sleep, ticks)
}

func main() {
	var numberOfInitialCells int
	var matrix Matrix

	var cols *int
	var rows *int
	var sleep *time.Duration

	cols = flag.Int("rows", 10, "Number of rows")
	rows = flag.Int("cols", 10, "Number of cols")
	sleep = flag.Duration("sleep", 5*time.Second, "Sleep between rounds (in seconds)")
	flag.Parse()

	numberOfInitialCells = int(float32(*cols*(*rows)) * 0.2)

	matrix = Matrix{}
	matrix.Cols = *cols
	matrix.Rows = *rows
	matrix.Population = numberOfInitialCells

	matrix.Initialize()

	var ticks chan bool
	ticks = make(chan bool)

	var done chan bool
	done = make(chan bool)

	go matrix.Update(ticks, done)
	go ticker(*sleep, ticks)

	<-done

	fmt.Println("Done")
}

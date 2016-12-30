package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	Data       [][]bool
	Cols       int
	Rows       int
	Population int
}

func (matrix *Matrix) Initialize() {
	var randomizer *rand.Rand

	randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

	var positionX int
	var positionY int

	for i := 1; i <= matrix.Cols; i++ {
		matrix.Data = append(matrix.Data, make([]bool, matrix.Rows))
	}

	for i := 1; i <= matrix.Population; i++ {
		positionX = randomizer.Intn(matrix.Cols)
		positionY = randomizer.Intn(matrix.Rows)

		fmt.Printf("Initializing cell in %d, %d\n", positionX, positionY)
		matrix.Data[positionX][positionY] = true
	}
}

func (matrix *Matrix) Update(ticks <-chan bool, done chan<- bool) {
	var neighbours int

	for range ticks {
		for j := 0; j < matrix.Cols; j++ {
			var line bytes.Buffer

			for i := 0; i < matrix.Rows; i++ {
				if matrix.Data[i][j] {
					line.WriteString("*")
				} else {
					line.WriteString("_")
				}

				neighbours = numberOfNeighbours(matrix.Data, i, j)

				if matrix.Data[i][j] {
					if neighbours == 0 {
						matrix.Data[i][j] = false
						matrix.Population = matrix.Population - 1
					}

					if neighbours >= 4 {
						matrix.Data[i][j] = false
						matrix.Population = matrix.Population - 1
					}
				}

				if !matrix.Data[i][j] {
					if neighbours == 3 {
						matrix.Data[i][j] = true
						matrix.Population = matrix.Population + 1
					}
				}
			}

			fmt.Println(line.String())
		}

		fmt.Println("")

		if matrix.Population <= 0 {
			done <- true
		}
	}
}

func hasNeighboursAt(matrix [][]bool, i, j, offsetI, offsetJ int) bool {
	return matrix[i+offsetI][j+offsetJ]
}

func numberOfNeighbours(matrix [][]bool, i, j int) int {
	var rows int
	var cols int
	var count int

	cols = len(matrix)
	rows = len(matrix[i])
	count = 0

	for row := -1; row <= 1; row++ {
		for col := -1; col <= 1; col++ {
			offsetI := col
			offsetJ := row

			if i+col < 0 || i+col >= cols {
				continue
			}

			if j+row < 0 || j+row >= rows {
				continue
			}

			if offsetI == 0 && offsetJ == 0 {
				continue
			}

			if hasNeighboursAt(matrix, i, j, offsetI, offsetJ) {
				count = count + 1
			}
		}
	}

	return count
}

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

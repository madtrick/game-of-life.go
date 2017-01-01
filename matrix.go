package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	Data       [][]Cell
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
		matrix.Data = append(matrix.Data, make([]Cell, matrix.Rows))
	}

	for row := 0; row < matrix.Rows; row++ {
		for col := 0; col < matrix.Cols; col++ {
			matrix.Data[row][col] = new(NullCell)
		}
	}

	for i := 1; i <= matrix.Population; i++ {
		positionX = randomizer.Intn(matrix.Cols)
		positionY = randomizer.Intn(matrix.Rows)

		fmt.Printf("Initializing cell in %d, %d\n", positionX, positionY)
		matrix.Data[positionX][positionY] = new(AliveCell)
	}
}

func (matrix *Matrix) Update(ticks <-chan bool, done chan<- bool) {
	var neighbours int

	for range ticks {
		matrix.Population = 0

		for j := 0; j < matrix.Cols; j++ {
			var line bytes.Buffer

			for i := 0; i < matrix.Rows; i++ {
				var cell = matrix.Data[i][j]
				if cell.IsAlive() {
					line.WriteString("*")
				} else {
					line.WriteString("_")
				}

				neighbours = numberOfNeighbours(matrix.Data, i, j)

				matrix.Data[i][j] = cell.Evolve(neighbours)
				if matrix.Data[i][j].IsAlive() {
					matrix.Population += 1
				}
			}

			fmt.Println(line.String())
		}

		fmt.Println("Population ", matrix.Population)
		fmt.Println("")

		if matrix.Population == 0 {
			done <- true
		}
	}
}

func hasNeighboursAt(matrix [][]Cell, i, j, offsetI, offsetJ int) bool {
	return matrix[i+offsetI][j+offsetJ].IsAlive()
}

func numberOfNeighbours(matrix [][]Cell, i, j int) int {
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

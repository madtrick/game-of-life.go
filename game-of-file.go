package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const rows = 10
const cols = 10

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

func main() {
	var matrix [][]bool
	var randomizer *rand.Rand
	var numberOfInitialCells int
	var population int

	numberOfInitialCells = 10
	population = numberOfInitialCells
	randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

	var positionX int
	var positionY int

	for i := 1; i <= cols; i++ {
		matrix = append(matrix, make([]bool, rows))
	}

	for i := 1; i <= numberOfInitialCells; i++ {
		positionX = randomizer.Intn(cols)
		positionY = randomizer.Intn(rows)

		fmt.Printf("Initializing cell in %d, %d\n", positionX, positionY)
		matrix[positionX][positionY] = true
	}

	var neighbours int

	for population > 0 {
		for j := 0; j < cols; j++ {
			var line bytes.Buffer

			for i := 0; i < rows; i++ {
				if matrix[i][j] {
					line.WriteString("*")
				} else {
					line.WriteString("_")
				}

				neighbours = numberOfNeighbours(matrix, i, j)

				if matrix[i][j] {
					if neighbours == 0 {
						matrix[i][j] = false
						population = population - 1
					}

					if neighbours >= 4 {
						matrix[i][j] = false
						population = population - 1
					}
				}

				if !matrix[i][j] {
					if neighbours == 3 {
						matrix[i][j] = true
						population = population + 1
					}
				}
			}

			fmt.Println(line.String())
		}

		fmt.Println("")
		time.Sleep(5 * time.Second)
	}
}

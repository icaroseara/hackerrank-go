package main

import (
	"fmt"
	"log"
	"math"
)

// Matrix is an unexported type that
// represents a n-dimensional array
type Matrix struct {
	dimension int
	elems     [][]int
}

// NewMatrix creates and returns objects of
// the unexported type Matrix.
func NewMatrix(dimension int) *Matrix {
	matrix := new(Matrix)

	matrix.dimension = dimension

	matrix.elems = make([][]int, matrix.dimension)

	for i := range matrix.elems {
		matrix.elems[i] = make([]int, matrix.dimension)
	}

	return matrix
}

// At returns the element in the ith row and jth column
func (m *Matrix) At(x, y int) int {
	return m.elems[y][x]
}

func (m *Matrix) readData() error {
	for y := 0; y < m.dimension; y++ {
		for x := 0; x < m.dimension; x++ {
			_, err := fmt.Scanf("%d", &m.elems[y][x])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Matrix) hourglassSum(x, y int) int {
	sum := 0
	sum += m.At(x, y)
	sum += m.At(x+1, y)
	sum += m.At(x+2, y)
	sum += m.At(x+1, y+1)
	sum += m.At(x, y+2)
	sum += m.At(x+1, y+2)
	sum += m.At(x+2, y+2)
	return sum
}

func (m *Matrix) maxHourglass() int {
	max := math.MinInt64

	for y := 0; y < m.dimension-2; y++ {
		for x := 0; x < m.dimension-2; x++ {
			sum := m.hourglassSum(x, y)
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	matrix := NewMatrix(6)

	err := matrix.readData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(matrix.maxHourglass())
}

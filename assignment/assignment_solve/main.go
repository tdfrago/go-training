package main

import (
	"fmt"
	"math"
)

func ComputeDeterminant(m [][]float64) float64 {
	minor_m := m[1:]
	total := 0.0
	for i := 0; i < len(m); i++ {

		m2x2 := [][]float64{}

		for j := 0; j < len(minor_m); j++ {
			row := []float64{}
			row = append(row, minor_m[j][0:i]...)
			row = append(row, minor_m[j][i+1:]...)
			m2x2 = append(m2x2, row)
		}

		sign := math.Pow(-1, float64(i))
		det_2x2 := (m2x2[0][0] * m2x2[1][1]) - (m2x2[0][1] * m2x2[1][0])
		total += sign * m[0][i] * det_2x2
	}
	return total
}
func GetSolution(m, mX, mY, mZ [][]float64) {
	detm := ComputeDeterminant(m)
	detmX := ComputeDeterminant(mX)
	X := detmX / detm
	detmY := ComputeDeterminant(mY)
	Y := detmY / detm
	detmZ := ComputeDeterminant(mZ)
	Z := detmZ / detm
	fmt.Println(X, Y, Z)
	fmt.Println(math.IsNaN(X), X == Y, Y == Z)
	if detm == 0 {
		if math.IsNaN(X) && math.IsNaN(Y) && math.IsNaN(Z) {
			fmt.Println("dependent - with multiple solutions")
		} else {
			fmt.Println("inconsistent - no solution")
		}
	} else {
		fmt.Printf("x = %.2v, y = %.2v, z = %.2v\n", X, Y, Z)
	}
}

func main() {
	m := [][]float64{{1, -1, 1}, {3, 2, -12}, {4, 1, -11}}
	mX := [][]float64{{7, -1, 1}, {11, 2, -12}, {18, 1, -11}}
	mY := [][]float64{{1, 7, 1}, {3, 11, -12}, {4, 18, -11}}
	mZ := [][]float64{{1, -1, 7}, {3, 2, 11}, {4, 1, 18}}
	GetSolution(m, mX, mY, mZ)
}

//no solution -inf, -inf +inf
//infinite solution all NaN NaN NaN

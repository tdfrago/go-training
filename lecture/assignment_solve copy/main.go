package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {
	var m = mat.NewDense(3, 3, []float64{
		-3, 5, 2,
		5, -1, 4,
		4, -2, 2,
	})

	var v = []float64{-19, -5, 2}

	x := make([]float64, len(v))
	b := make([]float64, len(v))
	d := math.Round(mat.Det(m)*100) / 100
	for c := range v {
		mat.Col(b, c, m)
		m.SetCol(c, v)
		x[c] = math.Round((mat.Det(m)/d)*100) / 100
		m.SetCol(c, b)
	}
	fmt.Println(x)
	for _, val := range x {
		if math.IsNaN(val) {
			fmt.Println("infinite soln")
		}
	}
}

//no solution -inf, -inf +inf
//infinite solution -Inf -inf NaN

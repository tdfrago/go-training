package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/solve" {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		if v, ok := r.Form["coef"]; ok {
			var a, b, c, d, e, f, g, h, i, j, k, l float64
			if n, _ := fmt.Sscanf(v[0], "%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v", &a, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k, &l); n == 12 {
				system, result := SolveEquation(a, b, c, d, e, f, g, h, i, j, k, l)
				fmt.Fprintf(w, "%v\n%v\n", system, result)
			} else {
				fmt.Fprintln(w, "incorrect number of coefficients must be 12")
			}
		} else {
			fmt.Fprintln(w, "incorrect parameter must have 'coef'")
		}
	} else {
		fmt.Fprintln(w, "incorrect url must be '/solve'")
	}

}

func SolveEquation(a, b, c, d, e, f, g, h, i, j, k, l float64) (string, string) {
	m := [][]float64{{a, b, c}, {e, f, g}, {i, j, k}}
	mX := [][]float64{{d, b, c}, {h, f, g}, {l, j, k}}
	mY := [][]float64{{a, d, c}, {e, h, g}, {i, l, k}}
	mZ := [][]float64{{a, b, d}, {e, f, h}, {i, j, l}}
	system := fmt.Sprintf("system:\n%vx + %vy + %vz = %v\n%vx + %vy + %vz = %v\n%vx + %vy + %vz = %v\n", a, b, c, d, e, f, g, h, i, j, k, l)
	result := GetSolution(m, mX, mY, mZ)
	return system, result
}

func GetSolution(m, mX, mY, mZ [][]float64) string {
	detm := ComputeDeterminant(m)
	detmX := ComputeDeterminant(mX)
	X := detmX / detm
	detmY := ComputeDeterminant(mY)
	Y := detmY / detm
	detmZ := ComputeDeterminant(mZ)
	Z := detmZ / detm
	result := ""
	if detm == 0 {
		if math.IsNaN(X) && math.IsNaN(Y) && math.IsNaN(Z) {
			result = "dependent - with multiple solutions"
		} else {
			result = "inconsistent - no solution"
		}
	} else {
		result = fmt.Sprintf("solution:\nx = %.2f, y = %.2f, z = %.2f\n", X, Y, Z)
	}
	return result
}

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

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

//handler receives and processes http requests from client
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

//SolveEquation provides the system of equations and solution results.
func SolveEquation(a, b, c, d, e, f, g, h, i, j, k, l float64) (string, string) {
	m := [][]float64{{a, b, c}, {e, f, g}, {i, j, k}}  //coefficient matrix
	mX := [][]float64{{d, b, c}, {h, f, g}, {l, j, k}} //matrix X
	mY := [][]float64{{a, d, c}, {e, h, g}, {i, l, k}} //matrix Y
	mZ := [][]float64{{a, b, d}, {e, f, h}, {i, j, l}} //matrix Z
	system := fmt.Sprintf("system:\n%vx + %vy + %vz = %v\n%vx + %vy + %vz = %v\n%vx + %vy + %vz = %v\n", a, b, c, d, e, f, g, h, i, j, k, l)
	result := GetSolution(m, mX, mY, mZ)
	return system, result
}

//GetSolution calculates the solution for the systems of equation by Cramer's rule.
func GetSolution(m, mX, mY, mZ [][]float64) string {
	detm := ComputeDeterminant(m)   //determinant of the coefficient matrix
	detmX := ComputeDeterminant(mX) //determinant of Matrix X
	X := detmX / detm               //solution for x
	detmY := ComputeDeterminant(mY) //determinant of Matrix Y
	Y := detmY / detm               //solution for y
	detmZ := ComputeDeterminant(mZ) //determinant of Matrix Z
	Z := detmZ / detm               //solution for z
	result := ""
	if detm == 0 { //If the determinant of the coefficient matrix is 0 then solution is either infinite or not solution at all.
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

//ComputeDeterminant calculates the determinant of the 3x3 matrix
func ComputeDeterminant(m [][]float64) float64 {
	minor_m := m[1:] //sub matrix with no first row
	total := 0.0     //determinant of 3x3 matrix
	for i := 0; i < len(m); i++ {
		m2x2 := [][]float64{} //sub matrix == 2x2 matrix
		for j := 0; j < len(minor_m); j++ {
			row := []float64{}                     //row of the 2x2 matrix
			row = append(row, minor_m[j][0:i]...)  //remove the focus column by getting the element on the left
			row = append(row, minor_m[j][i+1:]...) //remove the focus column by getting the element on the right
			m2x2 = append(m2x2, row)
		}
		sign := math.Pow(-1, float64(i))                                 //alternating signs for the determinant D = (+1)A*m1 (-1)B*m2 (+1)C*m3
		det_2x2 := (m2x2[0][0] * m2x2[1][1]) - (m2x2[0][1] * m2x2[1][0]) //determinant for sub matrix == 2x2 matrix
		total += sign * m[0][i] * det_2x2                                //total determinant for 3x3 matrix
	}
	return total
}

//Is This The Easiest Problem
package main

import "fmt"

func main() {
	var t, a, b, c int
	var s string
	_, err := fmt.Scan(&t)
	if err == nil && (t >= 1 && t <= 20) {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&a, &b, &c)
			if err == nil {
				if (a+b <= c) || (b+c <= a) || (c+a <= b) {
					s = "Invalid"
				} else if (a == b) && (a == c) {
					s = "Equilateral"
				} else if (a == b) || (b == c) || (c == a) {
					s = "Isosceles"
				} else {
					s = "Scalene"
				}
				fmt.Printf("Case %v: %v\n", i, s)
			}
		}
	}
}

//for cmd: go run main.go < 11479.in > 11479.out
// for powershell: type .\11479.in | go run main.go > 11479.out
//refer to https://www.udebug.com/ to verify inputs

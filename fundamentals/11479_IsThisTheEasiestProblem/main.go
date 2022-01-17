//Is This The Easiest Problem
package main

import "fmt"

func main() {
	var t, a, b, c int
	_, err := fmt.Scan(&t)
	if err == nil && (t >= 1 && t <= 20) {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&a, &b, &c)
			if err == nil {
				if (a+b <= c) || (b+c <= a) || (c+a <= b) {
					fmt.Printf("Case %v: Invalid\n", i)
				} else if (a == b) && (a == c) {
					fmt.Printf("Case %v: Equilateral\n", i)
				} else if (a == b) || (b == c) || (c == a) {
					fmt.Printf("Case %v: Isosceles\n", i)
				} else {
					fmt.Printf("Case %v: Scalene\n", i)
				}
			}
		}
	}
}

//for cmd: go run main.go < 11479.in > 11479.out
// for powershell: type .\11479.in | go run main.go > 11479.out
//refer to https://www.udebug.com/ to verify inputs

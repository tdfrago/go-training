//Cost Cutting
package main

import "fmt"

func main() {
	var t, a, b, c, middle_value int
	_, err := fmt.Scan(&t)
	if err == nil && (t >= 1 && t < 20) {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&a, &b, &c)
			if err == nil && (a >= 1000 && a <= 10000) && (b >= 1000 && b <= 10000) && (c >= 1000 && c <= 10000) {
				if (a >= b && a <= c) || (a >= c && a <= b) {
					middle_value = a
				} else if (b >= a && b <= c) || (b >= c && b <= a) {
					middle_value = b
				} else {
					middle_value = c
				}
				fmt.Printf("Case %v: %v\n", i, middle_value)
			}
		}
	}
}

//for cmd: go run main.go < 11727.in > 11727.out
// for powershell: type .\11727.in | go run main.go > 11727.out
//refer to https://www.udebug.com/ to verify inputs

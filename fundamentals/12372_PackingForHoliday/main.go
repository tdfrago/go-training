//Packing for Holiday
package main

import "fmt"

func main() {
	var t, l, w, h int
	_, err := fmt.Scan(&t)
	if err == nil && (t >= 1 && t <= 100) {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&l, &w, &h)
			if err == nil && (l >= 1 && l <= 50) && (w >= 1 && w <= 50) && (h >= 1 && h <= 50) {
				if l <= 20 && w <= 20 && h <= 20 {
					fmt.Printf("Case %v: good\n", i)
				} else {
					fmt.Printf("Case %v: bad\n", i)
				}
			}
		}

	}
}

//for cmd: go run main.go < 12372.in > 12372.out
// for powershell: type .\12372.in | go run main.go > 12372.out
//refer to https://www.udebug.com/ to verify inputs

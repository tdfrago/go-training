//Division of Nlogonia
package main

import "fmt"

func main() {
	var k, n, m, x, y int
	_, err := fmt.Scan(&k) //asks for number of coordinates
	for err == nil && (k >= 0 && k <= 1000) {
		if k == 0 {
			break
		}
		_, err := fmt.Scan(&n, &m) //asks for origin
		if err == nil && (m >= -10000 && m <= 10000) && (n >= -10000 && n <= 10000) {
			for i := 1; i <= k; i++ {
				_, err := fmt.Scan(&x, &y) //asks for coordinates
				if err == nil && (x >= -10000 && x <= 10000) && (y >= -10000 && y <= 10000) {
					ew := x - n //checks if east or west
					ns := y - m //checks if north or south
					if ns == 0 || ew == 0 {
						fmt.Println("divisa")
					} else if ns > 0 && ew < 0 {
						fmt.Println("NO")
					} else if ns > 0 && ew > 0 {
						fmt.Println("NE")
					} else if ns < 0 && ew > 0 {
						fmt.Println("SE")
					} else {
						fmt.Println("SO")
					}
				}
			}
		}
		_, err = fmt.Scan(&k)
	}
}

//for cmd: go run main.go < 11498.in > 11498.out
// for powershell: type .\11498.in | go run main.go > 11498.out
//refer to https://www.udebug.com/ to verify inputs

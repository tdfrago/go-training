//Egypt
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	for err == nil && (a >= 0 && a < 30000) && (b >= 0 && b < 30000) && (c >= 0 && c < 30000) {
		if a == 0 && a == b && a == c {
			break
		} else {
			x := []int{a, b, c}
			sort.Ints(x)
			if x[0]*x[0]+x[1]*x[1] == x[2]*x[2] {
				fmt.Println("right")
			} else {
				fmt.Println("wrong")
			}
		}
		_, err = fmt.Scan(&a, &b, &c)
	}
}

//for cmd: go run main.go < 11854.in > 11854.out
// for powershell: type .\11854.in | go run main.go > 11854.out
//refer to https://www.udebug.com/ to verify inputs

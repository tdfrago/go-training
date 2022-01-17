//Relational Operators
package main

import "fmt"

func main() {
	var t, a, b int
	_, err := fmt.Scan(&t)
	if err == nil && (t >= 1 && t < 15) {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&a, &b)
			if err == nil && (a < 1000000001) && (b < 1000000001) {
				if a == b {
					fmt.Println("=")
				} else if a < b {
					fmt.Println("<")
				} else {
					fmt.Println(">")
				}
			}
		}
	}
}

//for cmd: go run main.go < 11172.in > 11172.out
// for powershell: type .\11172.in | go run main.go > 11172.out
//refer to https://www.udebug.com/ to verify inputs

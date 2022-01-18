//Odd Sum
package main

import "fmt"

func main() {
	var t, a, b, sum int   //variable declaration
	_, err := fmt.Scan(&t) //input
	if err == nil && (t >= 1 && t <= 100) {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&a)
			if err == nil && (a >= 0 && a <= 100) {
				_, err := fmt.Scan(&b)
				if err == nil && (b >= 0 && b <= 100) {
					sum = 0
					if b < a {
						fmt.Printf("Case %v: %v\n", i, sum)
					} else {
						for a <= b { //loop from a to b, add 1 to a until it reaches b
							if a%2 != 0 { //check if odd number then adds odd number to the sum
								sum = sum + a
							}
							a = a + 1
						}
						fmt.Printf("Case %v: %v\n", i, sum) //print formating
					}
				}
			}
		}
	}

}

//for cmd: go run main.go < 10783.in > 10783.out
// for powershell: type .\10783.in | go run main.go > 10783.out
//refer to https://www.udebug.com/ to verify inputs

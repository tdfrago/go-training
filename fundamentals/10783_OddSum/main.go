//Odd Sum
package main

import "fmt"

func main() {
	var t, a, b, sum int
	_, err := fmt.Scan(&t)
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
						for a <= b {
							if a%2 != 0 {
								sum = sum + a
							}
							a = a + 1
						}
						fmt.Printf("Case %v: %v\n", i, sum)
					}
				}
			}
		}
	}

}

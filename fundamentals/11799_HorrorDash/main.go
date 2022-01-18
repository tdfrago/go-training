//Horror Dash
package main

import "fmt"

func main() {
	var t, n, max, speed int
	_, err := fmt.Scan(&t)
	if err == nil && (t <= 50) {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&n)
			max = 0
			if err == nil {
				for j := 1; j <= n; j++ {
					_, err := fmt.Scan(&speed)
					if err == nil && (speed >= 1 && speed <= 10000) {
						if speed > max {
							max = speed
						}
					}
				}
			}
			fmt.Printf("Case %v: %v\n", i, max)
		}
	}
}

//for cmd: go run main.go < 11799.in > 11799.out
// for powershell: type .\11799.in | go run main.go > 11799.out
//refer to https://www.udebug.com/ to verify inputs

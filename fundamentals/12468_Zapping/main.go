//Zapping
package main

import "fmt"

func main() {
	var a, b, up, down int
	_, err := fmt.Scan(&a, &b)
	for err == nil {
		if a == -1 && b == -1 {
			break
		} else {
			if a < b {
				up = b - a
				down = 100 - (b - a)
			} else {
				up = a - b
				down = 100 - (a - b)
			}
		}
		if up < down {
			fmt.Println(up)
		} else {
			fmt.Println(down)
		}
		_, err = fmt.Scan(&a, &b)
	}
}

//for cmd: go run main.go < 12468.in > 12468.out
// for powershell: type .\12468.in | go run main.go > 12468.out
//refer to https://www.udebug.com/ to verify inputs

//Hashmat The Brave Warrior
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, diff int
	c := math.Pow(2, 32)
	_, err := fmt.Scan(&a, &b)
	for err == nil && (a >= 0 && a <= int(c)) && (b >= 0 && b <= int(c)) {
		diff = a - b
		result := math.Abs(float64(diff))
		fmt.Println(int(result))
		_, err = fmt.Scan(&a, &b)
	}
}

//for cmd: go run main.go < 10055.in > 10055.out
// for powershell: type .\10055.in | go run main.go > 10055.out
//refer to https://www.udebug.com/ to verify inputs

//Ecological Premium
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, f, farm_size, num_animals, envi_friendly int
	_, err := fmt.Scan(&n) //get number of tests
	if err == nil && (n > 0 && n < 20) {
		for i := 1; i <= n; i++ {
			_, err := fmt.Scan(&f) //get number of farmers
			if err == nil && (f > 0 && f < 20) {
				sum := 0.0
				for j := 1; j <= f; j++ {
					_, err := fmt.Scan(&farm_size, &num_animals, &envi_friendly) //get farm size, number of animals, and degree of environment friendliness
					if err == nil && (farm_size >= 0 && farm_size <= 100000) && (num_animals >= 0 && num_animals <= 100000) && (envi_friendly >= 0 && envi_friendly <= 100000) {
						a := float64(farm_size) //assign int vars to float
						b := float64(num_animals)
						c := float64(envi_friendly)
						sum += (a / b) * c * b //compute premium and summation of farmers' premiums
					}
				}
				result := int(math.Round(sum))
				fmt.Println(result)
			}
		}
	}
}

//for cmd: go run main.go < 10300.in > 10300.out
// for powershell: type .\10300.in | go run main.go > 10300.out
//refer to https://www.udebug.com/ to verify inputs

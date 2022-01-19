//Search For Nessy
//minimum requirement for 1 sonar is 3x3
//so atleast 3 blocks for a row and 3 blocks for column
//get the total number of blocks in a row and divide by 3, get only the whole number
//do the same with the columns
//multiply the results: rows//3 * columns//3
package main

import "fmt"

func main() {
	var t, n, m, sonar int
	_, err := fmt.Scan(&t)
	if err == nil {
		for i := 1; i <= t; i++ {
			_, err := fmt.Scan(&n, &m)
			if err == nil && (n >= 6 && n <= 10000) && (m >= 6 && m <= 10000) {
				sonar = (n / 3) * (m / 3) //minimum requirement for 1 sonar is 3x3
			}
			fmt.Println(sonar)
		}
	}
}

//for cmd: go run main.go < 11044.in > 11044.out
// for powershell: type .\11044.in | go run main.go > 11044.out
//refer to https://www.udebug.com/ to verify inputs

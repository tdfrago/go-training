//What is the Median?
package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers []int
	var x, median int
	_, err := fmt.Scan(&x)
	for err == nil {
		numbers = append(numbers, x)
		sort.Ints(numbers)
		mid := len(numbers) / 2
		if len(numbers)%2 == 0 {
			median = (numbers[mid-1] + numbers[mid]) / 2
		} else {
			median = numbers[mid]
		}
		fmt.Println(median)
		_, err = fmt.Scan(&x)
	}
}

//for cmd: go run main.go < 10107.in > 10107.out
// for powershell: type .\10107.in | go run main.go > 10107.out
//refer to https://www.udebug.com/ to verify inputs

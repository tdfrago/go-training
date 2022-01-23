//Train Swapping
package main

import "fmt"

func main() {
	var n, l, num int
	_, err := fmt.Scan(&n) //input number of cases
	if err == nil {
		for num_cases := 1; num_cases <= n; num_cases++ {
			_, err := fmt.Scan(&l) //input length of train
			carriages := []int{}   //make a list for carriages
			count := 0
			if err == nil {
				for length := 1; length <= l; length++ {
					_, err := fmt.Scan(&num) //input the carriage number
					if err == nil {
						carriages = append(carriages, num) //append carriage to a list
					}
				}
				for i, value_1 := range carriages[:len(carriages)-1] { //apply bubble sort to count the switch
					for _, value_2 := range carriages[i+1:] {
						if value_1 > value_2 { //if current number is bigger than next number, switch/count
							count++
						}
					}
				}
			}
			fmt.Printf("Optimal train swapping takes %v swaps.\n", count)
		}
	}
}

//for cmd: go run main.go < 299.in > 299.out
// for powershell: type .\299.in | go run main.go > 299.out
//refer to https://www.udebug.com/ to verify inputs

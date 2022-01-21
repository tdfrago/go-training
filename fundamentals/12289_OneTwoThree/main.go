//One-Two-Three
package main

import "fmt"

func main() {
	var n int
	var word string
	one := "one"
	fmt.Scan(&n) //asks for number of words
	for i := 1; i <= n; i++ {
		fmt.Scan(&word)
		if len(word) > 3 { //any word exceeding 3 letters is three
			fmt.Println("3")
		} else {
			match := 0
			for i, _ := range one { //match the 3 letter word with one
				if word[i] == one[i] {
					match++
				}
			}
			if match >= 2 { //if match has 2 or more then it's a one
				fmt.Println("1")
			} else { //else it's already two
				fmt.Println("2")
			}
		}
	}
}

//for cmd: go run main.go < 12289.in > 12289.out
// for powershell: type .\12289.in | go run main.go > 12289.out
//refer to https://www.udebug.com/ to verify inputs

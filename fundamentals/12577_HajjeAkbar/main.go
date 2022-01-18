//Hajj-e-Akbar
//try user defined function
package main

import "fmt"

func hajj(word string) (meaning string) {
	if word == "Hajj" {
		meaning = "Hajj-e-Akbar"
	} else if word == "Umrah" {
		meaning = "Hajj-e-Asghar"
	}
	return meaning
}

func main() {
	var word string
	_, err := fmt.Scan(&word)
	count := 0
	for err == nil {
		count++
		if word == "*" {
			break
		}
		fmt.Printf("Case %v: %v\n", count, hajj(word))
		_, err = fmt.Scan(&word)
	}
}

//for cmd: go run main.go < 12577.in > 12577.out
// for powershell: type .\12577.in | go run main.go > 12577.out
//refer to https://www.udebug.com/ to verify inputs

//Back To High School Physics
//formula: v = d/t, d = vt

package main

import "fmt"

func main() {
	var v, t int
	_, err := fmt.Scan(&v, &t) //use _ if you don't want to use it
	//Scan returns 1. Number of variables it got from scanning
	// 2. error value
	for err == nil { //nil == no error
		fmt.Println(2 * v * t)
		_, err = fmt.Scan(&v, &t)
	}

}

//for cmd: go run main.go < 10071.in > 10071.out
// for powershell: type .\10071.in | go run main.go > 10071.out
//refer to https://www.udebug.com/ to verify inputs

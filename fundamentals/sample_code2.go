package main

import "fmt" //import format library

func main() {
	var x, y int     // variable declaration
	fmt.Scan(&x, &y) //standard input "&"" address of operator
	fmt.Println(x, "+", y, "=", x+y)
}

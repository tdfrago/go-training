package main

import "fmt"

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total
}
func main() {
	s := []int{8, 4, -3, 10, 2, 1}
	c := make(chan int)
	go sum(s[:3], c) //8 4 -3 = 9
	go sum(s[3:], c) //10 2 1 = 13
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

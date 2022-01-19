
package main

import "fmt"

func main() {
	var c, n, grade int
	_, err := fmt.Scan(&c) //get number of tests
	if err == nil {        //check for error
		for i := 1; i <= c; i++ { //loop according to number of tests
			_, err := fmt.Scan(&n)                   //get number of students
			if err == nil && (n >= 1 && n <= 1000) { //check for error
				grades := []int{}         //make empty list for grades
				for j := 1; j <= n; j++ { //loop according to number of students
					_, err := fmt.Scan(&grade)                      //get the grades each student
					if err == nil && (grade >= 0 && grade <= 100) { //check for error
						grades = append(grades, grade) //append the grade to lists
					}
				}
				sum := 0
				for _, num := range grades { //for loop to get the sum of students grades
					sum += num
				}
				num_students := float64(len(grades))   //assign int num of students to float num of students
				average := float64(sum) / num_students //get average of grades
				above_average := 0.0                   //number of students above average
				for _, num := range grades {
					if float64(num) > average { //check if grade in list is above the average
						above_average++
					}
				}
				result := (above_average / num_students) * 100 //percentage of above average students
				fmt.Printf("%.3f%%\n", result)
			}
		}
	}
}

//for cmd: go run main.go < 10370.in > 10370.out
// for powershell: type .\10370.in | go run main.go > 10370.out
//refer to https://www.udebug.com/ to verify inputs

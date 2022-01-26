package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ReadFile(filename string) [][]string {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if filepath.Ext(strings.TrimSpace(filename)) != ".csv" { //checks if file is in .csv format
		log.Fatal("Incorrect database format must be in .csv format")
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	lines, err := r.ReadAll() //read csv
	if err != nil {
		log.Fatalf("Failed to parse csv file: %v", err)
	}
	return lines
}

func CheckNumberOfQuestions(num_questions int, lines [][]string) {
	if num_questions > len(lines) { //if requested number of questions exceed number of questions in db return error log
		log.Fatalf("Insufficient Questions: There're only %v questions. User requested %v questions.", len(lines), num_questions)
	}
}

func quiz(lines [][]string, num_questions int) int {
	rand.Seed(time.Now().UnixNano())        //provide a random seed
	random_numbers := rand.Perm(len(lines)) //create a list of random numbers based on number of questions in db
	correct := 0
	for i, item := range random_numbers[:num_questions] { //loop on the list of random numbers but limit to number of questions specified by user
		var ans string
		fmt.Printf("Q%v: %s = ", i+1, lines[item][0]) //show question
		fmt.Scan(&ans)                                //input answer
		strings.TrimSpace(ans)                        //remove leading and trailing whitespaces
		if ans == lines[item][1] {                    //check if answer is correct
			correct++
		}
	}
	return correct
}

func main() {
	dbname := flag.String("csv", "problems.csv", "database")  //flag for csv
	num_questions := flag.Int("n", 10, "number of questions") //flag for number of questions
	flag.Parse()
	lines := ReadFile(*dbname)
	CheckNumberOfQuestions(*num_questions, lines)
	correct := quiz(lines, *num_questions)
	fmt.Printf("You answered %v out of %v questions correctly.", correct, *num_questions) //show result
}

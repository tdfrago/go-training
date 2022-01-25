package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	dbname := flag.String("csv", "problems.csv", "database")  //flag for csv
	num_questions := flag.Int("n", 10, "number of questions") //flag for number of questions
	flag.Parse()
	fp, err := os.Open(*dbname)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fp.Close()
	r := csv.NewReader(fp)
	lines, _ := r.ReadAll() //read csv

	if *num_questions > len(lines) { //if requested number of questions exceed number of questions in db return error log
		log.Fatalf("Insufficient Questions: There're only %v questions. User requested %v questions.", len(lines), *num_questions)
	}

	rand.Seed(time.Now().UnixNano())        //provide a random seed
	random_numbers := rand.Perm(len(lines)) //create a list of random numbers based on number of questions in db
	correct := 0
	for i, item := range random_numbers[:*num_questions] { //loop on the list of random numbers but limit to number of questions specified by user
		var ans string
		fmt.Printf("Q%v: %s = ", i+1, lines[item][0]) //show question
		fmt.Scan(&ans)                                //input answer
		strings.TrimSpace(ans)                        //remove leading and trailing whitespaces
		if ans == lines[item][1] {                    //check if answer is correct
			correct++
		}
	}
	fmt.Printf("You answered %v out of %v questions correctly.", correct, *num_questions) //show result
}

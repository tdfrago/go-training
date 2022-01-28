package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("sample.csv")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fp.Close()
	r := csv.NewReader(fp)
	//r.Comma = "\t"
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(lines)

	fp2, err := os.Create("output.csv")
	w := csv.NewWriter(fp2)
	for _, row := range lines {
		_ = w.Write(row)
	}
	w.Flush()
	fp2.Close()
}

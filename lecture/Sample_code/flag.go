package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fname := flag.String("file", "input.txt", "program database")
	flag.Parse()
	data, err := os.ReadFile(*fname)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%x\n", sha256.Sum256(data))
}

//.\flag.exe -file hello.txt

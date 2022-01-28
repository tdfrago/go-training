package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name   string `json:"name"`
	Job    string `json:"job"`
	Salary []struct {
		Value float64
	}
}

func main() {
	fp, _ := os.Open("sample2.json")
	defer fp.Close()
	byteData, _ := ioutil.ReadAll(fp)
	var u []User
	json.Unmarshal(byteData, &u)
	fmt.Println(u)
}

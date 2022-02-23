package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

type Contact struct {
	DSID               string `json:"DSID"`
	FirstName          string `json:"First Name"`
	FirstNameInitial   string `json:"First Name Initial"`
	MiddleName         string `json:"Middle Name"`
	MiddleNameInitial  string `json:"Middle Name Initial"`
	LastName           string `json:"Last Name"`
	LastNameInitial    string `json:"Last Name Initial"`
	PreferredName      string `json:"Preferred Name"`
	NamePrefix         string `json:"Name Prefix"`
	Company            string `json:"Company"`
	Email              string `json:"Email"`
	Phone              string `json:"Phone"`
	CompanyAddress     string `json:"Company Address"`
	CurrentTimezone    string `json:"Current Timezone"`
	SocialMedia        string `json:"Social Media"`
	Location           string `json:"Location"`
	NameSuffix         string `json:"Name Suffix"`
	YearsofExperience  int    `json:"Years of Experience"`
	TimeSpentatCompany int    `json:"Time Spent at Company"`
	Job                string `json:"Job"`
}

func readJson(filename string) Contact {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened contact.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var contact Contact

	json.Unmarshal(byteValue, &contact)

	return contact
}

func getmin(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func getfuzzyscore(word1, word2 string) float64 {
	var matrix [][]int
	var cost, min int
	var weight float64

	rows := len(word1) + 1
	cols := len(word2) + 1

	for i := 0; i < rows; i++ {
		var list []int
		for j := 0; j < cols; j++ {
			list = append(list, 0)
		}
		matrix = append(matrix, list)
	}

	for i := 1; i < rows; i++ {
		matrix[i][0] = i
	}

	for i := 1; i < cols; i++ {
		matrix[0][i] = i
	}

	for i := 1; i < cols; i++ {
		for j := 1; j < rows; j++ {
			if word1[j-1] == word2[i-1] {
				cost = 0
			} else {
				cost = 1
			}
			values := []int{matrix[j-1][i] + 1,
				matrix[j][i-1] + 1,
				matrix[j-1][i-1] + cost}

			min = getmin(values)
			matrix[j][i] = min
		}
	}

	match := 1 - float64(min)/float64(len(word2))
	if match < 0 {
		match = 0
	}

	weight = 1
	fuzzy_score := (math.Round(match*10) / 10) * weight
	return fuzzy_score
}

func main() {
	contact := readJson("contact.json")
	contact_val := readJson("contact_val.json")
	fmt.Println(contact)
	fmt.Println()
	fmt.Println(contact_val)
	fuzzy_score := compare(contact, contact_val)
	fmt.Println(fuzzy_score)
}

func compare(contact, contact_val Contact) float64 {
	fuzzy_score := getfuzzyscore(contact.DSID, contact_val.DSID)
	return fuzzy_score
}

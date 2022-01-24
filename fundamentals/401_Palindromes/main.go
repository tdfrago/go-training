//Palindromes
package main

import "fmt"

func main() {
	var word string
	//map letters with mirrored characters
	mirror := map[string]string{
		"A": "A", "E": "3", "H": "H", "I": "I", "J": "L",
		"L": "J", "M": "M", "O": "O", "S": "2", "T": "T",
		"U": "U", "V": "V", "W": "W", "X": "X", "Y": "Y",
		"Z": "5", "1": "1", "2": "S", "3": "E", "5": "Z",
		"8": "8",
	}
	_, err := fmt.Scan(&word) //get word
	for err == nil {
		palindrome := true
		mirror_word := true
		//check if word is a palindrome or not
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			if word[i] != word[j] {
				palindrome = false
			}
		}
		//check if word is a mirrored word or not
		for i, j := 0, len(word)-1; i <= j; i, j = i+1, j-1 {
			if mirror[string(word[i])] != string(word[j]) {
				mirror_word = false
			}
		}
		//print the result if palindrome/mirrored string
		if palindrome == true && mirror_word == true {
			fmt.Printf("%v -- is a mirrored palindrome.\n\n", word)
		} else if palindrome == true && mirror_word == false {
			fmt.Printf("%v -- is a regular palindrome.\n\n", word)
		} else if palindrome == false && mirror_word == true {
			fmt.Printf("%v -- is a mirrored string.\n\n", word)
		} else {
			fmt.Printf("%v -- is not a palindrome.\n\n", word)
		}
		_, err = fmt.Scan(&word)
	}
}

//for cmd: go run main.go < 401.in > 401.out
// for powershell: type .\401.in | go run main.go > 401.out
//refer to https://www.udebug.com/ to verify inputs

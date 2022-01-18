//Language Detection
package main

import "fmt"

func main() {
	country := map[string]string{
		"HELLO":        "ENGLISH",
		"HOLA":         "SPANISH",
		"HALLO":        "GERMAN",
		"BONJOUR":      "FRENCH",
		"CIAO":         "ITALIAN",
		"ZDRAVSTVUJTE": "RUSSIAN",
	}
	mesg := ""
	count := 0
	_, err := fmt.Scan(&mesg)
	for err == nil {
		count++
		if mesg == "#" {
			break
		} else {
			if lang, ok := country[mesg]; ok {
				fmt.Printf("Case %v: %v\n", count, lang)
			} else {
				fmt.Printf("Case %v: UNKNOWN\n", count)
			}
		}
		_, err = fmt.Scan(&mesg)
	}
}

//for cmd: go run main.go < 12250.in > 12250.out
// for powershell: type .\12250.in | go run main.go > 12250.out
//refer to https://www.udebug.com/ to verify inputs

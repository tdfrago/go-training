package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	if v, ok := r.Form["compute"]; ok {
		fmt.Fprintf(w, "%s %v\n", v[0], ok)
		var x, y int
		if n, _ := fmt.Sscanf(v[0], "%d+%d", &x, &y); n == 2 {
			fmt.Fprintf(w, "%d * %d = %d\n", x, y, x*y)
		}
	}
}

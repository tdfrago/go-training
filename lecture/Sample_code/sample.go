package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go getData(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func getData(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}
	duration := time.Since(start).Seconds()
	fmt.Printf("%.2fs %v %v\n", duration, len(b), url)
}

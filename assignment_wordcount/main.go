package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

//WordCounter is a struct containing mutex to safely access data across go routines and the mapping a word to its number of occurences
type WordCounter struct {
	mu    sync.Mutex
	count map[string]int
}

//Inc adds a new word as a key to the map or increment by 1 if found an existing word(key) in the map
func (word_count *WordCounter) Inc(key string) {
	word_count.mu.Lock()
	word_count.count[key]++
	word_count.mu.Unlock()
}

//Value returns the total count of the word(key)
func (word_count *WordCounter) Value(key string) int {
	word_count.mu.Lock()
	defer word_count.mu.Unlock()
	return word_count.count[key]
}

//ReadFile reads the file and returns a list of strings to the channel
func ReadFile(filename string, ch chan []string) {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if filepath.Ext(strings.TrimSpace(filename)) != ".txt" { //checks if file is in .csv format
		log.Fatal("Incorrect database format must be in .txt format")
	}
	defer fp.Close()

	file, err := ioutil.ReadAll(fp)

	//remove non alphanumeric characters
	n := 0
	for _, b := range file {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' || b == '\n' {
			file[n] = b
			n++
		}
	}
	list := strings.Fields(strings.ToLower(string(file[:n])))

	ch <- list
}

//Sortkey takes in a map and returns an alphabetize list of keys
func Sortkey(word_count WordCounter) []string {
	key_list := []string{}
	for key := range word_count.count {
		key_list = append(key_list, key)
	}
	sort.Strings(key_list)
	return key_list
}

func main() {
	all_words := []string{}
	word_count := WordCounter{count: make(map[string]int)}
	ch := make(chan []string)
	for _, file := range os.Args[1:] {
		go ReadFile(file, ch) //concurrently reads the files
	}

	for range os.Args[1:] {
		all_words = append(all_words, <-ch...)
	}

	for _, word := range all_words {
		go word_count.Inc(word) //concurrently counts the words
	}

	key_list := Sortkey(word_count)

	for _, key := range key_list {
		fmt.Println(key, word_count.count[key]) //prints the result
	}
}

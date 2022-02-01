package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Record struct {
	ID       int
	Last     string
	First    string
	Company  string
	Address  string
	Country  string
	Position string
}

type Database struct {
	nextID int
	mu     sync.Mutex
	recs   []Record
}

func main() {
	db := &Database{recs: []Record{}}
	http.ListenAndServe(":8080", db.handler())
}

func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/store" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/store/%d", &id); n == 1 {
			db.processID(id, w, r)
		}
	}
}

func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var rec Record
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, record := range db.recs {
			if rec.Last == record.Last && rec.First == record.First && rec.Company == record.Company && rec.Address == record.Address && rec.Country == record.Country && rec.Position == record.Position {
				w.WriteHeader(http.StatusConflict)
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(record); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		}
		db.mu.Lock()
		rec.ID = db.nextID
		db.nextID++
		db.recs = append(db.recs, rec)
		db.mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, http.StatusText(http.StatusCreated))
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(db.recs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "GET":
		for _, record := range db.recs {
			if id == record.ID {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(record); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	case "PUT":

	case "DELETE":
		for j, record := range db.recs {
			if id == record.ID {
				db.mu.Lock()
				db.recs = append(db.recs[:j], db.recs[j+1:]...)
				db.mu.Unlock()
				fmt.Fprintln(w, "Record has been deleted")
				return
			}
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

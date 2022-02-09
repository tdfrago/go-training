package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

//Record struct with ID, LastName, FirstName, Company, Address, Country, and Position fields
type Record struct {
	ID       int
	Last     string
	First    string
	Company  string
	Address  string
	Country  string
	Position string
}

//Database struct containing the Records
type Database struct {
	nextID int
	mu     sync.Mutex
	recs   []Record
}

func main() {
	db := &Database{recs: []Record{}}
	http.ListenAndServe(":8080", db.handler())
}

//handler receives and processes http requests from client
func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
			db.processID(id, w, r)
		} else {
			fmt.Fprintln(w, "incorrect url must be '/contacts'")

		}
	}
}

//process implements POST, GET, PUT, and DELETE Methods
func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST creates new record and sends 201 Created Status
	//if records exists, then sends the old record with 409 Conflict status
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
	//GET retrives all records
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(db.recs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	//PUT method not allowed in process
	//returns 405 Method not allowed status
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//DELETE method not allowed in process
	//returns 405 Method not allowed status
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//processID implements POST, GET, PUT, and DELETE Methods for existing records in database
func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST method not allowed in processID
	//returns 405 Method not allowed status
	case "POST":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//GET retrieves the record of the specified ID
	//if ID not in database, then returns 404 Not found status
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
	//PUT retrieves the record of the specified ID, and updates the record based on the record update
	//if ID not in database, then returns 404 Not found status
	case "PUT":
		for _, record := range db.recs {
			if id == record.ID {
				var rec Record
				if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				db.mu.Lock()
				if rec.Last != "" {
					db.recs[id].Last = rec.Last
				}
				if rec.First != "" {
					db.recs[id].First = rec.First
				}
				if rec.Company != "" {
					db.recs[id].Company = rec.Company
				}
				if rec.Address != "" {
					db.recs[id].Address = rec.Address
				}
				if rec.Country != "" {
					db.recs[id].Country = rec.Country
				}
				if rec.Position != "" {
					db.recs[id].Position = rec.Position
				}
				db.mu.Unlock()
				fmt.Fprintln(w, "Record has been updated")
				return
			}
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	//DELETE deletes the record of the specified ID on the database
	//if ID not in database, then returns 404 Not found status
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

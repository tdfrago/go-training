package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	LastName  string
	FirstName string
	UserName  string
	Password  string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("error validatin sql.open arguments")
		panic(err.Error())
	}
	http.ListenAndServe(":8080", handler())
}

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			//
		case "/signup":
			signup(w, r)
		case "/login":
			//
		case "/logout":
			//
		default:
			fmt.Fprintln(w, "incorrect url'")
		}
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		lastname := r.FormValue("lastname")
		firstname := r.FormValue("firstname")
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("lastname:%v,firstname:%v,username:%v,password:%v\n", lastname, firstname, username, password)

		stmt := "SELECT Id FROM users WHERE username = ?"
		row := db.QueryRow(stmt, username)
		var Id string
		err := row.Scan(&Id)
		if err != sql.ErrNoRows {
			fmt.Println("Username already exists:", err)
			fmt.Fprintln(w, "Username already exists")
			return
		}

		var hash []byte
		hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("hash:", hash)
		fmt.Println("string(hash):", string(hash))

		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO users (LastName, FirstName, UserName, Password) VALUES (?, ?, ?, ?);")
		if err != nil {
			fmt.Println("error statement:", err)
		}
		defer insert_stmt.Close()

		var result sql.Result
		result, err = insert_stmt.Exec(lastname, firstname, username, hash)
		rows_affected, _ := result.RowsAffected()
		last_insertedId, _ := result.LastInsertId()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Print("last inserted id:", last_insertedId)
		fmt.Println("error:", err)
		if err != nil {
			fmt.Println("Error registering new user:", err)
			return
		}

		fmt.Println("User has been created.")
		fmt.Fprintln(w, "User has been successfully created")
	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
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

var store = sessions.NewCookieStore([]byte("super-secret"))

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("error validatin sql.open arguments")
		panic(err.Error())
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/addmovie", addmovie)
	//http.HandleFunc("/updatemovie", updatemovie)
	//http.HandleFunc("/deletemovie",deletemovie)
	//http.HandleFunct("/viewmovies",viewmovies)
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
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
		if err != nil || rows_affected != 1 {
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

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("username: %v, password: %v\n", username, password)

		var Id, hash string
		stmt := "SELECT Id, Password FROM users WHERE UserName =?"
		row := db.QueryRow(stmt, username)
		err := row.Scan(&Id, &hash)
		fmt.Println("hass:", hash)
		if err != nil {
			fmt.Println("error getting hash", err)
			fmt.Fprint(w, "username not found")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			session, _ := store.Get(r, "session")
			session.Values["Id"] = Id
			session.Save(r, w)
			fmt.Println("User has logged in.")
			fmt.Fprint(w, "You have successfully logged in")
			return
		}
		fmt.Print("Incorrect password")
		fmt.Fprint(w, "Incorrect password")

	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	Id, ok := session.Values["Id"]
	fmt.Println("ok:", ok)
	if !ok {
		fmt.Fprint(w, "not logged in")
		return
	}
	stmt := "SELECT FirstName FROM users WHERE Id =?"
	row := db.QueryRow(stmt, Id)
	var firstname string
	err := row.Scan(&firstname)
	fmt.Println("username:", firstname)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("Hi! %v is seeing the home page.\n", firstname)
	fmt.Fprintf(w, "Hi! %v is seeing the homepage", firstname)
}

func logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		session, _ := store.Get(r, "session")
		delete(session.Values, "Id")
		session.Save(r, w)
		fmt.Println("User has logged out")
		fmt.Fprintln(w, "You have logged out")
	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func addmovie(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		session, _ := store.Get(r, "session")
		Id, ok := session.Values["Id"]
		fmt.Println("ok:", ok)
		if !ok {
			fmt.Fprint(w, "not logged in")
			return
		}
		stmt := "SELECT UserName, FirstName FROM users WHERE Id =?"
		row := db.QueryRow(stmt, Id)
		var username, firstname string
		err := row.Scan(&username, &firstname)
		fmt.Println("firstname:", firstname)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		r.ParseForm()
		title := r.FormValue("title")
		genre := r.FormValue("genre")
		year := r.FormValue("year")
		director := r.FormValue("director")
		language := r.FormValue("language")
		country := r.FormValue("country")
		status := r.FormValue("status")

		if title == "" || genre == "" || year == "" || director == "" || language == "" || country == "" || status == "" {
			fmt.Println("Must fill in all fields")
			fmt.Fprintln(w, "Please fill in all fields")
			return
		}

		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO movies (UserName, Title, Genre, Year, Director, Language, Country, Status) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			fmt.Println("error statement:", err)
		}
		defer insert_stmt.Close()

		result, err := insert_stmt.Exec(username, title, genre, year, director, language, country, status)
		rows_affected, _ := result.RowsAffected()
		last_insertedId, _ := result.LastInsertId()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Print("last inserted id:", last_insertedId)
		fmt.Println("error:", err)

		if err != nil || rows_affected != 1 {
			fmt.Println("Error adding a movie:", err)
			return
		}

		fmt.Printf("%v added a movie.\n", firstname)
		fmt.Fprintf(w, "Hi %v! Movie has been added.", firstname)
	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

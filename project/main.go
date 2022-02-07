package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

type UserLogin struct {
	UserName string
	Password string
}

type Movie struct {
	Id       int
	Title    string
	Genre    string
	Year     int
	Director string
	Language string
	Country  string
	Status   string
	UserName string
}

var db *sql.DB

var store = sessions.NewCookieStore([]byte("super-secret"))

func logger(username, message string) {
	date := time.Now().Format("01-02-2006")
	time := time.Now().Format("15:04:05.00")
	fmt.Println(username, message, date, time)
	var insert_stmt *sql.Stmt
	var err error
	insert_stmt, err = db.Prepare("INSERT INTO testdb.logs (Date, Time, UserName, Message) VALUES (?, ?, ?, ?);")
	if err != nil {
		fmt.Println("error statement:", err)
	}
	defer insert_stmt.Close()
	fmt.Println(insert_stmt)
	var result sql.Result
	result, err = insert_stmt.Exec(date, time, username, message)
	rows_affected, _ := result.RowsAffected()
	if err != nil || rows_affected != 1 {
		insert_stmt.Exec(date, time, username, "error log")
	}
}

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
	http.HandleFunc("/movies", movies)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			logger("", "error parsing json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var Id string

		lastname := user.LastName
		firstname := user.FirstName
		username := user.UserName
		password := user.Password

		fmt.Printf("lastname:%v,firstname:%v,username:%v,password:%v\n", lastname, firstname, username, password)

		logger(username, "user signup input")

		stmt := "SELECT Id FROM testdb.users WHERE username = ?"
		row := db.QueryRow(stmt, username)

		err := row.Scan(&Id)
		if err != sql.ErrNoRows {
			fmt.Fprintln(w, "Username already exists")
			logger(username, "username already exists")
			return
		}
		logger(username, "username does not yet exist")

		var hash []byte
		hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			logger(username, "error hashing password")
			return
		}
		fmt.Println("hash:", hash)
		fmt.Println("string(hash):", string(hash))
		logger(username, "password converted to hash")

		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO testdb.users (LastName, FirstName, UserName, Password) VALUES (?, ?, ?, ?);")
		if err != nil {
			logger(username, "error preparing insert statement")
			return
		}
		defer insert_stmt.Close()

		var result sql.Result
		result, err = insert_stmt.Exec(lastname, firstname, username, hash)
		rows_affected, _ := result.RowsAffected()

		if err != nil || rows_affected != 1 {
			logger(username, "error registering new user")
			return
		}
		logger(username, "user has been successfully created")
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
		var user UserLogin
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			logger("", "error parsing json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := user.UserName
		password := user.Password

		fmt.Printf("username: %v, password: %v\n", username, password)
		logger(username, "user login input")

		var Id, hash string
		stmt := "SELECT Id, Password FROM testdb.users WHERE UserName =?"
		row := db.QueryRow(stmt, username)
		err := row.Scan(&Id, &hash)
		fmt.Println("hass:", hash)
		if err != nil {
			logger(username, "username not found")
			fmt.Fprint(w, "username not found")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			session, _ := store.Get(r, "session")
			if Id, ok := session.Values["Id"]; ok {
				stmt = "SELECT UserName FROM testdb.users WHERE Id =?"
				row = db.QueryRow(stmt, Id)
				_ = row.Scan(&username)
				fmt.Println("ok:", ok)
				if ok {
					logger(username, "user already logged in")
					fmt.Fprint(w, "You are already logged in")
					return
				}
			}
			session.Values["Id"] = Id
			session.Save(r, w)
			logger(username, "user has logged in")
			fmt.Fprint(w, "You have successfully logged in")
			return
		}
		logger(username, "incorrect password")
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
	stmt := "SELECT FirstName FROM testdb.users WHERE Id =?"
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
		var username string
		session, _ := store.Get(r, "session")
		Id, ok := session.Values["Id"]
		stmt := "SELECT UserName FROM testdb.users WHERE Id =?"
		row := db.QueryRow(stmt, Id)
		err := row.Scan(&username)
		if err != nil || !ok {
			logger(username, "user not logged in")
			fmt.Fprintln(w, "You are not logged in.")
			return
		}
		delete(session.Values, "Id")
		session.Save(r, w)
		logger(username, "user has logged out")
		fmt.Fprintln(w, "You have logged out")
	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func movies(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	Id, ok := session.Values["Id"]
	fmt.Println("ok:", ok)
	stmt := "SELECT UserName, FirstName FROM testdb.users WHERE Id =?"
	row := db.QueryRow(stmt, Id)
	var username, firstname string
	err := row.Scan(&username, &firstname)
	if err != nil || !ok {
		logger(username, "user not logged in")
		fmt.Fprintln(w, "You are not logged in.")
		return
	}
	fmt.Println("firstname:", firstname)

	switch r.Method {
	case "POST":
		var movie Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			logger("", "error parsing json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		title := movie.Title
		genre := movie.Genre
		year := movie.Year
		director := movie.Director
		language := movie.Language
		country := movie.Country
		status := movie.Status

		logger(username, "user movie input")

		if title == "" || genre == "" || year == 0 || director == "" || language == "" || country == "" || status == "" {
			logger(username, "user must fill in all fields")
			fmt.Fprintln(w, "Please fill in all fields")
			return
		}

		movie_stmt := "SELECT Id FROM testdb.movies WHERE UserName = ? AND Title = ? AND Genre = ? AND Year = ? AND Director = ? AND Language = ? AND Country= ?"
		movie_row := db.QueryRow(movie_stmt, username, title, genre, year, director, language, country)
		var movieId string
		err = movie_row.Scan(&movieId)
		if err != sql.ErrNoRows {
			logger(username, "movie already added on list")
			fmt.Fprintln(w, "Movie already added on list")
			return
		}

		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO testdb.movies (UserName, Title, Genre, Year, Director, Language, Country, Status) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			logger(username, "error preparing insert statement")
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
			logger(username, "error adding movie")
			return
		}

		logger(username, "movie added")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Hi %v! Movie has been added.", firstname)
	case "GET":
		query_stmt := "SELECT * FROM testdb.movies WHERE UserName = ?;"
		rows, err := db.Query(query_stmt, username)
		if err != nil {
			logger(username, "empty movie watchlist")
			fmt.Fprintln(w, "empty movie watchlist")
		}
		defer rows.Close()

		var movies []Movie

		for rows.Next() {
			var movie Movie
			err = rows.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year, &movie.Director, &movie.Language, &movie.Country, &movie.Status, &movie.UserName)
			if err != nil {
				logger(username, "error movie watchlist")
				fmt.Fprintln(w, "error movie watchlist")
			}
			movies = append(movies, movie)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(movies); err != nil {
			logger(username, "error encoding json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		logger(username, "success retrieving movie watchlist")
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

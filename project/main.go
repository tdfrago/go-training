package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

//User struct for user details
type User struct {
	Id        int
	LastName  string
	FirstName string
	UserName  string
	Password  string
}

//UserLogin struct for user login details
type UserLogin struct {
	UserName string
	Password string
}

//Movie struct for movie details
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

//Database connection
var db *sql.DB

//Initialize session store
var store = sessions.NewCookieStore([]byte("super-secret"))

//main runs the handler which receives and processes http requests
func main() {
	var err error
	db, err = sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
	if err != nil {
		fmt.Println("error validatin sql.open arguments")
		panic(err.Error())
	}
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/movies", moviesHandler)
	http.HandleFunc("/movies/", moviesIdHandler)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}

//signupHandler creates new user in the database
func signupHandler(w http.ResponseWriter, r *http.Request) {
	//checks if url is correct
	if r.URL.Path != "/signup" {
		WarningLogger.Println("incorrect url")
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	//POST creates a new user
	case "POST":
		var user User
		//parse json request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			ErrorLogger.Println("error parsing json")
			fmt.Println("error parsing json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var Id string

		lastname := user.LastName
		firstname := user.FirstName
		username := user.UserName
		password := user.Password

		fmt.Printf("lastname:%v,firstname:%v,username:%v,password:%v\n", lastname, firstname, username, password)

		InfoLogger.Println(username + " entered signup details")
		fmt.Println(username + " entered signup details")

		//check if username already exists in database
		stmt := "SELECT Id FROM users WHERE username = ?"
		row := db.QueryRow(stmt, username)

		err := row.Scan(&Id)
		if err != sql.ErrNoRows {
			WarningLogger.Println(username + " already exists")
			fmt.Fprintln(w, "Username already exists")
			return
		}

		InfoLogger.Println(username + " username does not yet exist")
		fmt.Println(username + " username does not yet exist")

		//creates a hash for password
		var hash []byte
		hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			ErrorLogger.Println(username + " has error hashing password")
			fmt.Println(username + " has error hashing password")
			return
		}
		fmt.Println("hash:", hash)
		fmt.Println("string(hash):", string(hash))

		InfoLogger.Println(username + " password converted to hash")
		fmt.Println(username + " password converted to hash")

		//prepares insert statement for user creation
		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO users (LastName, FirstName, UserName, Password) VALUES (?, ?, ?, ?);")
		if err != nil {
			ErrorLogger.Println("error preparing insert statement")
			fmt.Println("error statement:", err)
			return
		}
		defer insert_stmt.Close()

		//executes insert statement for user creation
		var result sql.Result
		result, err = insert_stmt.Exec(lastname, firstname, username, hash)
		rows_affected, _ := result.RowsAffected()

		//checks for errors and if only 1 statement was implemented
		if err != nil || rows_affected != 1 {
			ErrorLogger.Println(username + " has error registering as new user")
			fmt.Println(username + " has error registering as new user")
			return
		}
		InfoLogger.Println(username + " account has been succesfully created")
		fmt.Fprintln(w, "User has been successfully created")
	//GET method not allowed
	case "GET":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//PUT method not allowed
	case "PUT":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//DELETE method not allowed
	case "DELETE":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//loginHandler authenticates user login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	//checks if url is correct
	if r.URL.Path != "/login" {
		WarningLogger.Println("incorrect url")
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	case "POST":
		var user UserLogin
		//parse json request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			ErrorLogger.Println("error parsing json")
			fmt.Println("error parsing json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := user.UserName
		password := user.Password

		fmt.Printf("username: %v, password: %v\n", username, password)
		InfoLogger.Println(username + " entered login details")
		fmt.Println(username + " entered login details")

		//retrieves hash from database
		var Id, hash string
		stmt := "SELECT Id, Password FROM users WHERE UserName =?"
		row := db.QueryRow(stmt, username)
		err := row.Scan(&Id, &hash)
		fmt.Println("hash:", hash)
		if err != nil {
			WarningLogger.Println(username + " not found")
			fmt.Fprint(w, "username not found")
			return
		}

		//compares retrieved hash and the entered login password's hash
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			//checks for existing session cookies - determines if a user is currently logged in
			session, _ := store.Get(r, "session")
			if Id, ok := session.Values["Id"]; ok {
				stmt = "SELECT UserName FROM users WHERE Id =?"
				row = db.QueryRow(stmt, Id)
				_ = row.Scan(&username)
				fmt.Println("ok:", ok)
				if ok {
					WarningLogger.Println(username + " already logged in")
					fmt.Fprint(w, "You are already logged in")
					return
				}
			}
			//creates session cookies
			session.Values["Id"] = Id
			session.Save(r, w)

			InfoLogger.Println(username + " has logged in")
			fmt.Fprint(w, "You have successfully logged in")
			return
		}

		//returned if login password did not match retrieved hash
		WarningLogger.Println(username + " entered incorrect password")
		fmt.Fprint(w, "Incorrect password")
	//GET method not allowed
	case "GET":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//PUT method not allowed
	case "PUT":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//DELETE method not allowed
	case "DELETE":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//homeHandler serves a homepage view
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//checks if url is correct
	if r.URL.Path != "/" {
		WarningLogger.Println("incorrect url")
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	//POST method not allowed
	case "POST":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//GET shows short webapp introduction
	case "GET":
		InfoLogger.Println("viewing home page")
		fmt.Fprintln(w, "Welcome to home page")
		fmt.Fprintln(w, "This webapp allows a user to make his/her own movie watchlist.")
		fmt.Fprintln(w, "A user can signup to the webapp, login to create, edit, and view the movie watchlist, and logout when done.")
	//PUT method not allowed
	case "PUT":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//DELETE method not allowed
	case "DELETE":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//logoutHandler deletes session cookies of logged in user
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	//checks if the url is correct
	if r.URL.Path != "/logout" {
		WarningLogger.Println("incorrect url")
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	//POST deletes the session cookies
	case "POST":
		//checks for existing session cookies - determines if a user is currently logged in
		var username string
		session, _ := store.Get(r, "session")
		Id, ok := session.Values["Id"]
		stmt := "SELECT UserName FROM users WHERE Id =?"
		row := db.QueryRow(stmt, Id)
		err := row.Scan(&username)
		if err != nil || !ok {
			WarningLogger.Println("user not logged in")
			fmt.Fprintln(w, "You are not logged in.")
			return
		}
		//deletes the cookies
		delete(session.Values, "Id")
		session.Save(r, w)
		InfoLogger.Println("user has logged out")
		fmt.Fprintln(w, "You have logged out")
	//GET method not allowed
	case "GET":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//PUT method not allowed
	case "PUT":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//DELETE method not allowed
	case "DELETE":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//moviesHandler creates a movie entry and shows the user's movie watchlist
func moviesHandler(w http.ResponseWriter, r *http.Request) {
	//checks if url is correct
	if r.URL.Path != "/movies" {
		WarningLogger.Println("incorrect url")
		fmt.Fprintln(w, "incorrect url")
		return
	}
	//checks for existing session cookies - determines if a user is currently logged in
	session, _ := store.Get(r, "session")
	Id, ok := session.Values["Id"]
	fmt.Println("ok:", ok)
	stmt := "SELECT UserName FROM users WHERE Id =?"
	row := db.QueryRow(stmt, Id)
	var username string
	err := row.Scan(&username)
	if err != nil || !ok {
		WarningLogger.Println("user not logged in")
		fmt.Fprintln(w, "You are not logged in.")
		return
	}
	InfoLogger.Println(username + " is currently logged in")
	fmt.Printf("User %v is logged in\n", username)

	switch r.Method {
	//POST creates a movie entry
	case "POST":
		//parse json request body
		var movie Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			ErrorLogger.Println("error parsing json")
			fmt.Println("error parsing json")
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

		InfoLogger.Println(username + " entered movie input")
		fmt.Println(username + " entered movie input")

		//checks for any empty fields
		if title == "" || genre == "" || year == 0 || director == "" || language == "" || country == "" || status == "" {
			ErrorLogger.Println(username + " must fill in all fields")
			fmt.Fprintln(w, "Please fill in all fields")
			return
		}
		//checks if movie entry already exists in database
		movie_stmt := "SELECT Id FROM movies WHERE UserName = ? AND Title = ? AND Genre = ? AND Year = ? AND Director = ? AND Language = ? AND Country= ?"
		movie_row := db.QueryRow(movie_stmt, username, title, genre, year, director, language, country)
		var movieId string
		err = movie_row.Scan(&movieId)
		if err != sql.ErrNoRows {
			WarningLogger.Println(username + " already added the movie on watchlist")
			fmt.Fprintln(w, "Movie already added on list")
			return
		}

		//prepares insert statement for movie entry creation
		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO movies (UserName, Title, Genre, Year, Director, Language, Country, Status) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			ErrorLogger.Println("error preparing insert statement")
			fmt.Println("error statement:", err)
			return
		}
		defer insert_stmt.Close()

		//executes insert statement for movie entry creation
		result, err := insert_stmt.Exec(username, title, genre, year, director, language, country, status)
		rows_affected, _ := result.RowsAffected()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Println("error:", err)

		if err != nil || rows_affected != 1 {
			ErrorLogger.Println("error adding movie")
			return
		}
		//updates user that the movie has been added
		InfoLogger.Println(username + " successfully added the move")
		fmt.Println(username + " successfully added the move")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Hi %v! Movie has been added.", username)
	//GET retrieves the user's movie watchlist
	case "GET":
		//queries all the movies under the user
		var movies []Movie
		query_stmt := "SELECT * FROM movies WHERE UserName = ?;"
		rows, err := db.Query(query_stmt, username)
		if err != nil {
			WarningLogger.Println("empty movie watchlist")
			fmt.Fprintln(w, "empty movie watchlist")
			return
		}
		defer rows.Close()
		//appends each movie to movies (watchlist)
		for rows.Next() {
			var movie Movie
			err = rows.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year, &movie.Director, &movie.Language, &movie.Country, &movie.Status, &movie.UserName)
			if err != nil {
				ErrorLogger.Println("error movie watchlist")
				fmt.Fprintln(w, "error movie watchlist")
				return
			}
			movies = append(movies, movie)
		}

		//show the user's movie watchlist
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(movies); err != nil {
			ErrorLogger.Println("error encoding json")
			fmt.Println("error encoding json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		InfoLogger.Println(username + " successfully retrieved the movie watchlist")
		fmt.Println(username + " successfully retrieved the movie watchlist")
	//PUT method not allowed
	case "PUT":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//DELETE method not allowed
	case "DELETE":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//moviesIdHandler retrieves movie by id, updates movie by id, and deletes movie by id
func moviesIdHandler(w http.ResponseWriter, r *http.Request) {
	//checks if url is correct
	if r.URL.Path != "/movies/" {
		WarningLogger.Println("incorrect url must be '/movies/?Id={id}'")
		fmt.Fprintln(w, "incorrect url must be '/movies/?Id={id}'")
		return
	}
	//checks for existing session cookies - determines if a user is currently logged in
	session, _ := store.Get(r, "session")
	Id, ok := session.Values["Id"]
	fmt.Println("ok:", ok)
	stmt := "SELECT UserName FROM users WHERE Id =?"
	row := db.QueryRow(stmt, Id)
	var username string
	err := row.Scan(&username)
	if err != nil || !ok {
		WarningLogger.Println("user not logged in")
		fmt.Fprintln(w, "You are not logged in.")
		return
	}

	InfoLogger.Println(username + " is currently logged in")
	fmt.Printf("User %v is logged in\n", username)

	//retrieves movie Id and checks if movie is in database
	r.ParseForm()
	id := r.FormValue("Id")
	movie_stmt := "SELECT * FROM movies WHERE UserName = ? AND Id = ?"
	movie_row := db.QueryRow(movie_stmt, username, id)
	var movie Movie
	err = movie_row.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year, &movie.Director, &movie.Language, &movie.Country, &movie.Status, &movie.UserName)
	if err != nil {
		InfoLogger.Println(username + " did not find the movie")
		fmt.Fprintln(w, "movie not found")
		return
	}
	switch r.Method {
	//POST method not allowed
	case "POST":
		WarningLogger.Println("status method not allowed")
		fmt.Println("status method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	//GET retrieves the movie by id
	case "GET":
		//shows the movie details to the user
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(movie); err != nil {
			ErrorLogger.Println("error encoding json")
			fmt.Println("error encoding json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		InfoLogger.Println(username + " successfully retrieved the movie")
		fmt.Println(username + " successfully retrieved the movie")
	//PUT updates the movie details
	case "PUT":
		//parse json request body
		var movie_update Movie
		if err := json.NewDecoder(r.Body).Decode(&movie_update); err != nil {
			ErrorLogger.Println("error parsing json")
			fmt.Println("error parsing json")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//updates/edits movie fields from request
		if movie_update.Title != "" {
			movie.Title = movie_update.Title
		}
		if movie_update.Genre != "" {
			movie.Genre = movie_update.Genre
		}
		if movie_update.Year != 0 {
			movie.Year = movie_update.Year
		}
		if movie_update.Director != "" {
			movie.Director = movie_update.Director
		}
		if movie_update.Language != "" {
			movie.Language = movie_update.Language
		}
		if movie_update.Country != "" {
			movie.Country = movie_update.Country
		}
		if movie_update.Status != "" {
			movie.Status = movie_update.Status
		}
		//prepares update statement for movie entry
		var update_stmt *sql.Stmt
		update_stmt, err := db.Prepare("UPDATE movies SET Title = ?, Genre = ?, Year = ?, Director = ?, Language = ?, Country = ?, Status = ? WHERE UserName = ? AND Id =?;")
		if err != nil {
			ErrorLogger.Println("error preparing update statement")
			fmt.Println("error statement:", err)
			return
		}
		defer update_stmt.Close()

		//executes update statement for movie entry
		result, err := update_stmt.Exec(movie.Title, movie.Genre, movie.Year, movie.Director, movie.Language, movie.Country, movie.Status, username, id)
		rows_affected, _ := result.RowsAffected()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Println("error:", err)
		if err != nil || rows_affected != 1 {
			InfoLogger.Println(username + " was not able to update the movie")
			fmt.Println("error statement:", err)
			fmt.Fprintf(w, "Hi %v! No changes made.", username)
			return
		}
		//updates user that the movie entry has been edited
		InfoLogger.Println(username + " updated the movie")
		fmt.Fprintf(w, "Hi %v! Movie has been updated.", username)
	//DELETE deletes the movie entry
	case "DELETE":
		//prepares delete statement for movie entry
		var delete_stmt *sql.Stmt
		delete_stmt, err = db.Prepare("DELETE FROM movies WHERE UserName = ? AND Id =?;")
		if err != nil {
			ErrorLogger.Println("error preparing delete statement")
			fmt.Println("error statement:", err)
			return
		}
		defer delete_stmt.Close()

		//executes delete statement for movie entry
		result, err := delete_stmt.Exec(username, id)
		rows_affected, _ := result.RowsAffected()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Println("error:", err)
		if err != nil || rows_affected != 1 {
			InfoLogger.Println(username + " was not able to delete the movie")
			fmt.Fprintf(w, "Hi %v! No changes made.", username)
			return
		}
		//updates user that the movie entry has been deleted
		InfoLogger.Println(username + " deleted the movie")
		fmt.Fprintf(w, "Hi %v! Movie has been delete.", username)
	}
}

//Custom logger variables
var (
	WarningLogger *log.Logger //log for not found/mismatched entries
	InfoLogger    *log.Logger //log for user movement
	ErrorLogger   *log.Logger //log for errors
)

//initializes logs
func init() {
	//saves logs to a txt file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//initializes custom loggers
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

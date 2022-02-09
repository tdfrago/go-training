package main

import (
	"database/sql"
	"encoding/json"
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

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	case "POST":
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var Id string

		lastname := user.LastName
		firstname := user.FirstName
		username := user.UserName
		password := user.Password

		fmt.Printf("lastname:%v,firstname:%v,username:%v,password:%v\n", lastname, firstname, username, password)

		stmt := "SELECT Id FROM users WHERE username = ?"
		row := db.QueryRow(stmt, username)

		err := row.Scan(&Id)
		if err != sql.ErrNoRows {
			fmt.Fprintln(w, "Username already exists")
			return
		}

		var hash []byte
		hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		fmt.Println("hash:", hash)
		fmt.Println("string(hash):", string(hash))

		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO users (LastName, FirstName, UserName, Password) VALUES (?, ?, ?, ?);")
		if err != nil {
			return
		}
		defer insert_stmt.Close()

		var result sql.Result
		result, err = insert_stmt.Exec(lastname, firstname, username, hash)
		rows_affected, _ := result.RowsAffected()

		if err != nil || rows_affected != 1 {
			return
		}
		fmt.Fprintln(w, "User has been successfully created")
	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	case "POST":
		var user UserLogin
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := user.UserName
		password := user.Password

		fmt.Printf("username: %v, password: %v\n", username, password)

		var Id, hash string
		stmt := "SELECT Id, Password FROM users WHERE UserName =?"
		row := db.QueryRow(stmt, username)
		err := row.Scan(&Id, &hash)
		fmt.Println("hass:", hash)
		if err != nil {
			fmt.Fprint(w, "username not found")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			session, _ := store.Get(r, "session")
			if Id, ok := session.Values["Id"]; ok {
				stmt = "SELECT UserName FROM users WHERE Id =?"
				row = db.QueryRow(stmt, Id)
				_ = row.Scan(&username)
				fmt.Println("ok:", ok)
				if ok {
					fmt.Fprint(w, "You are already logged in")
					return
				}
			}
			session.Values["Id"] = Id
			session.Save(r, w)
			fmt.Fprint(w, "You have successfully logged in")
			return
		}
		fmt.Fprint(w, "Incorrect password")

	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	case "POST":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "GET":
		fmt.Fprintf(w, "Welcome to home page")
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		fmt.Fprintln(w, "incorrect url")
		return
	}
	switch r.Method {
	case "POST":
		var username string
		session, _ := store.Get(r, "session")
		Id, ok := session.Values["Id"]
		stmt := "SELECT UserName FROM users WHERE Id =?"
		row := db.QueryRow(stmt, Id)
		err := row.Scan(&username)
		if err != nil || !ok {
			fmt.Fprintln(w, "You are not logged in.")
			return
		}
		delete(session.Values, "Id")
		session.Save(r, w)
		fmt.Fprintln(w, "You have logged out")
	case "GET":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/movies" {
		fmt.Fprintln(w, "incorrect url")
		return
	}
	session, _ := store.Get(r, "session")
	Id, ok := session.Values["Id"]
	fmt.Println("ok:", ok)
	stmt := "SELECT UserName FROM users WHERE Id =?"
	row := db.QueryRow(stmt, Id)
	var username string
	err := row.Scan(&username)
	if err != nil || !ok {
		fmt.Fprintln(w, "You are not logged in.")
		return
	}
	fmt.Printf("User %v is logged in", username)

	switch r.Method {
	case "POST":
		var movie Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
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

		if title == "" || genre == "" || year == 0 || director == "" || language == "" || country == "" || status == "" {
			fmt.Fprintln(w, "Please fill in all fields")
			return
		}

		movie_stmt := "SELECT Id FROM movies WHERE UserName = ? AND Title = ? AND Genre = ? AND Year = ? AND Director = ? AND Language = ? AND Country= ?"
		movie_row := db.QueryRow(movie_stmt, username, title, genre, year, director, language, country)
		var movieId string
		err = movie_row.Scan(&movieId)
		if err != sql.ErrNoRows {
			fmt.Fprintln(w, "Movie already added on list")
			return
		}

		var insert_stmt *sql.Stmt
		insert_stmt, err = db.Prepare("INSERT INTO movies (UserName, Title, Genre, Year, Director, Language, Country, Status) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			fmt.Println("error statement:", err)
			return
		}
		defer insert_stmt.Close()

		result, err := insert_stmt.Exec(username, title, genre, year, director, language, country, status)
		rows_affected, _ := result.RowsAffected()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Println("error:", err)

		if err != nil || rows_affected != 1 {
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Hi %v! Movie has been added.", username)

	case "GET":
		query_stmt := "SELECT * FROM movies WHERE UserName = ?;"
		rows, err := db.Query(query_stmt, username)
		if err != nil {
			fmt.Fprintln(w, "empty movie watchlist")
			return
		}
		defer rows.Close()

		var movies []Movie

		for rows.Next() {
			var movie Movie
			err = rows.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year, &movie.Director, &movie.Language, &movie.Country, &movie.Status, &movie.UserName)
			if err != nil {
				fmt.Fprintln(w, "error movie watchlist")
				return
			}
			movies = append(movies, movie)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(movies); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "PUT":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "DELETE":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func moviesIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/movies/" {
		fmt.Fprintln(w, "incorrect url must be '/movies/?Id={id}'")
		return
	}
	session, _ := store.Get(r, "session")
	Id, ok := session.Values["Id"]
	fmt.Println("ok:", ok)
	stmt := "SELECT UserName FROM users WHERE Id =?"
	row := db.QueryRow(stmt, Id)
	var username string
	err := row.Scan(&username)
	if err != nil || !ok {
		fmt.Fprintln(w, "You are not logged in.")
		return
	}
	fmt.Printf("User %v is logged in", username)

	r.ParseForm()
	id := r.FormValue("Id")
	movie_stmt := "SELECT * FROM movies WHERE UserName = ? AND Id = ?"
	movie_row := db.QueryRow(movie_stmt, username, id)
	var movie Movie
	err = movie_row.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year, &movie.Director, &movie.Language, &movie.Country, &movie.Status, &movie.UserName)
	if err != nil {
		fmt.Fprintln(w, "movie not found")
		return
	}

	switch r.Method {
	case "POST":
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(movie); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PUT":
		var movie_update Movie
		if err := json.NewDecoder(r.Body).Decode(&movie_update); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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

		var update_stmt *sql.Stmt
		update_stmt, err := db.Prepare("UPDATE movies SET Title = ?, Genre = ?, Year = ?, Director = ?, Language = ?, Country = ?, Status = ? WHERE UserName = ? AND Id =?;")
		if err != nil {
			fmt.Println("error statement:", err)
			return
		}
		defer update_stmt.Close()

		result, err := update_stmt.Exec(movie.Title, movie.Genre, movie.Year, movie.Director, movie.Language, movie.Country, movie.Status, username, id)
		rows_affected, _ := result.RowsAffected()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Println("error:", err)

		if err != nil || rows_affected != 1 {
			fmt.Println("error statement:", err)
			fmt.Fprintf(w, "Hi %v! No changes made.", username)
			return
		}

		fmt.Fprintf(w, "Hi %v! Movie has been updated.", username)

	case "DELETE":
		var delete_stmt *sql.Stmt
		delete_stmt, err = db.Prepare("DELETE FROM movies WHERE UserName = ? AND Id =?;")
		if err != nil {
			fmt.Println("error statement:", err)
			return
		}
		defer delete_stmt.Close()

		result, err := delete_stmt.Exec(username, id)
		rows_affected, _ := result.RowsAffected()
		fmt.Println("number of rows affected:", rows_affected)
		fmt.Println("error:", err)

		if err != nil || rows_affected != 1 {
			fmt.Fprintf(w, "Hi %v! No changes made.", username)
			return
		}

		fmt.Fprintf(w, "Hi %v! Movie has been delete.", username)
	}
}

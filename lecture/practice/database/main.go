package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
}

var tpl *template.Template

var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/productsearch", productSearchHandler)
	http.HandleFunc("/productsearch2", productSearchHandler2)
	http.ListenAndServe("localhost:8080", nil)
	/*
		err = db.Ping()
		if err != nil {
			fmt.Println("error verrifying connection with db.Ping")
			panic(err.Error())
		}

		insert, err := db.Query("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('2', 'Ben', 'Ford');")
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		fmt.Println("Successful Connection to Database")
	*/
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home")
}

func productSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "productsearch.html", nil)
		return
	}
	r.ParseForm()
	var P Product
	name := r.FormValue("productName")
	fmt.Println("name:", name)

	stmt := "SELECT * FROM products WHERE name = ?;"

	row := db.QueryRow(stmt, name)

	err := row.Scan(&P.ID, &P.Name, &P.Price, &P.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintln(w, "Not found")
		} else {
			log.Fatal(err)
		}
	} else {
		tpl.ExecuteTemplate(w, "productsearch.html", P)
	}
}

func productSearchHandler2(w http.ResponseWriter, r *http.Request) {

}

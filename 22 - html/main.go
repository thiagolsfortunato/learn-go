package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Name string
	Email string
}

var templates *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	u := user{"John", "john.peter@gmail.com"}

	templates.ExecuteTemplate(w, "users.html", u)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)


	fmt.Println("Server listening on port 5001")
	
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
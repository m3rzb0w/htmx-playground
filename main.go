package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Book struct {
	Title  string
	Author string
}

const (
	maxBooks = 15
)

var books = map[string][]Book{
	"Books": {
		{Title: "Hacker's Delight", Author: "Henry S. Warren"},
		{Title: "Hypermedia Systems", Author: "CARSON GROSS, ADAM STEPINSKI, DENİZ AKŞİMŞEK"},
		{Title: "Game Engine Black Book: DOOM: v1.1", Author: "Fabien Sanglard"},
		{Title: "Masters of Doom: How Two Guys Created an Empire and Transformed Pop Culture", Author: "David Kushner"},
	},
}

func main() {
	fmt.Println("[INFO]: Starting server...")
	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, books)
	}

	addBookHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		if len(title) == 0 || len(author) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Fill out every inputs!\n"))
			return
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "book-list-element", Book{Title: title, Author: author})
	}

	upBookHandler := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		if len(title) == 0 || len(author) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Fill out every inputs!\n"))
			return
		}
		if len(books["Books"]) == maxBooks {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Sorry all slots are filled!\n"))
			return
		}
		book := Book{Title: title, Author: author}
		books["Books"] = append([]Book{book}, books["Books"]...)
		w.WriteHeader(http.StatusAccepted)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/add-book", addBookHandler)
	http.HandleFunc("/up-book", upBookHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

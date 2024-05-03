package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", index)

	router.Post("/subscribe", subscribe)

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))

	fmt.Println("Sever is running on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := renderTemplate()
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func renderTemplate() *template.Template {
	tmpl, err := template.ParseFiles("ui/tmpl/index.html")
	if err != nil {
		panic(err)
	}
	return tmpl
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	fmt.Printf("Name: %s, Email: %s\n", name, email)
	w.Write([]byte("<h2>Subscription successful!</h2>"))
}

package main

import (
	"html/template"
	"net/http"
)

func main() {
	// Serve static files (CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/skills", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/skills.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/blog.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/links", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/links.html"))
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe("localhost:8080", nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// messageHandler handles requests to the /api/message endpoint.
func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to text/html
	w.Header().Set("Content-Type", "text/html")

	// Write a simple HTML snippet as the response
	fmt.Fprint(w, "<div><p>Hello, this is a message from the server!</p></div>")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Register the message handler for the /api/message endpoint
	http.HandleFunc("/api/message", messageHandler)

	http.ListenAndServe("localhost:8080", nil)
}

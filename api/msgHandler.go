// In Go, add a handler for /api/message that returns a partial HTML snippet.
package main

import (
    "fmt"
    "net/http"
)

// messageHandler handles requests to the /api/message endpoint.
func messageHandler(w http.ResponseWriter, r *http.Request) {
    // Set the content type to text/html
    w.Header().Set("Content-Type", "text/html")

    // Write a simple HTML snippet as the response
    fmt.Fprint(w, "<div><p>Hello, this is a message from the server!</p></div>")
}


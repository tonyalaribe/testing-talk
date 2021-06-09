package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// Check if aa user request is authenticated
func authFunction(r *http.Request) bool {
  return r.URL.Query().Get("auth") == "ABC"
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authFunction(r) {
      http.Redirect(w, r, "/notloggedIn", http.StatusUnauthorized)
			return
		}

		// User authenticated successfully
		next(w, r)
	})
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func notLoggedInHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not logged in"))
}

func main() {
	http.HandleFunc("/submit", authMiddleware(handlerFunc))
	http.HandleFunc("/notloggedIn", notLoggedInHandlerFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

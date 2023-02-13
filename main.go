package main

import (
	"fmt"
	"net/http"
)

//	type Handler interface {
//		ServeHTTP(ResponseWriter, *Request)
//	}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my website!`)
}

func myLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "USING GET for login endpoint\n")
	case "POST":
		fmt.Fprintf(w, "USING POST for login endpoint\n")

	}
	fmt.Fprintf(w, "Login")
}

func main() {
	// http.HandleFunc("/login", myLogin)
	// http.HandleFunc("/welcome", myWelcome)
	http.Handle("/login", http.HandlerFunc(myLogin))
	http.Handle("/welcome", http.HandlerFunc(myWelcome))
	fmt.Println("Listening on port 8080....")

	http.ListenAndServe("localhost:8080", nil)
}

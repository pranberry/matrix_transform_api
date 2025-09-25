package main

import (
	"league_challenge/handlers"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/invert", handlers.Transpose)
	http.HandleFunc("/flatten", handlers.Flatten)
	http.HandleFunc("/add", handlers.Addition)
	http.HandleFunc("/mul", handlers.Multiply)
	http.ListenAndServe(":8080", nil)
}
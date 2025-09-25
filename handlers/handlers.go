package handlers

import (
	"fmt"
	"league_challenge/matrix"
	"log"
	"net/http"
)

// Print back the matrix
func Echo(w http.ResponseWriter, r *http.Request) {
	// expect a file with the key "file" in form data
	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, matrix.Echo())
}

// Invert/Transpose a NxN matrix...rows become columns, columns become rows
func Transpose(w http.ResponseWriter, r *http.Request){

	log.Println("Request: /invert from", r.RemoteAddr)

	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	matrix.Transpose()
	
	fmt.Fprint(w, matrix.Echo())
}


// Returns the flattened representation of the matrix
func Flatten(w http.ResponseWriter, r *http.Request){
	log.Println("Request: /flatten from", r.RemoteAddr)

	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	fmt.Fprint(w, matrix.Flatten())
}


// Returns the sum of all values in a matrix
func Addition(w http.ResponseWriter, r *http.Request){
	log.Println("Request: /add from", r.RemoteAddr)
	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sum, err := matrix.Add()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, sum)
}

// Returns the product of all values in a matrix
func Multiply(w http.ResponseWriter, r *http.Request){
	log.Println("Request: /mul from", r.RemoteAddr)
	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	prod, err := matrix.Multiply()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, prod)
}



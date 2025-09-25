package handlers

import (
	"fmt"
	"league_challenge/matrix"
	"log"
	"net/http"
)


func Echo(w http.ResponseWriter, r *http.Request) {
	// expect a file with the key "file" in form data
	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, matrix.Echo())
}

// "invert" a NxN matrix...rows become columns, columns become rows
func Transpose(w http.ResponseWriter, r *http.Request){

	log.Println("Request: /invert from", r.RemoteAddr)

	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	matrix.Transpose()
	
	fmt.Fprint(w, matrix.Echo())
}



func Flatten(w http.ResponseWriter, r *http.Request){
	log.Println("Request: /flatten from", r.RemoteAddr)

	matrix, err := matrix.NewMatrix(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	fmt.Fprint(w, matrix.Flatten())
}





/*
	In this file/package live functions which perform arithmetic on the uploaded file.
*/

// Returns the sum of all values in the matrix-file
func Addition(w http.ResponseWriter, r *http.Request){
	log.Println("Request: /add from", r.RemoteAddr)
	fmt.Fprint(w, "addition endpoint")
}

// Returns the sum of all values in the matrix-file
func Multiply(w http.ResponseWriter, r *http.Request){
	log.Println("Request: /mul from", r.RemoteAddr)
	fmt.Fprint(w, "multipily endpoint")
}



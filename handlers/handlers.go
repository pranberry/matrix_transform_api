package handlers

import (
	"fmt"
	"league_challenge/matrix"
	"log"
	"net/http"
)

// Print back the matrix
func Echo(w http.ResponseWriter, r *http.Request) {

	// before and after logging
	log.Printf("REQUEST: method=%v on url=%v from remote=%v", r.Method, r.URL.Path, r.RemoteAddr)
	reqStatus := http.StatusBadRequest
	defer func() {
		log.Printf("RESPONSE: status=%d on method=%v on url=%v from %v", reqStatus, r.Method, r.URL.Path, r.RemoteAddr)
	}()

	// get matrix
	matrix, err := matrix.NewMatrix(r)
	if err != nil {
		http.Error(w, err.Error(), reqStatus)
		return
	}

	// write and return response
	w.Header().Set("Content-Type", "text/csv")
	reqStatus = http.StatusOK
	w.WriteHeader(reqStatus)
	fmt.Fprint(w, matrix.Echo())
}

// Invert/Transpose a NxN matrix...rows become columns, columns become rows
func Transpose(w http.ResponseWriter, r *http.Request) {

	log.Printf("REQUEST: method=%v on url=%v from remote=%v", r.Method, r.URL.Path, r.RemoteAddr)
	reqStatus := http.StatusBadRequest
	defer func() {
		log.Printf("RESPONSE: status=%d on method=%v on url=%v from %v", reqStatus, r.Method, r.URL.Path, r.RemoteAddr)
	}()

	matrix, err := matrix.NewMatrix(r)
	if err != nil {
		http.Error(w, err.Error(), reqStatus)
		return
	}

	// call to actually mutate (transpose) the matrix
	matrix.Transpose()

	// write and return response
	w.Header().Set("Content-Type", "text/csv")
	reqStatus = http.StatusOK
	w.WriteHeader(reqStatus)
	fmt.Fprint(w, matrix.Echo())
}

// Returns the flattened representation of the matrix
func Flatten(w http.ResponseWriter, r *http.Request) {

	log.Printf("REQUEST: method=%v on url=%v from remote=%v", r.Method, r.URL.Path, r.RemoteAddr)
	reqStatus := http.StatusBadRequest
	defer func() {
		log.Printf("RESPONSE: status=%d on method=%v on url=%v from %v", reqStatus, r.Method, r.URL.Path, r.RemoteAddr)
	}()

	matrix, err := matrix.NewMatrix(r)
	if err != nil {
		http.Error(w, err.Error(), reqStatus)
		return
	}

	reqStatus = http.StatusOK
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(reqStatus)
	// flattened response returned here
	fmt.Fprint(w, matrix.Flatten())
}

// Returns the sum of all values in a matrix
func Addition(w http.ResponseWriter, r *http.Request) {

	log.Printf("REQUEST: method=%v on url=%v from remote=%v", r.Method, r.URL.Path, r.RemoteAddr)
	reqStatus := http.StatusBadRequest
	defer func() {
		log.Printf("RESPONSE: status=%d on method=%v on url=%v from %v", reqStatus, r.Method, r.URL.Path, r.RemoteAddr)
	}()

	matrix, err := matrix.NewMatrix(r)
	if err != nil {
		http.Error(w, err.Error(), reqStatus)
		return
	}

	// calculate sum
	sum, err := matrix.Add()
	if err != nil {
		http.Error(w, err.Error(), reqStatus)
		return
	}

	reqStatus = http.StatusOK
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(reqStatus)
	fmt.Fprint(w, sum)
}

// Returns the product of all values in a matrix
func Multiply(w http.ResponseWriter, r *http.Request) {

	log.Printf("REQUEST: method=%v on url=%v from remote=%v", r.Method, r.URL.Path, r.RemoteAddr)
	reqStatus := http.StatusBadRequest
	defer func() {
		log.Printf("RESPONSE: status=%d on method=%v on url=%v from %v", reqStatus, r.Method, r.URL.Path, r.RemoteAddr)
	}()

	matrix, err := matrix.NewMatrix(r)
	if err != nil {
		http.Error(w, err.Error(), reqStatus)
		return
	}

	// calculate product
	prod, err := matrix.Multiply()
	if err != nil {
		http.Error(w, err.Error(), reqStatus)
		return
	}

	reqStatus = http.StatusOK
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(reqStatus)
	fmt.Fprint(w, prod)
}

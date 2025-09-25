package matrix

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

// IN this file we....
// could use a custom type in the return here.
// - type would hold the matrix, an error, an http status code

type ParseResp struct{
	Matrix *Matrix
	Status int
	Error error
}

// Loads transforms and
func NewMatrix(r *http.Request) (*Matrix, error) {

	// read from file
	const keyName string = "file"
	file, _, err := r.FormFile(keyName)
	if err != nil {
		return nil, fmt.Errorf("error: %s. must upload form file with key '%s'", err.Error(), keyName)
	}
	defer file.Close()
	
	// use use csv reader to return records
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error: %s", err.Error())
	}
	
	err = validateNxN(records)
	if err != nil{
		return nil, err
	}


	matrix := &Matrix{
		Data : records,
		Size: len(records),
	}

	return matrix, nil
}


// Strictly validates if matrix is !empty, and is NxN
func validateNxN(records [][]string) error {
	
	// csv.ReadAll() returns valid on empty file, check for empty records
	if len(records) == 0 {
		return fmt.Errorf("error: empty matrix")
	}

	// number of #rows == #cols(/#elements in each row) 
	for _, row := range records {
		if len(row) != len(records) {
			return fmt.Errorf("error: not an NxN matrix")
		}
	}
	return nil
}
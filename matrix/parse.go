package matrix

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

/*
	This file has ELT Operations:
	- Extracts file from Http.Request
	- Extracts matrix form file
	- Sanitizes the retrieved matrix
	- Loads into Matrix struct
*/


// Extracts file from http.request and returns valid Matrix
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

	// csv.ReadAll() returns valid on empty file, check for empty records
	if len(records) == 0 {
		return nil, fmt.Errorf("error: empty matrix")
	}
	
	// Validate matrix for NxN
	if err = validateNxN(records); err != nil {
		return nil, err
	}

	// Call to sanitize matrix
	cleanMatrix(records)

	// Initialize matrix
	matrix := &Matrix{
		Data: records,
		Size: len(records),
	}

	return matrix, nil
}

// Strictly validates on NxN-ness of matrix
// A matrix is a valid NxN if number of rows == number of columns per row
func validateNxN(records [][]string) error {
	// number of #rows must == #cols(/#elements in each row)
	for _, row := range records {
		if len(row) != len(records) {
			return fmt.Errorf("error: not an NxN matrix")
		}
	}
	return nil
}

// Sanitize matrix as desired.
// Trims spaces in each element.
// Commented out code to optionally, replace empty cells with "NA"
func cleanMatrix(records [][]string) {
	// trim spaces on each element. avoid conversion failures downstream
	for _, row := range records {
		for i, elem := range row {
			row[i] = strings.Trim(elem, " ")
			// if row[i] == "" {
			// 	row[i] = "NA"
			// }
		}
	}
}

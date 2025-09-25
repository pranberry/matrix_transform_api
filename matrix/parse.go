package matrix

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

// IN this file we....
// could use a custom type in the return here.
// - type would hold the matrix, an error, an http status code

type ParseResp struct {
	Matrix *Matrix
	Status int
	Error  error
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

	if err = validateNxN(records); err != nil {
		return nil, err
	}

	cleanMatrix(records)

	matrix := &Matrix{
		Data: records,
		Size: len(records),
	}

	return matrix, nil
}

// Parses and validates raw records returned by CSV.Reader.
// returns errors on empty matrix, and !NxN matrix.
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

// Sanitize matrix as desired.
// Trims spaces in each element.
// Optionally, replace empty cells with "NA"
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

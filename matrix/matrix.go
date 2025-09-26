package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	This files contains the Matrix struct definition and methods acting on the matrix type
*/

type Matrix struct {
	Data [][]string
	Size int
}

// Returns a string representation of the matrix.
// Echo() is frequently used.
func (m *Matrix) Echo() string {
	var b strings.Builder
	for _, row := range m.Data {
		b.WriteString(strings.Join(row, ","))
		b.WriteByte('\n')
	}
	return b.String()
}

// Transposes the matrix in-memory.
// Returns nothing, use m.Echo() to print.
func (m *Matrix) Transpose() {
	for row := 0; row < m.Size; row++ {
		for col := row + 1; col < m.Size; col++ {
			m.Data[row][col], m.Data[col][row] = m.Data[col][row], m.Data[row][col]
		}
	}
}

// Returns a flattened representation of the string.
// Nested slices get reduced to single slice.
// Sliced gets joined into a string.
func (m *Matrix) Flatten() string {
	var totalElements int
	for _, row := range m.Data {
		totalElements += len(row)
	}

	retMatrix := make([]string, 0, totalElements)
	for _, row := range m.Data {
		retMatrix = append(retMatrix, row...)
	}
	return strings.Join(retMatrix, ",")
}

// Returns the sum of all the values in the matrix.
// Returns error if non-ints are encountered.
func (m *Matrix) Add() (int, error) {
	sum := 0
	for _, row := range m.Data {
		for j := range row {
			v, err := strconv.Atoi(row[j])
			if err != nil {
				return 0, fmt.Errorf("error: non-int values in matrix. all values must be of type int for addition")
			}
			sum += v
		}
	}
	return sum, nil
}

// Returns the product of all the values in the matrix
// Returns error if non-ints are encountered.
func (m *Matrix) Multiply() (int, error) {
	prod := 1
	for _, row := range m.Data {
		for j := range row {
			v, err := strconv.Atoi(row[j])
			if err != nil {
				return 0, fmt.Errorf("error: non-int values in matrix. all values must be of type int for multiplication")
			}
			prod *= v
		}
	}
	return prod, nil
}

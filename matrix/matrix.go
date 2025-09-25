package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	This files contains the struct/object definition and methods acting on the matrix type
*/

type Matrix struct{
	Data [][]string
	Size int
}

// Echo() returns a string representation of the matrix
func (m *Matrix) Echo() string {
	var matrixStr string
	for _, row := range m.Data {
		// this sprintF looks funny, what is the response obj doing here?
		matrixStr = fmt.Sprintf("%s%s\n", matrixStr, strings.Join(row, ","))
	}
	return matrixStr
}

// Transposes the matrix in memory
func (m *Matrix) Transpose() {
	for row := 0; row < m.Size; row++ {
		for col := row + 1; col < m.Size; col++{
			 m.Data[row][col],  m.Data[col][row] = m.Data[col][row], m.Data[row][col]
		}
	}
}

// Returns a flattened representation of the string.
// Nested slices get reduced to single slice.
// Sliced gets joined into a string
func (m *Matrix) Flatten() string {
	retMatrix := make([]string,0,m.Size)
	for _, v := range m.Data{
		retMatrix = append(retMatrix, v...)
	}
	return strings.Join(retMatrix, ",")
}

func (m *Matrix) Add() (int, error) {
	sum := 0
	for _, row := range m.Data{
		for j := range row{
			v, err := strconv.Atoi(row[j])
			if err != nil{
				return 0, fmt.Errorf("error: non-int values in matrix. All values must be of type int for addition")
			}
			sum += v
		}
	}
	return sum, nil
}

func (m *Matrix) Multiply() (int, error) {
	prod := 1
	for _, row := range m.Data{
		for j := range row{
			v, err := strconv.Atoi(row[j])
			if err != nil{
				return 0, fmt.Errorf("error: non-int values in matrix. All values must be of type int for multiplication")
			}
			prod *= v
		}
	}
	return prod, nil
}

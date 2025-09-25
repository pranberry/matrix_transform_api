package tests

import (
	"league_challenge/matrix"
	"testing"
)

func TestTranspose(t *testing.T) {
	m := &matrix.Matrix{
		Data: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		Size: int(3),
	}

	want := "1,4,7\n2,5,8\n3,6,9\n"

	m.Transpose()
	got := m.Echo()

	if want != got {
		t.Errorf("Failed while transposing matrix.\nWant:\n%s\nGot:\n%s", want, got)
	}

}

func TestFlatten(t *testing.T) {

	m := &matrix.Matrix{
		Data: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		Size: int(3),
	}

	want := "1,2,3,4,5,6,7,8,9"
	got := m.Flatten()
	if want != got {
		t.Errorf("Failed while flattening matrix.\nWant:\n%s\nGot:\n%s", want, got)
	}

}

func TestAddition(t *testing.T) {

	m := &matrix.Matrix{
		Data: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		Size: int(3),
	}

	want := 1+2+3+4+5+6+7+8+9
	got, err := m.Add()
	if err != nil{
		t.Errorf("Failed while adding matrix values with error : %s\n", err.Error())
		return
	}
	if want != got {
		t.Errorf("Failed while adding matrix.\nWant:\n%d\nGot:\n%d", want, got)
	}
}

func TestMultiply(t *testing.T) {

	m := &matrix.Matrix{
		Data: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		Size: int(3),
	}

	want := 1*2*3*4*5*6*7*8*9
	got, err := m.Multiply()
	if err != nil{
		t.Errorf("Failed while multiplying matrix with error : %s\n", err.Error())
		return
	}
	if want != got {
		t.Errorf("Failed while multiplying matrix.\nWant:\n%d\nGot:\n%d", want, got)
		return
	}
}

func TestEcho(t *testing.T) {
	
	m := &matrix.Matrix{
		Data: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		Size: int(3),
	}

	want := "1,2,3\n4,5,6\n7,8,9\n"
	got := m.Echo()

	if want != got {
		t.Errorf("Failed while echoing matrix.\nWant:\n%s\nGot:\n%s", want, got)
	}

}

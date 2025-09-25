package matrix

import (
	"strconv"
	"testing"
)

// matrixFromInts converts an integer grid into the Matrix format used by the
// handlers. Tests rely on this helper to keep the happy-path setup compact
// while still exercising the string parsing logic inside Add/Multiply.
func matrixFromInts(rows [][]int) *Matrix {
	data := make([][]string, len(rows))
	for i, row := range rows {
		data[i] = make([]string, len(row))
		for j, val := range row {
			data[i][j] = strconv.Itoa(val)
		}
	}

	return &Matrix{
		Data: data,
		Size: len(rows),
	}
}

// TestTranspose verifies that Transpose mutates the matrix in-place and that
// the returned view from Echo matches the expected row/column orientation.
func TestTranspose(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		want   string
	}{
		{
			name:   "asymmetric 3x3",
			matrix: matrixFromInts([][]int{{1, 2, 3}, {0, -5, 8}, {9, 4, 11}}),
			want:   "1,0,9\n2,-5,4\n3,8,11\n",
		},
		{
			name:   "already diagonal",
			matrix: matrixFromInts([][]int{{4, 0, 0}, {0, -3, 0}, {0, 0, 7}}),
			want:   "4,0,0\n0,-3,0\n0,0,7\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.matrix.Transpose()
			if got := tt.matrix.Echo(); got != tt.want {
				t.Fatalf("transpose mismatch.\nwant:\n%s\ngot:\n%s", tt.want, got)
			}
		})
	}
}

// TestFlatten ensures we flatten multi-row matrices in row-major order with a
// comma delimiter and without a trailing newline, matching the API contract.
func TestFlatten(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		want   string
	}{
		{
			name:   "single row",
			matrix: matrixFromInts([][]int{{1, 2, 3, 4}}),
			want:   "1,2,3,4",
		},
		{
			name:   "multi row",
			matrix: matrixFromInts([][]int{{-2, 0}, {5, 9}}),
			want:   "-2,0,5,9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Flatten(); got != tt.want {
				t.Fatalf("flatten mismatch.\nwant: %s\ngot: %s", tt.want, got)
			}
		})
	}
}

// TestAdd covers representative positive, zero, and negative values so we know
// the accumulator starts at zero and that sign handling is correct.
func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		want   int
	}{
		{
			name:   "positive values",
			matrix: matrixFromInts([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
			want:   45,
		},
		{
			name:   "with negatives",
			matrix: matrixFromInts([][]int{{-2, 4}, {-3, 6}}),
			want:   5,
		},
		{
			name:   "includes zero",
			matrix: matrixFromInts([][]int{{0, 0}, {0, 0}}),
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.matrix.Add()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("sum mismatch: want %d got %d", tt.want, got)
			}
		})
	}
}

// TestAddError confirms we surface parsing errors rather than silently
// skipping bad data and that the fallback sum remains zero when an error is
// encountered.
func TestAddError(t *testing.T) {
	m := &Matrix{
		Data: [][]string{{"1", "two"}},
		Size: 2,
	}

	got, err := m.Add()
	if err == nil {
		t.Fatalf("expected error for non-integer element; got sum %d", got)
	}
	if got != 0 {
		t.Fatalf("expected zero sum on error, got %d", got)
	}
}

// TestMultiply demonstrates that the product accumulator starts at one and
// that zeros/negatives propagate through the final product as expected.
func TestMultiply(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		want   int
	}{
		{
			name:   "positive values",
			matrix: matrixFromInts([][]int{{1, 2, 3}, {4, 5, 6}}),
			want:   720,
		},
		{
			name:   "contains zero",
			matrix: matrixFromInts([][]int{{7, 0}, {3, 5}}),
			want:   0,
		},
		{
			name:   "includes negative",
			matrix: matrixFromInts([][]int{{-1, 2}, {3, 4}}),
			want:   -24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.matrix.Multiply()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("product mismatch: want %d got %d", tt.want, got)
			}
		})
	}
}

// TestMultiplyError mirrors TestAddError to ensure non-integer elements
// short-circuit multiplication and provide a deterministic zero result.
func TestMultiplyError(t *testing.T) {
	m := &Matrix{
		Data: [][]string{{"foo", "2"}},
		Size: 1,
	}

	got, err := m.Multiply()
	if err == nil {
		t.Fatalf("expected error for non-integer element; got product %d", got)
	}
	if got != 0 {
		t.Fatalf("expected zero product on error, got %d", got)
	}
}

// TestEcho ensures the multi-line representation always ends with a newline
// so handlers can stream it directly without extra formatting logic.
func TestEcho(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		want   string
	}{
		{
			name:   "multi row",
			matrix: matrixFromInts([][]int{{1, 2, 3}, {4, 5, 6}}),
			want:   "1,2,3\n4,5,6\n",
		},
		{
			name:   "single element",
			matrix: matrixFromInts([][]int{{42}}),
			want:   "42\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Echo(); got != tt.want {
				t.Fatalf("echo mismatch.\nwant:\n%s\ngot:\n%s", tt.want, got)
			}
		})
	}
}

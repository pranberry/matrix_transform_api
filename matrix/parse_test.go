package matrix

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// buildMultipartRequest constructs a multipart/form-data request with an optional file payload.
func buildMultipartRequest(includeFile bool, contents string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if includeFile {
		fileWriter, err := writer.CreateFormFile("file", "matrix.csv")
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(fileWriter, strings.NewReader(contents)); err != nil {
			panic(err)
		}
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}

func TestNewMatrix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		includeFile bool
		contents    string
		wantErr     string
		wantSize    int
		wantData    [][]string
	}{
		{
			name:        "valid 2x2 matrix",
			includeFile: true,
			contents:    "1,2\n3,4\n",
			wantSize:    2,
			wantData:    [][]string{{"1", "2"}, {"3", "4"}},
		},
		{
			name: 		"ints with spaces",
			includeFile: true,
			contents: 	" 1,2 \n3 , 4\n",
			wantSize:   2,
			wantData: 	[][]string{{"1", "2"}, {"3", "4"}},
		},
		{
			name:    "empty file",
			includeFile: true,
			contents: "",
			wantErr: "empty matrix",
		},
		{
			name: 		"empty cells",
			includeFile: true,
			contents: 	" 1,,3\n4,5,6\n,8,9\n",
			wantSize:   3,
			wantData: 	[][]string{{"1", "","3"}, {"4", "5","6"},{"", "8","9"}},
		},
		{
			name:        "missing file",
			includeFile: false,
			wantErr:     "must upload form file",
		},
		{
			name:        "non square",
			includeFile: true,
			contents:    "1,2,3\n4,5,6\n",
			wantErr:     "not an NxN matrix",
		},
		{
			name:        "empty file",
			includeFile: true,
			contents:    "",
			wantErr:     "empty matrix",
		},
		{
			name:        "bad csv",
			includeFile: true,
			contents:    "1,\"2\n3,4\n",
			wantErr:     "parse error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := buildMultipartRequest(tc.includeFile, tc.contents)
			m, err := NewMatrix(req)

			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
				}
				if m != nil {
					t.Fatalf("expected nil matrix on error, got %#v", m)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if m.Size != tc.wantSize {
				t.Fatalf("expected size %d, got %d", tc.wantSize, m.Size)
			}

			if !reflect.DeepEqual(tc.wantData, m.Data) {
				t.Fatalf("matrix data mismatch. want %#v, got %#v", tc.wantData, m.Data)
			}
		})
	}
}

func TestValidateNxN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		records [][]string
		wantErr string
	}{
		{
			name:    "valid 3x3",
			records: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
		},
		{
			name:    "non square",
			records: [][]string{{"1", "2"}, {"3"}},
			wantErr: "not an NxN matrix",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := validateNxN(tc.records)
			if tc.wantErr == "" && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
				}
			}
		})
	}
}

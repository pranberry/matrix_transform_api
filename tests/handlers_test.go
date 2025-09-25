package tests

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"league_challenge/handlers"
)

// sampleMatrixCSV provides a consistent 3x3 matrix payload reused across
// success and failure scenarios to keep the assertions focused on handler
// behaviour.
const sampleMatrixCSV = "1,2,3\n4,5,6\n7,8,9\n"

// TestHandlersSuccess ensures every HTTP handler returns a 200 response with
// the expected body when a well-formed matrix file is uploaded.
func TestHandlersSuccess(t *testing.T) {
	for _, tc := range handlerExpectations() {
		t.Run(tc.name, func(t *testing.T) {
			req := newMultipartRequest(t, tc.target, &tc.input)
			rec := httptest.NewRecorder()

			tc.handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("expected status 200, got %d", rec.Code)
			}

			if body := rec.Body.String(); body != tc.wantBody {
				t.Fatalf("unexpected response body: %q", body)
			}
		})
	}
}

// TestHandlersMissingFile verifies that omitting the multipart file triggers a
// client error response describing the missing upload requirement.
func TestHandlersMissingFile(t *testing.T) {
	for _, tc := range handlerExpectations() {
		t.Run(tc.name, func(t *testing.T) {
			req := newMultipartRequest(t, tc.target, nil)
			rec := httptest.NewRecorder()

			tc.handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400, got %d", rec.Code)
			}

			if body := rec.Body.String(); !strings.Contains(body, "must upload form file with key 'file'") {
				t.Fatalf("expected missing file message, got %q", body)
			}
		})
	}
}

// TestHandlersEmptyFile confirms that uploading an empty payload returns a
// clear validation error indicating an empty matrix.
func TestHandlersEmptyFile(t *testing.T) {
	for _, tc := range handlerExpectations() {
		t.Run(tc.name, func(t *testing.T) {
			content := ""
			req := newMultipartRequest(t, tc.target, &content)
			rec := httptest.NewRecorder()

			tc.handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400, got %d", rec.Code)
			}

			if body := rec.Body.String(); !strings.Contains(body, "error: empty matrix") {
				t.Fatalf("expected empty matrix error, got %q", body)
			}
		})
	}
}

// TestHandlersMalformedFile asserts that malformed CSV uploads are rejected
// with a descriptive parsing error across all handlers.
func TestHandlersMalformedFile(t *testing.T) {
	malformed := "1,2,3\n4,5\n7,8,9\n"
	for _, tc := range handlerExpectations() {
		t.Run(tc.name, func(t *testing.T) {
			req := newMultipartRequest(t, tc.target, &malformed)
			rec := httptest.NewRecorder()

			tc.handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400, got %d", rec.Code)
			}

			if body := rec.Body.String(); !strings.Contains(body, "wrong number of fields") {
				t.Fatalf("expected malformed matrix error, got %q", body)
			}
		})
	}
}

type handlerExpectation struct {
	name     string
	target   string
	handler  http.HandlerFunc
	input    string
	wantBody string
}

// handlerExpectations enumerates each handler alongside the endpoint, sample
// matrix payload, and expected response body.
func handlerExpectations() []handlerExpectation {
	return []handlerExpectation{
		{
			name:     "echo",
			target:   "/echo",
			handler:  http.HandlerFunc(handlers.Echo),
			input:    sampleMatrixCSV,
			wantBody: "1,2,3\n4,5,6\n7,8,9\n",
		},
		{
			name:     "transpose",
			target:   "/transpose",
			handler:  http.HandlerFunc(handlers.Transpose),
			input:    sampleMatrixCSV,
			wantBody: "1,4,7\n2,5,8\n3,6,9\n",
		},
		{
			name:     "flatten",
			target:   "/flatten",
			handler:  http.HandlerFunc(handlers.Flatten),
			input:    sampleMatrixCSV,
			wantBody: "1,2,3,4,5,6,7,8,9",
		},
		{
			name:     "addition",
			target:   "/addition",
			handler:  http.HandlerFunc(handlers.Addition),
			input:    sampleMatrixCSV,
			wantBody: "45",
		},
		{
			name:     "multiply",
			target:   "/multiply",
			handler:  http.HandlerFunc(handlers.Multiply),
			input:    sampleMatrixCSV,
			wantBody: "362880",
		},
	}
}

// newMultipartRequest builds a POST request that optionally attaches a matrix
// file, enabling success and failure cases to share a single helper.
func newMultipartRequest(t *testing.T, target string, content *string) *http.Request {
	t.Helper()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if content != nil {
		part, err := writer.CreateFormFile("file", "matrix.csv")
		if err != nil {
			t.Fatalf("failed to create form file: %v", err)
		}

		if _, err := io.Copy(part, strings.NewReader(*content)); err != nil {
			t.Fatalf("failed to write file contents: %v", err)
		}
	}

	if err := writer.Close(); err != nil {
		t.Fatalf("failed to close multipart writer: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, target, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

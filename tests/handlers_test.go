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

// sampleMatrixCSV is reused across tests to exercise successful handler paths.
const sampleMatrixCSV = "1,2,3\n4,5,6\n7,8,9\n"

func TestEchoHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Provide a well-formed CSV file and expect the handler to
		// stream the contents back to the caller unchanged.
		r := newFileUploadRequest(t, "/echo", fileContent(sampleMatrixCSV))
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Echo)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got: %d", rec.Code)
		}

		expected := "1,2,3\n4,5,6\n7,8,9\n"
		if rec.Body.String() != expected {
			t.Errorf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("missing file", func(t *testing.T) {
		// Requests without a file should be rejected with a 400.
		r := newFileUploadRequest(t, "/echo", noFileContent())
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Echo)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got: %d", rec.Code)
		}
	})
}

func TestTransposeHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Expect the handler to transpose the 3x3 matrix correctly.
		r := newFileUploadRequest(t, "/transpose", fileContent(sampleMatrixCSV))
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Transpose)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got: %d", rec.Code)
		}

		expected := "1,4,7\n2,5,8\n3,6,9\n"
		if rec.Body.String() != expected {
			t.Errorf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("missing file", func(t *testing.T) {
		// Requests without a file should be rejected with a 400.
		r := newFileUploadRequest(t, "/transpose", noFileContent())
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Transpose)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got: %d", rec.Code)
		}
	})
}

func TestFlattenHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// A successful flatten request should return a single comma-separated row.
		r := newFileUploadRequest(t, "/flatten", fileContent(sampleMatrixCSV))
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Flatten)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got: %d", rec.Code)
		}

		expected := "1,2,3,4,5,6,7,8,9"
		if rec.Body.String() != expected {
			t.Errorf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("missing file", func(t *testing.T) {
		// Requests without a file should be rejected with a 400.
		r := newFileUploadRequest(t, "/flatten", noFileContent())
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Flatten)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got: %d", rec.Code)
		}
	})
}

func TestAdditionHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Expect the handler to sum all matrix values.
		r := newFileUploadRequest(t, "/addition", fileContent(sampleMatrixCSV))
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Addition)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got: %d", rec.Code)
		}

		expected := "45"
		if rec.Body.String() != expected {
			t.Errorf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("missing file", func(t *testing.T) {
		// Requests without a file should be rejected with a 400.
		r := newFileUploadRequest(t, "/addition", noFileContent())
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Addition)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got: %d", rec.Code)
		}
	})
}

func TestMultiplyHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Expect the handler to multiply all matrix values.
		r := newFileUploadRequest(t, "/multiply", fileContent(sampleMatrixCSV))
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Multiply)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got: %d", rec.Code)
		}

		expected := "362880"
		if rec.Body.String() != expected {
			t.Errorf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("missing file", func(t *testing.T) {
		// Requests without a file should be rejected with a 400.
		r := newFileUploadRequest(t, "/multiply", noFileContent())
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Multiply)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got: %d", rec.Code)
		}
	})
}

func TestHandlersEmptyFile(t *testing.T) {
	for _, tc := range handlerTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			// An empty file is structurally present but should still trigger
			// the handlers' validation error paths.
			r := newFileUploadRequest(t, tc.target, fileContent(""))
			rec := httptest.NewRecorder()

			tc.handler.ServeHTTP(rec, r)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected 400, got: %d", rec.Code)
			}

			if body := rec.Body.String(); !strings.Contains(body, "error: empty matrix") {
				t.Fatalf("expected empty matrix error, got: %q", body)
			}
		})
	}
}

func TestHandlersMalformedFile(t *testing.T) {
	malformed := "1,2,3\n4,5\n7,8,9\n"
	for _, tc := range handlerTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			// Malformed CSV content should be rejected as a bad request.
			r := newFileUploadRequest(t, tc.target, fileContent(malformed))
			rec := httptest.NewRecorder()

			tc.handler.ServeHTTP(rec, r)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected 400, got: %d", rec.Code)
			}

			if body := rec.Body.String(); !strings.Contains(body, "wrong number of fields") {
				t.Fatalf("expected malformed matrix error, got: %q", body)
			}
		})
	}
}

type filePayload struct {
	content *string
}

// fileContent returns a payload that will populate the multipart upload with the
// provided CSV data.
func fileContent(s string) filePayload {
	return filePayload{content: &s}
}

// noFileContent returns a payload that leaves the multipart form without a file part.
func noFileContent() filePayload {
	return filePayload{content: nil}
}

// newFileUploadRequest constructs a multipart/form-data request that optionally
// includes a "file" form part used by the handlers under test.
func newFileUploadRequest(t *testing.T, target string, payload filePayload) *http.Request {
	t.Helper()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if payload.content != nil {
		part, err := writer.CreateFormFile("file", "matrix.csv")
		if err != nil {
			t.Fatalf("failed to create form file: %v", err)
		}
		if _, err := io.Copy(part, bytes.NewBufferString(*payload.content)); err != nil {
			t.Fatalf("failed to write file contents: %v", err)
		}
	}

	if err := writer.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, target, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

type handlerCase struct {
	name    string
	target  string
	handler http.HandlerFunc
}

func handlerTestCases() []handlerCase {
	return []handlerCase{
		{name: "echo", target: "/echo", handler: http.HandlerFunc(handlers.Echo)},
		{name: "transpose", target: "/transpose", handler: http.HandlerFunc(handlers.Transpose)},
		{name: "flatten", target: "/flatten", handler: http.HandlerFunc(handlers.Flatten)},
		{name: "addition", target: "/addition", handler: http.HandlerFunc(handlers.Addition)},
		{name: "multiply", target: "/multiply", handler: http.HandlerFunc(handlers.Multiply)},
	}
}

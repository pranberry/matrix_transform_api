package tests

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"league_challenge/handlers"
)

const sampleMatrixCSV = "1,2,3\n4,5,6\n7,8,9\n"

func TestEchoHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := newFileUploadRequest(t, "/echo", true)
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
		r := newFileUploadRequest(t, "/echo", false)
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
		r := newFileUploadRequest(t, "/transpose", true)
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
		r := newFileUploadRequest(t, "/transpose", false)
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
		r := newFileUploadRequest(t, "/flatten", true)
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
		r := newFileUploadRequest(t, "/flatten", false)
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
		r := newFileUploadRequest(t, "/addition", true)
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
		r := newFileUploadRequest(t, "/addition", false)
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
		r := newFileUploadRequest(t, "/multiply", true)
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
		r := newFileUploadRequest(t, "/multiply", false)
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.Multiply)
		handler.ServeHTTP(rec, r)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected 400, got: %d", rec.Code)
		}
	})
}

func newFileUploadRequest(t *testing.T, target string, includeFile bool) *http.Request {
	t.Helper()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if includeFile {
		part, err := writer.CreateFormFile("file", "matrix.csv")
		if err != nil {
			t.Fatalf("failed to create form file: %v", err)
		}
		if _, err := io.Copy(part, bytes.NewBufferString(sampleMatrixCSV)); err != nil {
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

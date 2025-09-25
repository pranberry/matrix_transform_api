package tests

import (
	"league_challenge/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/echo", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.Echo)
	handler.ServeHTTP(rec, r)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got: %d", rec.Code)
	}
}

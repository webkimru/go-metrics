package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	req, _ := http.NewRequest("POST", "/", nil)
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	w := httptest.NewRecorder()
	h := Middleware(testHandler)
	h.ServeHTTP(w, req)
}

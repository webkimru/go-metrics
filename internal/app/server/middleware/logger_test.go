package middleware

import (
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithLogging(t *testing.T) {
	req, _ := http.NewRequestWithContext(context.Background(), "POST", "/", nil)
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	w := httptest.NewRecorder()
	h := WithLogging(testHandler)
	h.ServeHTTP(w, req)
}

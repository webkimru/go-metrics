package middleware

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	req, _ := http.NewRequestWithContext(context.Background(), "POST", "/", nil)
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	w := httptest.NewRecorder()
	h := TextPlain(testHandler)
	h.ServeHTTP(w, req)

	assert.Equal(t, w.Header().Get("Content-Type"), "text/plain; charset=utf-8")
}

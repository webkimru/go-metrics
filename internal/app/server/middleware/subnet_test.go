package middleware

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrustedSubnet(t *testing.T) {
	a := config.AppConfig{}
	app = &a
	NewMiddleware(app)

	handler := TrustedSubnet(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	srv := httptest.NewServer(handler)
	defer srv.Close()

	t.Run("valid ip", func(t *testing.T) {
		app.TrustedSubnet = "127.0.0.0/8"
		r := httptest.NewRequest("POST", srv.URL, bytes.NewReader([]byte("")))
		r.Header.Set("X-Real-IP", "127.0.0.1")
		r.RequestURI = ""

		resp, err := http.DefaultClient.Do(r)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()
	})

	t.Run("invalid ip", func(t *testing.T) {
		app.TrustedSubnet = "192.168.1.0/32"
		r := httptest.NewRequest("POST", srv.URL, bytes.NewReader([]byte("")))
		r.Header.Set("X-Real-IP", "192.168.0.1")
		r.RequestURI = ""

		resp, err := http.DefaultClient.Do(r)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		defer resp.Body.Close()
	})

	t.Run("without subnet", func(t *testing.T) {
		app.TrustedSubnet = ""
		r := httptest.NewRequest("POST", srv.URL, bytes.NewReader([]byte("")))
		r.RequestURI = ""

		resp, err := http.DefaultClient.Do(r)
		assert.NoError(t, err)
		defer resp.Body.Close()
	})

	t.Run("invalid subnet", func(t *testing.T) {
		app.TrustedSubnet = "none"
		r := httptest.NewRequest("POST", srv.URL, bytes.NewReader([]byte("")))
		r.RequestURI = ""

		resp, err := http.DefaultClient.Do(r)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		defer resp.Body.Close()
	})
}

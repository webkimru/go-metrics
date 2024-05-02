package middleware

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"io"
	"net/http"
)

func Decrypt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.PrivateKeyPEM == nil {
			next.ServeHTTP(w, r)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Log.Errorf("failed ReadAll()=%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := hex.DecodeString(string(b))
		data, err = rsa.DecryptPKCS1v15(rand.Reader, app.PrivateKeyPEM, data)
		if err != nil {
			logger.Log.Errorf("failed ReadAll()=%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.Body = io.NopCloser(bytes.NewReader(data))

		next.ServeHTTP(w, r)
	})
}

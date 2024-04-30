package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewServer(middleware(routes))
	defer ts.Close()

	tests := []struct {
		name               string
		url                string
		method             string
		expectedStatusCode int
	}{
		{"Positive test: counter", "/update/counter/someMetric/123", http.MethodPost, http.StatusOK},
		{"positive test: gauge", "/update/gauge/someMetric/123", http.MethodPost, http.StatusOK},
		{"positive test: gauge", "/update/gauge/someMetric/123.123", http.MethodPost, http.StatusOK},
		{"positive test: default", "/", http.MethodGet, http.StatusOK},
		{"positive test: counter value", "/value/counter/someMetric", http.MethodGet, http.StatusOK},
		{"positive test: gauge value", "/value/gauge/someMetric", http.MethodGet, http.StatusOK},
		{"negative test: post batch", "/updates/", http.MethodPost, http.StatusBadRequest},
		{"negative test: counter", "/update/counter/someMetric/123.123", http.MethodPost, http.StatusBadRequest},
		{"negative test: counter", "/update/counter/someMetric/none", http.MethodPost, http.StatusBadRequest},
		{"negative test: counter", "/update/counter/someMetric/none", http.MethodPost, http.StatusBadRequest},
		{"negative test: gauge", "/update/gauge/someMetric/none", http.MethodPost, http.StatusBadRequest},
		{"negative test: metric name", "/update/counter//123", http.MethodPost, http.StatusNotFound},
		{"nagative test: http method", "/update/counter/someMetric/123", http.MethodGet, http.StatusMethodNotAllowed},
		{"nagative test: wrong url", "/someurl/", http.MethodPost, http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name+":"+tt.url, func(t *testing.T) {
			switch tt.method {
			case http.MethodPost:
				req, err := http.NewRequestWithContext(context.Background(), "POST", ts.URL+tt.url, nil)
				assert.NoError(t, err)
				req.Header.Set("Content-Type", "text/plain")
				client := &http.Client{}
				resp, err := client.Do(req)
				defer resp.Body.Close()
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)

			case http.MethodGet:
				req, err := http.NewRequestWithContext(context.Background(), "GET", ts.URL+tt.url, nil)
				assert.NoError(t, err)
				client := &http.Client{}
				resp, err := client.Do(req)
				defer resp.Body.Close()
				assert.NoError(t, err)
				assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
			}
		})
	}
}

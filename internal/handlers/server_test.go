package handlers

import (
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
				resp, err := ts.Client().Post(ts.URL+tt.url, "text/plain", nil)
				resp.Body.Close()
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)

			case http.MethodGet:
				resp, err := ts.Client().Get(ts.URL + tt.url)
				resp.Body.Close()
				assert.NoError(t, err)
				assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
			}
		})
	}
}

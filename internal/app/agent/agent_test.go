package agent

import (
	"net/http"
	"testing"
)

func TestSendMetric(t *testing.T) {
	tests := []struct {
		name               string
		metric             Metric
		expectedStatusCode int
	}{
		{
			name: "positive test",
			metric: Metric{
				RandomValue: 123.123,
				PollCount:   1,
			},
			expectedStatusCode: http.StatusOK,
		},
		//{"positive test: gauge", "http://localhost:8080/update/gauge/someMetric/123", http.StatusOK},
		//{"positive test: gauge", "http://localhost:8080/update/gauge/someMetric/123.123", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMetric(tt.metric, "http://localhost:8080")
		})
	}
}

package agent

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent/config"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent/metrics"
	"sync"
	"testing"
	"time"
)

func TestGetMetrics(t *testing.T) {
	a := config.AppConfig{
		PollInterval: 1,
	}
	app = a

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	m := metrics.Metric{}
	wg.Add(1)
	go GetMetrics(ctx, &wg, &m)

	time.Sleep(3 * time.Second)
	cancel()

	if m.Alloc == 0 {
		t.Error("Expected Alloc value > 0, but got 0")
	}
}

func TestSend(t *testing.T) {
	a := config.AppConfig{
		SecretKey: "123",
	}
	app = a

	var tests = []struct {
		name   string
		url    string
		metric metrics.RequestMetricSlice
	}{
		{
			name: "positive test",
			url:  "http://localhost:8080/update/",
			metric: metrics.RequestMetricSlice{
				{
					ID:    "someMetric",
					MType: "gauge",
					Value: 123.123,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Send(context.Background(), tt.url, tt.metric)
			assert.Error(t, err)
		})
	}
}

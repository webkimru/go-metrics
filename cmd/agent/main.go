package main

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/agent"
	"time"
)

var m agent.Metric

const (
	pollInterval   = 2
	reportInterval = 10
	targetUrl      = "http://localhost:8080"
)

func main() {
	go agent.GetMetric(&m, pollInterval)

	for {
		time.Sleep(reportInterval * time.Second)
		agent.SendMetric(m, targetUrl)
	}
}

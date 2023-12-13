package main

import (
	"flag"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent"
	"time"
)

var m agent.Metric

var (
	serverAddress  = flag.String("a", "localhost:8080", "server address")
	reportInterval = flag.Duration("r", 10*time.Second, "report interval")
	pollInterval   = flag.Duration("p", 2*time.Second, "poll interval")
)

func main() {

	flag.Parse()

	go agent.GetMetric(&m, *pollInterval)

	for {
		time.Sleep(*reportInterval)
		agent.SendMetric(m, *serverAddress)
	}
}

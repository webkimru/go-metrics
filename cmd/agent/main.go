package main

import (
	"flag"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent"
	"time"
)

var m agent.Metric

var (
	serverAddress  = flag.String("a", "localhost:8080", "server address")
	reportInterval = flag.Int("r", 10, "report interval (in seconds)")
	pollInterval   = flag.Int("p", 2, "poll interval (in seconds)")
)

func main() {

	flag.Parse()

	go agent.GetMetric(&m, *pollInterval)

	for {
		time.Sleep(time.Duration(*reportInterval))
		agent.SendMetric(m, *serverAddress)
	}
}

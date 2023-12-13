package main

import (
	"flag"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent"
	"log"
	"os"
	"strconv"
	"time"
)

var m agent.Metric

var (
	serverAddress  = flag.String("a", "localhost:8080", "server address")
	reportInterval = flag.Int("r", 2, "report interval (in seconds)")
	pollInterval   = flag.Int("p", 1, "poll interval (in seconds)")
)

func main() {

	// разбор командой строки - флаги
	flag.Parse()
	// определение переменных окружения
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		serverAddress = &envRunAddr
	}
	if envReportInterval := os.Getenv("REPORT_INTERVAL"); envReportInterval != "" {
		ri, err := strconv.Atoi(envReportInterval)
		if err != nil {
			log.Fatal(err)
		}
		reportInterval = &ri
	}
	if envPollInterval := os.Getenv("POLL_INTERVAL"); envPollInterval != "" {
		pi, err := strconv.Atoi(envPollInterval)
		if err != nil {
			log.Fatal(err)
		}
		pollInterval = &pi
	}

	go agent.GetMetric(&m, *pollInterval)

	for {
		time.Sleep(time.Duration(*reportInterval))
		agent.SendMetric(m, *serverAddress)
	}
}

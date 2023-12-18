package main

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/agent"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent/metrics"
	"log"
	"time"
)

var m metrics.Metric

func main() {

	// настраиваем/инициализируем приложение
	serverAddress, reportInterval, pollInterval, err := agent.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// получаем метрики
	go agent.GetMetric(&m, *pollInterval)

	// отдаем метрики
	reportDuration := time.Duration(*reportInterval)
	for {
		time.Sleep(reportDuration)
		agent.SendMetric(m, *serverAddress)
	}
}

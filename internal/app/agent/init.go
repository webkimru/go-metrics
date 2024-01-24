package agent

import (
	"flag"
	"github.com/webkimru/go-yandex-metrics/internal/app/agent/logger"
	"log"
	"os"
	"strconv"
)

type ServerAddress *string
type ReportInterval *int
type PollInterval *int

func Setup() (ServerAddress, ReportInterval, PollInterval, error) {

	// задаем флаги для агента
	serverAddress := flag.String("a", "localhost:8080", "server address")
	reportInterval := flag.Int("r", 10, "report interval (in seconds)")
	pollInterval := flag.Int("p", 2, "poll interval (in seconds)")

	// разбор командой строки
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

	// инициализируем логер
	if err := logger.Initialize("info"); err != nil {
		return nil, nil, nil, err
	}

	logger.Log.Infoln(
		"Starting configuration:",
		"ADDRESS", *serverAddress,
		"REPORT_INTERVAL", *reportInterval,
		"POLL_INTERVAL", *pollInterval,
	)

	return serverAddress, reportInterval, pollInterval, nil
}

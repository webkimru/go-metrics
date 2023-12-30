package agent

import (
	"flag"
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

	return serverAddress, reportInterval, pollInterval, nil
}

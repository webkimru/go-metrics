package main

import (
	"flag"
	"github.com/webkimru/go-yandex-metrics/internal/app/server"
	"log"
	"net/http"
	"os"
)

// main начало приложения
func main() {
	// указываем имя флага, значение по умолчанию и описание
	serverAddress := flag.String("a", "localhost:8080", "server address")
	// разбор командной строки
	flag.Parse()
	// определяем переменные окружения
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		serverAddress = &envRunAddr
	}

	if err := server.Setup(); err != nil {
		log.Fatal(err)
	}

	// стартуем сервер
	err := http.ListenAndServe(*serverAddress, server.Middleware(server.Routes()))
	panic(err)
}

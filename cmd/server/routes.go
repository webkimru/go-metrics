package main

import (
	"github.com/webkimru/go-yandex-metrics/internal/handlers"
	"net/http"
)

// routes задаем маршруты для всего сервиса
func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handlers.Repo.Default)
	mux.HandleFunc(`/update/`, handlers.Repo.PostMetrics)

	return mux
}

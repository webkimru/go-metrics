package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/webkimru/go-yandex-metrics/internal/handlers"
	"net/http"
)

// routes задаем маршруты для всего сервиса
func routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", handlers.Repo.Default)
	r.Post("/update/{metric}/{name}/{value}", handlers.Repo.PostMetrics)
	r.Get("/value/{metric}/{name}", handlers.Repo.GetMetric)

	return r
}

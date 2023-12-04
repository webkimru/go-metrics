package main

import (
	"github.com/webkimru/go-yandex-metrics/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, handlers.PostMetrics)

	return mux
}

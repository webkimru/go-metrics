package main

import (
	"github.com/webkimru/go-yandex-metrics/internal/handlers"
	"net/http"
	"strings"
)

func middleware(next http.Handler) http.Handler {
	// получаем Handler приведением типа http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// устанавливаем заголовок для всех ответов
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// Убираем дубли слешей
		r.URL.Path = deduplicate(r.URL.String(), "/")
		next.ServeHTTP(w, r)
	})
}

// Задаем маршруты для всего сервиса
func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handlers.Default)
	mux.HandleFunc(`/update/`, handlers.PostMetrics)

	return mux
}

func deduplicate(str string, cut string) string {
	var newStr strings.Builder
	var old rune
	for _, r := range str {
		switch {
		case r != old, r != int32(cut[0]):
			newStr.WriteRune(r)
			old = r
		}
		continue
	}
	return newStr.String()
}

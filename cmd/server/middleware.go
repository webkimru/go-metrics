package main

import (
	"net/http"
)

// middleware посредник для обработки входящих звпросов
func middleware(next http.Handler) http.Handler {
	// получаем Handler приведением типа http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// устанавливаем заголовок для всех ответов
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		next.ServeHTTP(w, r)
	})
}

package main

import "net/http"

// run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	return http.ListenAndServe(`:8080`, middleware(routes()))
}

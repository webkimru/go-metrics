package main

import "net/http"

// main исходный код программы
func main() {
	if err := run(); err != nil {
		panic(err)
	}

	// стартуем сервер
	err := http.ListenAndServe(`:8080`, middleware(routes()))
	panic(err)
}

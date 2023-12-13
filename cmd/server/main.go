package main

import (
	"flag"
	"net/http"
)

// main исходный код программы
func main() {
	// указываем имя флага, значение по умолчанию и описание
	serverAddress := flag.String("a", "localhost:8080", "server address")
	// разбор командной строки
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}

	// стартуем сервер
	err := http.ListenAndServe(*serverAddress, middleware(routes()))
	panic(err)
}

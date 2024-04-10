package handlers

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories"
	"net/http"
)

func ExampleRepository_Default() {
	type suite struct {
		Store repositories.StoreRepository
		w     http.ResponseWriter
		r     *http.Request
	}

	var example suite

	m := &Repository{
		Store: example.Store,
	}
	m.Default(example.w, example.r)

	// Output HTML:
	//
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	//     <meta charset="UTF-8">
	//     <meta name="viewport" content="width=device-width, initial-scale=1.0">
	//     <title>Metrics</title>
	// </head>
	// <body>
	//
	//
	// </body>
	// </html>
}

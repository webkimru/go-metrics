package main

import "net/http"

func run() error {
	return http.ListenAndServe(`:8080`, middleware(routes()))
}

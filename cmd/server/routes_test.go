package main

import (
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	mux := routes()
	switch v := mux.(type) {
	case *http.ServeMux:
		//
	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}

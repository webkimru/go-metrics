package main

import (
	"net/http"
	"strings"
)

// middleware посредник для обработки входящих звпросов
func middleware(next http.Handler) http.Handler {
	// получаем Handler приведением типа http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// устанавливаем заголовок для всех ответов
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		r.Header.Set("X-Raw-Path", r.URL.Path)
		//// Убираем дубли слешей
		r.URL.Path = deduplicate(r.URL.String(), "/")

		//не нашел варианта:
		//log.Println("URL =", r.URL)
		//log.Println("Path =", r.URL.Path)
		//log.Println("RawPath =", r.URL.RawPath)
		//log.Println("EscapedPath() =", r.URL.EscapedPath())
		//log.Println("String() =", r.URL.String())
		//log.Println("RequestURI() =", r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

// deduplicate убираем дубли слешей из адреса входящего запроса
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

package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/config"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// конфигурация приложения
	a := config.AppConfig{
		StorePriority: config.Memory,
	}
	app = &a

	if err := file.Initialize(app); err != nil {
		log.Fatal(err)
	}

	testStorage := store.NewFakeStorage()
	repo := NewRepo(testStorage)
	NewHandlers(repo, app)

	os.Exit(m.Run())
}

func getRoutes() http.Handler {
	r := chi.NewRouter()
	r.Post("/update/{metric}/{name}/{value}", Repo.PostMetrics)
	r.Get("/value/{metric}/{name}", Repo.GetMetric)
	r.Get("/", Repo.Default)
	r.Post("/updates/", Repo.PostBatchMetrics)

	return r
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		r.Header.Set("X-Raw-Path", r.URL.Path)
		r.URL.Path = deduplicate(r.URL.String(), "/")

		next.ServeHTTP(w, r)
	})
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

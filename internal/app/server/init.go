package server

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/server/handlers"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store"
)

// Setup будет полезна при инициализации зависимостей сервера перед запуском
func Setup() error {
	// задаем вариант хранения
	memStorage := store.NewMemStorage()
	// инициализируем репозиторий хендлеров с указанным вариантом хранения
	repo := handlers.NewRepo(memStorage)
	// инициализвруем хендлеры для работы с репозиторием
	handlers.NewHandlers(repo)

	return nil
}

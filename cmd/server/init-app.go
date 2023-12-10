package main

import (
	"github.com/webkimru/go-yandex-metrics/internal/handlers"
	"github.com/webkimru/go-yandex-metrics/internal/repositories"
	"github.com/webkimru/go-yandex-metrics/internal/repositories/store"
)

// run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	// задаем вариант хранения
	memStorage := store.NewMemStorage()
	// иницилизируем новое хранилище
	storage := repositories.NewStore(memStorage)
	// инициализируем новый репозиторий
	repo := handlers.NewRepo(storage)
	// инициализвруем хендлеры для работы с репозиторием
	handlers.NewHandlers(repo)
	//handlers.NewHandlers(storage)

	return nil
}

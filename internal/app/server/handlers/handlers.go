package handlers

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/server/config"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories"
)

// Repo - репозиторий испльзуется хендлерами
var Repo *Repository
var app *config.AppConfig

// Repository описываем структуру репозитория для хендлеров
type Repository struct {
	Store repositories.StoreRepository
}

// NewRepo создаем новый репозиторий
func NewRepo(repository repositories.StoreRepository) *Repository {
	return &Repository{
		Store: repository,
	}
}

// NewHandlers устанавливаем репозиторий для хендлеров
func NewHandlers(r *Repository, a *config.AppConfig) {
	Repo = r
	app = a
}

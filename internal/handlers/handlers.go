package handlers

import "github.com/webkimru/go-yandex-metrics/repositories"

// Repo - репозиторий испльзуется хендлерами
var Repo *Repository

// Repository описываем структуру репозитория для хендлеров
type Repository struct {
	Store repositories.StoreRepository
}

// NewRepo создаем новый репозиторий
func NewRepo(repository repositories.StoreRepository) *Repository {
	return &Repository{
		Store: repositories.NewStore(repository),
	}
}

// NewHandlers устанавливаем репозиторий для хендлеров
func NewHandlers(r *Repository) {
	Repo = r
}

package repositories

import (
	"context"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/models"
)

// StoreRepository интерфейс хранилища всего сервиса - контракт
// ниже описываем, все, что он должен уметь делать - методы
type StoreRepository interface {
	UpdateCounter(ctx context.Context, name string, value int64) (int64, error)
	UpdateGauge(ctx context.Context, name string, value float64) (float64, error)
	UpdateBatchMetrics(ctx context.Context, metrics []models.Metrics) error
	GetCounter(ctx context.Context, metric string) (int64, error)
	GetGauge(ctx context.Context, metric string) (float64, error)
	GetAllMetrics(ctx context.Context) (map[string]interface{}, error)
}

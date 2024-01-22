package repositories

import "context"

// StoreRepository интерфейс хранилища всего сервиса - контракт
// ниже описываем, все, что он должен уметь делать - методы
type StoreRepository interface {
	UpdateCounter(ctx context.Context, name string, value int64) (int64, error)
	UpdateGauge(ctx context.Context, name string, value float64) (float64, error)
	GetCounter(ctx context.Context, metric string) (int64, error)
	GetGauge(ctx context.Context, metric string) (float64, error)
	GetAllMetrics(ctx context.Context) (map[string]interface{}, error)
}

package repositories

// StoreRepository интерфейс хранилища всего сервиса - контракт
// ниже описываем, все, что он должен уметь делать - методы
type StoreRepository interface {
	UpdateCounter(metric map[string]string) error
	UpdateGauge(metric map[string]string) error
	GetCounter(metric string) (int64, error)
	GetGauge(metric string) (float64, error)
	GetAllMetrics() (map[string]interface{}, error)
}

// Store описываем структура самого хранилища
type Store struct {
	repo StoreRepository
}

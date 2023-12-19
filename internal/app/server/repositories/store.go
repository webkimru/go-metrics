package repositories

// StoreRepository интерфейс хранилища всего сервиса - контракт
// ниже описываем, все, что он должен уметь делать - методы
type StoreRepository interface {
	UpdateCounter(name string, value int64) error
	UpdateGauge(name string, value float64) error
	GetCounter(metric string) (int64, error)
	GetGauge(metric string) (float64, error)
	GetAllMetrics() (map[string]interface{}, error)
}

// Store описываем структура самого хранилища
type Store struct {
	repo StoreRepository
}

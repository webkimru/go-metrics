package repositories

// StoreRepository интерфейс хранилища всего сервиса - контракт
// ниже описываем, все, что он должен уметь делать - методы
type StoreRepository interface {
	UpdateCounter(metric map[string]string) error
	UpdateGauge(metric map[string]string) error
}

// Store описываем структура самого хранили
type Store struct {
	repo StoreRepository
}

// NewStore через конструктор нашего типа реализуем DI
func NewStore(repo StoreRepository) *Store {
	return &Store{
		repo: repo,
	}
}

// UpdateCounter метод обновления метрики Counter
// описываем данный метод, чтобы исполнить контракт интерфейсного типа хранилища
func (s *Store) UpdateCounter(metric map[string]string) error {
	err := s.repo.UpdateCounter(metric)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGauge метод обновлеия метрики Gauge
func (s *Store) UpdateGauge(metric map[string]string) error {
	err := s.repo.UpdateGauge(metric)
	if err != nil {
		return err
	}

	return nil
}

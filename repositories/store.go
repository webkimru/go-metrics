package repositories

// StoreRepository интерфейс хранилища всего сервиса - контракт
// ниже описываем, все, что он должен уметь делать - методы
type StoreRepository interface {
	Update(metric map[string]string) error
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

// Update описываем данный метод, что испольнить контракт интерфейсного типа хранилища
func (s *Store) Update(metric map[string]string) error {
	err := s.repo.Update(metric)
	if err != nil {
		return err
	}

	return nil
}

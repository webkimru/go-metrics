package store

type FakeStorage struct{}

// NewFakeStorage конструктур типа FakeStorage
func NewFakeStorage() *FakeStorage {
	return &FakeStorage{}
}

func (f *FakeStorage) UpdateCounter(_ map[string]string) error {
	return nil
}

func (f *FakeStorage) UpdateGauge(_ map[string]string) error {
	return nil
}

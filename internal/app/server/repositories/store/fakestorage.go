package store

import "context"

type FakeStorage struct{}

// NewFakeStorage конструктур типа FakeStorage
func NewFakeStorage() *FakeStorage {
	return &FakeStorage{}
}

func (f *FakeStorage) UpdateCounter(_ context.Context, _ string, _ int64) (int64, error) {
	return 0, nil
}

func (f *FakeStorage) UpdateGauge(_ context.Context, _ string, _ float64) (float64, error) {
	return 0, nil
}

func (f *FakeStorage) GetCounter(_ context.Context, _ string) (int64, error) {
	return 0, nil
}

func (f *FakeStorage) GetGauge(_ context.Context, _ string) (float64, error) {
	return 0, nil
}

func (f *FakeStorage) GetAllMetrics(_ context.Context) (map[string]interface{}, error) {
	return nil, nil
}

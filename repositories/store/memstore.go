package store

import (
	"errors"
	"github.com/webkimru/go-yandex-metrics/internal/utils"
)

var (
	ErrUpdateFailed = errors.New("update failed")
)

type Counter int64
type Gauge float64

// MemStorage описываем структуру хранилища в памяти
type MemStorage struct {
	Counter map[string]Counter
	Gauge   map[string]Gauge
}

// NewMemStorage конструктур типа MemStorage
func NewMemStorage() *MemStorage {
	return &MemStorage{
		Counter: map[string]Counter{},
		Gauge:   map[string]Gauge{},
	}
}

// Update описываем метод в соответствии с контактном интерфейсного типа StoreRepository
func (ms *MemStorage) Update(metric map[string]string) error {
	switch value := utils.CheckTypeOfMetricValue(metric["value"]).(type) {
	case int64:
		ms.Counter[metric["name"]] += Counter(value)
	case float64:
		ms.Gauge[metric["name"]] = Gauge(value)
	}
	// log.Printf("%#v", ms)
	return ErrUpdateFailed
}

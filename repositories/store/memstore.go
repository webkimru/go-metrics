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
	return &MemStorage{}
}

// Update описываем метод в соответствии с контактном интерфейсного типа StoreRepository
func (ms *MemStorage) Update(metric map[string]string) error {
	switch value := utils.CheckTypeOfMetricValue(metric["value"]).(type) {
	case int64:
		if _, ok := ms.Counter[metric["name"]]; !ok {
			ms.Counter = make(map[string]Counter)
		}
		ms.Counter[metric["name"]] += Counter(value)
	case float64:
		if _, ok := ms.Counter[metric["name"]]; !ok {
			ms.Gauge = make(map[string]Gauge)
		}
		ms.Gauge[metric["name"]] = Gauge(value)
	}
	// log.Printf("%#v", ms)
	return ErrUpdateFailed
}

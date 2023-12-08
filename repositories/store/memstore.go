package store

import (
	"errors"
	"github.com/webkimru/go-yandex-metrics/internal/utils"
	"log"
	"strconv"
	"sync"
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
	mu      sync.Mutex
}

// NewMemStorage конструктур типа MemStorage
func NewMemStorage() *MemStorage {
	return &MemStorage{
		Counter: map[string]Counter{},
		Gauge:   map[string]Gauge{},
	}
}

// UpdateCounter обновляем поле Counter
// описываем метод в соответствии с контактном интерфейсного типа StoreRepository
func (ms *MemStorage) UpdateCounter(metric map[string]string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	value, _ := strconv.ParseInt(metric["value"], 10, 64)
	ms.Counter[metric["name"]] += Counter(value)
	log.Printf("%#v", ms)
	return nil
}

// UpdateGauge обновляем поле UpdateGauge
func (ms *MemStorage) UpdateGauge(metric map[string]string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	switch value := utils.CheckTypeOfMetricValue(metric["value"]).(type) {
	case int64:
	case float64:
		ms.Gauge[metric["name"]] = Gauge(value)
	default:
		return ErrUpdateFailed
	}
	log.Printf("%#v", ms)
	return nil
}

package store

import (
	"errors"
	"fmt"
	"github.com/webkimru/go-yandex-metrics/internal/utils"
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
		Counter: make(map[string]Counter, 1),
		Gauge:   make(map[string]Gauge, 28),
	}
}

// UpdateCounter обновляем поле Counter
// описываем метод в соответствии с контактном интерфейсного типа StoreRepository
func (ms *MemStorage) UpdateCounter(metric map[string]string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	value, _ := strconv.ParseInt(metric["value"], 10, 64)
	ms.Counter[metric["name"]] += Counter(value)
	//log.Printf("%#v", ms)
	return nil
}

// UpdateGauge обновляем поле UpdateGauge
func (ms *MemStorage) UpdateGauge(metric map[string]string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	switch value := utils.CheckTypeOfMetricValue(metric["value"]).(type) {
	case int64:
		ms.Gauge[metric["name"]] = Gauge(value)
	case float64:
		ms.Gauge[metric["name"]] = Gauge(value)
	default:
		return ErrUpdateFailed
	}
	//log.Printf("%#v", ms)
	return nil
}

func (ms *MemStorage) GetCounter(metric string) (int64, error) {
	value, ok := ms.Counter[metric]
	if !ok {
		return 0, errors.New(fmt.Sprintf("%s does not exists", metric))
	}
	return int64(value), nil
}

func (ms *MemStorage) GetGauge(metric string) (float64, error) {
	value, ok := ms.Gauge[metric]
	if !ok {
		return 0, errors.New(fmt.Sprintf("%s does not exists", metric))
	}
	return float64(value), nil
}

func (ms *MemStorage) GetAllMetrics() (map[string]interface{}, error) {
	all := make(map[string]interface{}, 30)
	all["counter"] = ms.Counter
	all["gauge"] = ms.Gauge

	return all, nil
}

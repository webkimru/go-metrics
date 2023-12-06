package store

import (
	"errors"
	"log"
)

var (
	ErrUpdateFailed = errors.New("update failed")
)

// MemStorage описываем структуру хранилища в памяти
type MemStorage struct {
	Counter map[string]int64
	Gauge   map[string]float64
}

// NewMemStorage конструктур типа MemStorage
func NewMemStorage() *MemStorage {
	return &MemStorage{}
}

// Update описываем метод в соответствии с контактном интерфейсного типа StoreRepository
func (ms *MemStorage) Update(metric map[string]string) error {
	log.Println(metric)
	//var i interface{} = metric["value"]
	//switch v := i.(type) {
	//case int64:
	//	ms.Counter[metric["name"]] += v
	//case float64:
	//	ms.Gauge[metric["name"]] = v
	//}

	return ErrUpdateFailed
}

package handlers

import (
	"errors"
	"github.com/webkimru/go-yandex-metrics/internal/utils"
	"net/http"
	"strings"
)

var (
	// ErrURLIsInvalid возвращает ошибку, если ссылка некорректная.
	ErrURLIsInvalid = errors.New("URL is invalid")
)

// Default задет дефолтный маршрут
func (m *Repository) Default(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// PostMetrics обрабатывае входящие метрики
func (m *Repository) PostMetrics(w http.ResponseWriter, r *http.Request) {
	// Принимать метрики по протоколу HTTP методом `POST`
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// При попытке передать запрос без имени метрики возвращать `http.StatusNotFound`.
	// пример: /update/gauge//10
	rawPath := r.Header.Get("X-Raw-Path")
	if rawPath != r.URL.String() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Парсим маршрут и получаем мапку: метрика, значение
	metric, err := parseURL(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// При попытке передать запрос с некорректным типом метрики возвращать `http.StatusBadRequest`.
	if metric["type"] != "counter" && metric["type"] != "gauge" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// При попытке передать запрос с некорректным типом метрики или значением возвращать `http.StatusBadRequest`.
	switch utils.CheckTypeOfMetricValue(metric["value"]).(type) {
	case int64:
		// Запись значения метрики Counter
		err := m.Store.Update(metric)
		if err != nil {
			return
		}
	case float64:
		// Некорректное значение для типа - `http.StatusBadRequest`.
		// Пример: /update/counter/allocCount/20.0003
		if metric["type"] == "counter" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Запись значения метрики gauge
		err := m.Store.Update(metric)
		if err != nil {
			return
		}

	case string:
		// Пример: /update/counter/allocCount/text
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

// parseURL парсит маршруты входящих запросов
func parseURL(r *http.Request) (map[string]string, error) {
	// Реализовать парсинг маршрутов
	// Пример: /update/gauge/speedAverage/200 - len = 5
	// Пример: /update/counter/               - len = 4
	// Пример: /update/                       - len = 3
	// Пример: /update/another/               - len = 4
	metric := map[string]string{}
	slice := strings.Split(r.URL.String(), "/")

	//log.Println("len(slice)=", len(slice))

	// Проверяем корректность маршрута
	if len(slice) == 5 {
		metric["type"] = slice[2]
		metric["name"] = slice[3]
		metric["value"] = slice[4]
		return metric, nil
	}

	return metric, ErrURLIsInvalid
}

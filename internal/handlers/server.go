package handlers

import (
	"log"
	"net/http"
)

func PostMetrics(w http.ResponseWriter, r *http.Request) {
	// 1. Принимать метрики по протоколу HTTP методом `POST`.
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//  2. Парсим маршрут и получаем пару: метрика и значение, либо дефолтные значения
	metric, value := parseMetric(r)
	log.Println(metric, value)

	// 3. Проверяем наличие метрики и значения, и др. пограничные случаи
	switch {
	// Пример: /update/counter/
	// Пример: /update/
	// Пример: /update/another/
	case metric == "" && value == "":
		w.WriteHeader(http.StatusBadRequest)
		return
	// 5. При попытке передать запрос без имени метрики возвращать `http.StatusNotFound`.
	// Пример: /update/counter//123
	case metric == "":
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// 4. Проверяем тип метрики и в случае корректных данных производим запись
	switch v := value.(type) {
	case int:
		// При попытке передать запрос с некорректным типом метрики или значением возвращать `http.StatusBadRequest`.
		// Пример: /update/gauge/speedAverage/200
		if metric == "gauge" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: 5. Запись значение метрики counter
		log.Println(v)
	case float64:
		// При попытке передать запрос с некорректным типом метрики или значением возвращать `http.StatusBadRequest`.
		// Пример: /update/counter/allocCount/20.0003
		if metric == "counter" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: 6. Запись значения метрики gauge
		log.Println(v)
	default:
		// При попытке передать запрос с некорректным типом метрики или значением возвращать `http.StatusBadRequest`.
		// Пример: /update/counter/allocCount/text
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func parseMetric(r *http.Request) (string, interface{}) {
	log.Println("path", r.URL)

	return "", ""
}

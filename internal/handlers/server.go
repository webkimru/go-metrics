package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/webkimru/go-yandex-metrics/internal/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

// Default задет дефолтный маршрут
func (m *Repository) Default(w http.ResponseWriter, _ *http.Request) {
	stringHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Metrics</title>
</head>
<body>
    {{range $k, $v := .counter}}
    	{{$k}} {{$v}}<br>
	{{end}}
    {{range $k, $v := .gauge}}
    	{{$k}} {{$v}}<br>
	{{end}}
</body>
</html>
`

	res, err := m.Store.GetAllMetrics()
	if err != nil {
		return
	}
	log.Println(res)

	t := template.New("Metrics")
	t, err = t.Parse(stringHTML)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "Content-Type")
	w.WriteHeader(http.StatusOK)
	err = t.Execute(w, res)
	if err != nil {
		return
	}
}

// PostMetrics обрабатывае входящие метрики
func (m *Repository) PostMetrics(w http.ResponseWriter, r *http.Request) {
	// Парсим маршрут и получаем мапку: метрика, значение
	metric := make(map[string]string, 3)
	metric["type"] = chi.URLParam(r, "metric")
	metric["name"] = chi.URLParam(r, "name")
	metric["value"] = chi.URLParam(r, "value")

	// При попытке передать запрос с некорректным типом метрики возвращать `http.StatusBadRequest`.
	if metric["type"] != Counter && metric["type"] != Gauge {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// При попытке передать запрос с некорректным типом метрики или значением возвращать `http.StatusBadRequest`.
	switch utils.CheckTypeOfMetricValue(metric["value"]).(type) {
	case int64:
		if metric["type"] == Gauge {
			err := m.Store.UpdateGauge(metric)
			if err != nil {
				return
			}
		}
		if metric["type"] == Counter {
			err := m.Store.UpdateCounter(metric)
			if err != nil {
				return
			}
		}
	case float64:
		// Некорректное значение для типа - `http.StatusBadRequest`.
		// Пример: /update/counter/allocCount/20.0003
		if metric["type"] == Counter {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Запись значения метрики gauge
		err := m.Store.UpdateGauge(metric)
		if err != nil {
			return
		}

	case string:
		// Пример: /update/counter/allocCount/text
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (m *Repository) GetMetric(w http.ResponseWriter, r *http.Request) {
	metric := chi.URLParam(r, "metric")
	name := chi.URLParam(r, "name")

	switch metric {
	case "counter":
		res, err := m.Store.GetCounter(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(strconv.FormatInt(res, 10)))
		if err != nil {
			return
		}
	case "gauge":
		res, err := m.Store.GetGauge(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(strconv.FormatFloat(res, 'f', -1, 64)))
		if err != nil {
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

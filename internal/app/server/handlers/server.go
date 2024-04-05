package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/models"
	"github.com/webkimru/go-yandex-metrics/internal/utils"
	"html/template"
	"net/http"
	"strconv"
)

const (
	Gauge   = "gauge"
	Counter = "counter"

	ContentTypeJSON = "application/json"
)

// Default выдает список всех метрик и их значения в HTML.
func (m *Repository) Default(w http.ResponseWriter, r *http.Request) {
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

	res, err := m.Store.GetAllMetrics(r.Context())
	if err != nil {
		logger.Log.Errorln("failed to get the data from storage, GetAllMetrics() = ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := template.New("Metrics")
	t, err = t.Parse(stringHTML)
	if err != nil {
		logger.Log.Errorln("HTML template is not parsed, Parse() = ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	err = t.Execute(w, res)
	if err != nil {
		logger.Log.Errorln("template execution error, Execute() = ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// PostMetrics обрабатывает входящие метрики.
func (m *Repository) PostMetrics(w http.ResponseWriter, r *http.Request) {
	var metrics models.Metrics
	// application/json
	if r.Header.Get("Content-Type") == ContentTypeJSON {
		if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// text/plain
		metrics.MType = chi.URLParam(r, "metric")
		metrics.ID = chi.URLParam(r, "name")
		switch metrics.MType {
		case Counter:
			value, err := utils.GetInt64ValueFromSting(chi.URLParam(r, "value"))
			if err != nil {
				logger.Log.Errorln(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			metrics.Delta = &value
		case Gauge:
			value, err := utils.GetFloat64ValueFromSting(chi.URLParam(r, "value"))
			if err != nil {
				logger.Log.Errorln(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			metrics.Value = &value
		}
	}

	// При попытке передать запрос с некорректным типом метрики возвращать `http.StatusBadRequest`.
	if metrics.MType != Counter && metrics.MType != Gauge {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// При попытке передать запрос с некорректным значением возвращать `http.StatusBadRequest`.
	switch metrics.MType {
	case Gauge:
		// Обновление данных в хранилище.
		res, err := m.Store.UpdateGauge(r.Context(), metrics.ID, *metrics.Value)
		if err != nil {
			logger.Log.Errorln("failed to update the data from storage, UpdateGauge() = ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		metrics.Value = &res
		// Сохранение данных в файл.
		if err := file.SyncWriter(r.Context(), m.Store.GetAllMetrics); err != nil {
			logger.Log.Errorln("failed to write the data to the file, SyncWriter() =", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Ответ клиенту.
		if err := m.WriteResponseGauge(w, r, metrics); err != nil {
			logger.Log.Errorln("failed to write the data to the connection, WriteResponseGauge() =", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	case Counter:
		// Обновление данных в хранилище.
		res, err := m.Store.UpdateCounter(r.Context(), metrics.ID, *metrics.Delta)
		if err != nil {
			logger.Log.Errorln("failed to update the data from storage, UpdateCounter() = ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		metrics.Delta = &res
		// Сохранение данных в файл.
		if err := file.SyncWriter(r.Context(), m.Store.GetAllMetrics); err != nil {
			logger.Log.Errorln("failed to write the data to the file, SyncWriter() =", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Ответ клиенту.
		if err := m.WriteResponseCounter(w, r, metrics); err != nil {
			logger.Log.Errorln("failed to write the data to the connection, WriteResponseCounter() =", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetMetric выдает данные по запрашиваемой метрике.
func (m *Repository) GetMetric(w http.ResponseWriter, r *http.Request) {
	var metrics models.Metrics
	// application/json
	if r.Header.Get("Content-Type") == ContentTypeJSON {
		if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// text/plain
		metrics.MType = chi.URLParam(r, "metric")
		metrics.ID = chi.URLParam(r, "name")
	}

	switch metrics.MType {
	case Counter:
		res, err := m.Store.GetCounter(r.Context(), metrics.ID)
		if err != nil {
			logger.Log.Infoln("failed to get the data from storage, GetCounter() = ", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		metrics.Delta = &res
		if err := m.WriteResponseCounter(w, r, metrics); err != nil {
			logger.Log.Errorln("failed to write the data to the connection, WriteResponseCounter() =", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case Gauge:
		res, err := m.Store.GetGauge(r.Context(), metrics.ID)
		if err != nil {
			logger.Log.Infoln("failed to get the data from storage, GetGauge() = ", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		metrics.Value = &res
		if err := m.WriteResponseGauge(w, r, metrics); err != nil {
			logger.Log.Errorln("failed to write the data to the connection, WriteResponseGauge() =", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

// WriteResponseCounter отдает кдиенту данные и статус по счетчику Counter.
func (m *Repository) WriteResponseCounter(w http.ResponseWriter, r *http.Request, metrics models.Metrics) error {
	// application/json
	if r.Header.Get("Content-Type") == ContentTypeJSON {
		w.Header().Set("Content-Type", ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(metrics); err != nil {
			return err
		}

		return nil
	}

	// text/plain
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(strconv.Itoa(int(*metrics.Delta))))
	if err != nil {
		return err
	}

	return nil
}

// WriteResponseGauge отдает клиенту данные и статус по счетчику Gauge.
func (m *Repository) WriteResponseGauge(w http.ResponseWriter, r *http.Request, metrics models.Metrics) error {
	// application/json
	if r.Header.Get("Content-Type") == ContentTypeJSON {
		w.Header().Set("Content-Type", ContentTypeJSON)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(metrics); err != nil {
			return err
		}

		return nil
	}

	// text/plain
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(strconv.FormatFloat(*metrics.Value, 'f', -1, 64)))
	if err != nil {
		return err
	}

	return nil
}

// PostBatchMetrics обрабатывает входящие батчи данных с метриками.
func (m *Repository) PostBatchMetrics(w http.ResponseWriter, r *http.Request) {
	var metrics []models.Metrics

	// application/json
	if r.Header.Get("Content-Type") == ContentTypeJSON {
		// Приняли даныне по http в metrics.
		if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Обновляем данные в хранилище.
		err := m.Store.UpdateBatchMetrics(r.Context(), metrics)
		if err != nil {
			logger.Log.Errorln("failed to update the data from storage, UpdateBatchMetrics() = ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Сохранение данных в файл.
		if err = file.SyncWriter(r.Context(), m.Store.GetAllMetrics); err != nil {
			logger.Log.Errorln("failed to write the data to the file, SyncWriter() =", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

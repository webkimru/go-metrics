package utils

import "strconv"

// CheckTypeOfMetricValue определяет тип метрики из маршрута запросов
func CheckTypeOfMetricValue(s string) interface{} {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f
	}
	return ""
}

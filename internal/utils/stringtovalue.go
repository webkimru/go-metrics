package utils

import "strconv"

// GetFloat64ValueFromSting определяет тип метрики из маршрута запросов
func GetFloat64ValueFromSting(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// GetInt64ValueFromSting определяет тип метрики из маршрута запросов
func GetInt64ValueFromSting(s string) (int64, error) {
	f, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

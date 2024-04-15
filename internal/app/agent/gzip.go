package agent

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

// Compress сжимает слайс байт.
func Compress(data *[]byte) error {
	var b bytes.Buffer
	// создаём переменную w — в неё будут записываться входящие данные,
	// которые будут сжиматься и сохраняться в bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return fmt.Errorf("failed init compress writer: %w", err)
	}
	// запись данных
	_, err = w.Write(*data)
	if err != nil {
		return fmt.Errorf("failed write data to compress temporary buffer: %w", err)
	}
	// обязательно нужно вызвать метод Close() — в противном случае часть данных
	// может не записаться в буфер b; если нужно выгрузить все упакованные данные
	// в какой-то момент сжатия, используйте метод Flush()
	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed compress data: %w", err)
	}
	// переменная b содержит сжатые данные
	*data = b.Bytes()

	return nil
}

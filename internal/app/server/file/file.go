package file

import (
	"bufio"
	"encoding/json"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store"
	"os"
)

var Recorder recorder

type StructFile struct {
	Counter map[string]store.Counter
	Gauge   map[string]store.Gauge
}

type recorder struct {
	StoreInterval int
	StoreRestore  bool
	StoreFilePath string
}

func Initialize(storeInterval int, storeFilePath string, storeRestore bool) error {

	Recorder.StoreInterval = storeInterval
	Recorder.StoreFilePath = storeFilePath
	Recorder.StoreRestore = storeRestore

	// пустое значение отключает функцию записи на диск
	if storeFilePath == "" {
		return nil
	}

	// загружать или нет ранее сохранённые значения из указанного файла при старте сервера (по умолчанию true)
	//if storeRestore == true {
	//	//counter, gauge, err := Reader()
	//	//if err != nil {
	//	//	return err
	//	//}
	//	err := handlers.Repo.Store.SetAllMetrics("1") // counter, gauge
	//	if err != nil {
	//		return err
	//	}
	//	//log.Println(counter, gauge)
	//}

	//producer, err := NewProducer(storeFilePath)
	//if err != nil {
	//	return err
	//}
	//Recorder.Producer = producer
	//
	//consumer, err := NewConsumer(storeFilePath)
	//if err != nil {
	//	return err
	//}
	//Recorder.Consumer = consumer

	return nil
}

type Producer struct {
	file *os.File
	// добавляем Writer в Producer
	writer *bufio.Writer
}

func NewProducer(filename string) (*Producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0666)
	if err != nil {
		return nil, err
	}

	return &Producer{
		file: file,
		// создаём новый Writer
		writer: bufio.NewWriter(file),
	}, nil
}

func (p *Producer) WriteJson(metrics map[string]interface{}) error {
	data, err := json.Marshal(&metrics)
	if err != nil {
		return err
	}

	// записываем в буфер
	if _, err := p.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err := p.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	return p.writer.Flush()
}

func (p *Producer) Close() error {
	return p.file.Close()
}

type Consumer struct {
	file *os.File
	// заменяем Reader на Scanner
	scanner *bufio.Scanner
}

func NewConsumer(filename string) (*Consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		file: file,
		// создаём новый scanner
		scanner: bufio.NewScanner(file),
	}, nil
}

func (c *Consumer) ReadJson() (*StructFile, error) {
	// одиночное сканирование до следующей строки
	if !c.scanner.Scan() {
		return nil, c.scanner.Err()
	}
	// читаем данные из scanner
	data := c.scanner.Bytes()

	metrics := StructFile{}
	err := json.Unmarshal(data, &metrics)
	if err != nil {
		return nil, err
	}

	return &metrics, nil
}

func (c *Consumer) Close() error {
	return c.file.Close()
}
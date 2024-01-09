package file

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store"
	"log"
)

func SyncWriter(getAllMetrics func() (map[string]interface{}, error)) error {
	// пустое значение отключает функцию записи на диск
	if Recorder.StoreFilePath == "" {
		return nil
	}

	// значение 0 делает запись синхронной
	if Recorder.StoreInterval > 0 {
		return nil
	}

	res, err := getAllMetrics()
	if err != nil {
		log.Println("failed to get the data from storage, GetAllMetrics() = ", err)
	}

	// записываем в файл
	producer, err := NewProducer(Recorder.StoreFilePath)
	if err != nil {
		return err
	}
	if err := producer.WriteJson(res); err != nil {
		return err
	}
	defer producer.Close()

	return nil
}

func AsyncWriter() {}

func Reader() (map[string]store.Counter, map[string]store.Gauge, error) {
	// читаем из файла
	consumer, err := NewConsumer(Recorder.StoreFilePath)
	if err != nil {
		return nil, nil, err
	}
	res, err := consumer.ReadJson()
	if err != nil {
		return nil, nil, err
	}
	defer consumer.Close()

	return res.Counter, res.Gauge, nil
}

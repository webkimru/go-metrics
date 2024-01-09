package async

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/handlers"
	"log"
	"time"
)

func Writer() {
	// 1. Интервал времени в секундах, по истечении которого текущие показания сервера сохраняются на диск
	// (по умолчанию 300 секунд, значение 0 делает запись синхронной)
	// 2. Синхронная запись реализовано по месту в хендлере сервера в методе PostMetrics()
	if file.Recorder.StoreInterval > 0 {
		go func() {
			storeInterval := time.Duration(file.Recorder.StoreInterval) * time.Second
			for {
				time.Sleep(storeInterval)

				res, err := handlers.Repo.Store.GetAllMetrics()
				if err != nil {
					log.Println("failed to get the data from storage, GetAllMetrics() = ", err)
				}

				// записываем в файл
				producer, err := file.NewProducer(file.Recorder.StoreFilePath)
				if err != nil {
					log.Println(err)
				}
				if err := producer.WriteJson(res); err != nil {
					log.Println(err)
				}
				producer.Close()
			}
		}()
	}
}

package async

import (
	"context"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/handlers"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"time"
)

func FileWriter(ctx context.Context) {
	// 1. Интервал времени в секундах, по истечении которого текущие показания сервера сохраняются на диск
	// (по умолчанию 300 секунд, значение 0 делает запись синхронной)
	// 2. Синхронная запись реализовано по месту в хендлере сервера в методе PostMetrics()
	if file.Recorder.StoreInterval > 0 {
		go func() {
			storeInterval := time.Duration(file.Recorder.StoreInterval) * time.Second
			for {
				time.Sleep(storeInterval)

				SaveData(ctx)
			}
		}()
	}
}

func SaveData(ctx context.Context) {
	res, err := handlers.Repo.Store.GetAllMetrics(ctx)
	if err != nil {
		logger.Log.Errorln("failed to get the data from storage, GetAllMetrics() = ", err)
	}

	// записываем в файл
	producer, err := file.NewProducer(file.Recorder.StoreFilePath)
	if err != nil {
		logger.Log.Errorln(err)
	}
	if err := producer.WriteJSON(res); err != nil {
		logger.Log.Errorln(err)
	}
	producer.Close()
}

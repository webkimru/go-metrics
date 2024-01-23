package async

import (
	"context"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/config"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/handlers"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"time"
)

var app *config.AppConfig

func WriterInitialize(a *config.AppConfig) error {
	app = a

	return nil
}

func FileWriter(ctx context.Context) {
	// Если используется база данных в качестве хранилища, то ничего не делаем
	if app.StorePriority == config.Database {
		return
	}
	// 1. Интервал времени в секундах, по истечении которого текущие показания сервера сохраняются на диск
	// (по умолчанию 300 секунд, значение 0 делает запись синхронной)
	// 2. Синхронная запись реализовано по месту в хендлере сервера в методе PostMetrics()
	if app.FileStore.Interval > 0 {
		go func() {
			storeInterval := time.Duration(app.FileStore.Interval) * time.Second
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
	producer, err := file.NewProducer(app.FileStore.FilePath)
	if err != nil {
		logger.Log.Errorln(err)
	}
	if err := producer.WriteJSON(res); err != nil {
		logger.Log.Errorln(err)
	}
	producer.Close()
}

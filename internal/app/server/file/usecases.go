package file

import (
	"context"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/config"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
)

func SyncWriter(ctx context.Context, getAllMetrics func(ctx context.Context) (map[string]interface{}, error)) error {
	// Если используется база данных в качестве хранилища, то ничего не делаем
	if app.StorePriority == config.Database {
		return nil
	}

	// пустое значение отключает функцию записи на диск
	if app.FileStore.FilePath == "" {
		return nil
	}

	// значение 0 делает запись синхронной
	if app.FileStore.Interval > 0 {
		return nil
	}

	res, err := getAllMetrics(ctx)
	if err != nil {
		logger.Log.Errorln("failed to get the data from storage, GetAllMetrics() = ", err)
	}

	// записываем в файл
	producer, err := NewProducer(app.FileStore.FilePath)
	if err != nil {
		return err
	}
	if err := producer.WriteJSON(res); err != nil {
		return err
	}
	defer producer.Close()

	return nil
}

func Reader() (*StructFile, error) {
	// читаем из файла
	consumer, err := NewConsumer(app.FileStore.FilePath)
	if err != nil {
		return nil, err
	}
	res, err := consumer.ReadJSON()
	if err != nil {
		return nil, err
	}
	defer consumer.Close()

	return res, nil
}

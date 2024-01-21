package server

import (
	"flag"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/handlers"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store/pg"
	"log"
	"os"
	"strconv"
)

// Setup будет полезна при инициализации зависимостей сервера перед запуском
func Setup() (*string, error) {

	// указываем имя флага, значение по умолчанию и описание
	serverAddress := flag.String("a", "localhost:8080", "server address")
	// интервал времени в секундах, по истечении которого текущие показания сервера сохраняются на диск
	// (по умолчанию 300 секунд, значение 0 делает запись синхронной)
	storeInterval := flag.Int("i", 300, "store interval")
	storeFilePath := flag.String("f", "/tmp/metrics-db.json", "file storage path")
	storeRestore := flag.Bool("r", true, "restore saved data")
	// DB
	databaseDSN := flag.String("d", "", "database dsn")
	// разбор командной строки
	flag.Parse()
	// определяем переменные окружения
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		serverAddress = &envRunAddr
	}
	if envStoreInterval := os.Getenv("STORE_INTERVAL"); envStoreInterval != "" {
		si, err := strconv.Atoi(envStoreInterval)
		if err != nil {
			log.Fatal(err)
		}
		storeInterval = &si
	}
	if envStoreFilePath := os.Getenv("FILE_STORAGE_PATH"); envStoreFilePath != "" {
		storeFilePath = &envStoreFilePath
	}
	if envStoreRestore := os.Getenv("RESTORE"); envStoreRestore != "" {
		sr, err := strconv.ParseBool(envStoreRestore)
		if err != nil {
			log.Fatal(err)
		}
		storeRestore = &sr
	}
	if envDatabaseDSN := os.Getenv("DATABASE_DSN"); envDatabaseDSN != "" {
		databaseDSN = &envDatabaseDSN
	}

	// инициализируем логер
	if err := logger.Initialize("info"); err != nil {
		return nil, err
	}

	// инициализируем хранение метрик в файле
	if err := file.Initialize(*storeInterval, *storeFilePath, *storeRestore); err != nil {
		return nil, err
	}

	// инициализируем работу с PostgreSQL
	if err := pg.ConnectToDB(*databaseDSN); err != nil {
		return nil, err
	}

	// задаем вариант хранения
	memStorage := store.NewMemStorage()
	// загружать ранее сохранённые значения из указанного файла при старте сервера
	if *storeRestore {
		res, err := file.Reader()
		if err != nil {
			return nil, err
		}
		// если не пустой файл
		if res != nil {
			memStorage.Counter = res.Counter
			memStorage.Gauge = res.Gauge
		}
	}
	// инициализируем репозиторий хендлеров с указанным вариантом хранения
	repo := handlers.NewRepo(memStorage)
	// инициализвруем хендлеры для работы с репозиторием
	handlers.NewHandlers(repo)

	return serverAddress, nil
}

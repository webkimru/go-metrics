package main

import (
	"context"
	"github.com/webkimru/go-yandex-metrics/internal/app/server"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file/async"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store/pg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// main начало приложения
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// при штатном завершении сервера все накопленные данные должны сохраняться
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		async.SaveData(ctx)
		logger.Log.Infoln("Successful shutdown")
		pg.DB.Conn.Close()
		cancel()
		os.Exit(0)
	}()

	// настраиваем/инициализируем приложение
	serverAddress, err := server.Setup(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// асинхронная запись метрик
	async.FileWriter(ctx)

	// стартуем сервер
	err = http.ListenAndServe(*serverAddress, server.Routes())
	if err != nil {
		panic(err)
	}

	<-ctx.Done()
}

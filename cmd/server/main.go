package main

import (
	"context"
	"fmt"
	"github.com/webkimru/go-yandex-metrics/internal/app/server"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/file/async"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

// main начало приложения
func main() {
	fmt.Println("Build version:", checkVarBuild(buildVersion))
	fmt.Println("Build date:", checkVarBuild(buildDate))
	fmt.Println("Build commit:", checkVarBuild(buildCommit))

	ctx, cancel := context.WithCancel(context.Background())
	// при штатном завершении сервера все накопленные данные должны сохраняться
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// настраиваем/инициализируем приложение
	serverAddress, err := server.Setup(ctx)
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:    *serverAddress,
		Handler: server.Routes(),
	}

	// gracefully shutdown
	go func() {
		<-c
		async.SaveData(ctx)
		logger.Log.Infoln("Successful shutdown")
		server.Shutdown(ctx, srv)
		cancel()
	}()

	// асинхронная запись метрик
	async.FileWriter(ctx)

	// стартуем сервер
	logger.Log.Infof("Starting metric server on %s", *serverAddress)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log.Fatal(err)
	}

	<-ctx.Done()
}

func checkVarBuild(s string) string {
	if s == "" {
		return "N/A"
	}

	return s
}

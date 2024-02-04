package handlers

import (
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/repositories/store/pg"
	"net/http"
)

func (m *Repository) PingPostgreSQL(w http.ResponseWriter, _ *http.Request) {
	err := pg.DB.Conn.Ping()
	if err != nil {
		logger.Log.Errorln("failed to ping PostgreSQL, PingPostgreSQL() = ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

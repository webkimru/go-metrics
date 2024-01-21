package pg

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/webkimru/go-yandex-metrics/internal/app/server/logger"
)

var DB *sql.DB

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDB(dsn string) (err error) {
	DB, err = OpenDB(dsn)
	if err != nil {
		logger.Log.Errorln("Postgres not yet ready...")
		return err
	}

	logger.Log.Infoln("Connected to Postgres")

	return nil
}

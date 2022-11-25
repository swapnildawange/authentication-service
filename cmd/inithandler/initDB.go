package inithandler

import (
	"time"

	"github.com/go-kit/log"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// This function will make a connection to the database only once.
func InitDB(logger log.Logger) *sqlx.DB {
	return connectToDB(logger, 3, 10)
}

func openDB(logger log.Logger, dsn string) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)
	db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		logger.Log("Failed to open Database", err.Error())
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB(logger log.Logger, retries, delay int) *sqlx.DB {
	var (
		connection *sqlx.DB
		err        error
	)
	for i := 0; i < retries; i++ {
		dsn := viper.GetString("DSN")
		connection, err = openDB(logger, dsn)
		if err != nil {
			logger.Log("[debug]", " Postgres not ready yet", "err", err.Error())
			time.Sleep(time.Duration(delay) * time.Second)
			continue
		}
		logger.Log("[debug]", "Successfully connected postgres DB")
		return connection
	}
	return nil
}

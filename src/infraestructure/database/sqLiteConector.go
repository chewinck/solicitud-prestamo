package database

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"

	"github.com/getsentry/sentry-go"
)

func ConnectorSQLite() *sql.DB {
	db, err := sql.Open("sqlite", os.Getenv("DB_PATH"))
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("No se pudo conectar a SQLite: %v", err)
	}

	if err = db.Ping(); err != nil {
		sentry.CaptureException(err)
		log.Fatalf("No se pudo hacer ping a SQLite: %v", err)
	}

	log.Println("Conexión a SQLite exitosa")
	return db
}

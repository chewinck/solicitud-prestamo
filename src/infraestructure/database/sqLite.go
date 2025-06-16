package database

import (
	"log"

	"github.com/getsentry/sentry-go"
	_ "github.com/mattn/go-sqlite3"
)

func IniciarEsquema() error {
	query := `
    CREATE TABLE IF NOT EXISTS solicitud_prestamos (
        uuid TEXT PRIMARY KEY,
		documento_identidad TEXT NOT NULL,
        nombre TEXT NOT NULL,
        monto INTEGER NOT NULL,
		estado TEXT NOT NULL DEFAULT 'Solicitud Inciada',
    );
    `
	db := ConnectorSQLite() // Adjust the path as needed
	_, err := db.Exec(query)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Error al crear la tabla solicitud_prestamos: %v", err)
	}

	return nil
}

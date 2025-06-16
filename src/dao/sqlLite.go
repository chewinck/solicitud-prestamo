package dao

import (
	"solicitudPrestamo/src/infraestructure/database"
	"solicitudPrestamo/src/view/dto"

	"github.com/getsentry/sentry-go"
)

type SqlLite struct {
}

func NewSqlLite() *SqlLite {
	return &SqlLite{}
}

func (s *SqlLite) IniciarSolicitud(solicitudPrestamoDto dto.SolicitudPrestamoDto) error {

	db := database.ConnectorSQLite()

	defer db.Close() // Cierra la conexión después de usarla

	query := `
	INSERT INTO solicitud_prestamos (
		uuid, documento_identidad, nombre, monto, estado
	) VALUES (?, ?, ?, ?, ?);
	`

	_, err := db.Exec(query, solicitudPrestamoDto.UUID, solicitudPrestamoDto.DocumentoIdentidad, solicitudPrestamoDto.NombreCompleto, solicitudPrestamoDto.MontoSolicitado, solicitudPrestamoDto.Estado)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	return nil
}

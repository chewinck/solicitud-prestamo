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

func (s *SqlLite) GuardarScore(uuid string, score int) error {
	db := database.ConnectorSQLite()
	defer db.Close()

	query := `
	UPDATE solicitud_prestamos
	SET score = ?
	WHERE uuid = ?;
	`

	_, err := db.Exec(query, score, uuid)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	return nil
}

func (s *SqlLite) VerificarIdentidad(solicitudPrestamoDto dto.SolicitudPrestamoDto) (dto.SolicitudPrestamoDto, error) {
	db := database.ConnectorSQLite()
	defer db.Close()

	query := `
	SELECT uuid, documento_identidad, nombre, monto, estado, score FROM solicitud_prestamos
	WHERE documento_identidad = ?;
	`

	var documento, nombre, uuid, estado string
	var monto float64
	var score int

	err := db.QueryRow(query, solicitudPrestamoDto.DocumentoIdentidad).Scan(&uuid, &documento, &nombre, &monto, &estado, &score)
	if err != nil {
		sentry.CaptureException(err)
		return dto.SolicitudPrestamoDto{}, err
	}

	solicitudPrestamoDto = dto.SolicitudPrestamoDto{
		UUID:               uuid,
		DocumentoIdentidad: documento,
		NombreCompleto:     nombre,
		MontoSolicitado:    monto,
		Estado:             estado,
		Score:              score,
	}

	if !(documento == solicitudPrestamoDto.DocumentoIdentidad && nombre == solicitudPrestamoDto.NombreCompleto) {
		s.ActualizarEstado(uuid, "Identidad No Verificada Exitosamente")
		return solicitudPrestamoDto, nil
	}

	s.ActualizarEstado(uuid, "Identidad Verificada Exitosamente")
	return solicitudPrestamoDto, nil
}

func (s *SqlLite) ActualizarEstado(uuid string, estado string) error {
	db := database.ConnectorSQLite()
	defer db.Close()

	query := `
	UPDATE solicitud_prestamos
	SET estado = ?
	WHERE uuid = ?;
	`

	_, err := db.Exec(query, estado, uuid)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	return nil
}

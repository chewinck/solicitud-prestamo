package dao

import (
	"database/sql"
	"fmt"
	"solicitudPrestamo/src/infraestructure/database"
	"solicitudPrestamo/src/view/dto"
	"strings"

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

	fmt.Println("solicitudPrestamoDto / inciar solicitud")
	fmt.Println(solicitudPrestamoDto)

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
func (s *SqlLite) GuardarScore(documento string, score int, estado string) error {

	fmt.Println("Documento / guardar score", documento)
	db := database.ConnectorSQLite()
	defer db.Close()

	query := `
	UPDATE solicitud_prestamos
	SET score = ?, estado = ?
	WHERE documento_identidad = ?;
	`

	_, err := db.Exec(query, score, estado, documento)
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
	var score sql.NullInt64

	err := db.QueryRow(query, solicitudPrestamoDto.DocumentoIdentidad).Scan(
		&uuid, &documento, &nombre, &monto, &estado, &score,
	)
	if err != nil {
		sentry.CaptureException(err)
		return dto.SolicitudPrestamoDto{}, fmt.Errorf("no se encontró una solicitud con documento: %s", solicitudPrestamoDto.DocumentoIdentidad)
	}

	// Si score es válido, usamos su valor. Si no, ponemos un default (por ejemplo -1)
	finalScore := -1
	if score.Valid {
		finalScore = int(score.Int64)
	}

	solicitudPrestamoDto = dto.SolicitudPrestamoDto{
		UUID:               uuid,
		DocumentoIdentidad: documento,
		NombreCompleto:     nombre,
		MontoSolicitado:    monto,
		Estado:             estado,
		Score:              finalScore,
	}

	fmt.Println("documento:", documento)
	fmt.Println("nombre:", nombre)
	fmt.Println("solicitudPrestamoDto.DocumentoIdentidad:", solicitudPrestamoDto.DocumentoIdentidad)
	fmt.Println("solicitudPrestamoDto.NombreCompleto:", solicitudPrestamoDto.NombreCompleto)

	docIngresado := strings.TrimSpace(strings.ToLower(solicitudPrestamoDto.DocumentoIdentidad))
	nombreIngresado := strings.TrimSpace(strings.ToLower(solicitudPrestamoDto.NombreCompleto))
	docBD := strings.TrimSpace(strings.ToLower(documento))
	nombreBD := strings.TrimSpace(strings.ToLower(nombre))

	if !(docIngresado == docBD && nombreIngresado == nombreBD) {
		s.ActualizarEstado(uuid, "Identidad No Verificada Exitosamente")
		fmt.Println("entra en: Identidad No Verificada Exitosamente")
		return solicitudPrestamoDto, nil
	}

	s.ActualizarEstado(uuid, "Identidad Verificada Exitosamente")
	fmt.Print("inrgesa Identidad SI Verificada Exitosamente")
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

func (s *SqlLite) ConsultarEstado(uuid string) string {
	db := database.ConnectorSQLite()
	defer db.Close()

	query := `
	SELECT estado FROM solicitud_prestamos
	WHERE uuid= ?;
	`

	var estado string
	err := db.QueryRow(query, uuid).Scan(&estado)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No se encontró ningún estado para el UUID:", uuid)
			return "Estado no encontrado"
		}
		sentry.CaptureException(err)
		return "Error al consultar estado"
	}
	return estado
}

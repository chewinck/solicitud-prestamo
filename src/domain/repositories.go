package domain

import "solicitudPrestamo/src/view/dto"

type SolicitudPrestamoRepository interface {
	IniciarSolicitud(solicitudPrestamoDto dto.SolicitudPrestamoDto) error
	GuardarScore(documento string, score int, estado string) error
	VerificarIdentidad(solicitudPrestamoDto dto.SolicitudPrestamoDto) (dto.SolicitudPrestamoDto, error)
	ActualizarEstado(uuid string, estado string) error
	ConsultarEstado(uuid string) string
}

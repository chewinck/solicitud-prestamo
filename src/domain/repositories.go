package domain

import "solicitudPrestamo/src/view/dto"

type SolicitudPrestamoRepository interface {
	IniciarSolicitud(solicitudPrestamoDto dto.SolicitudPrestamoDto) error
	// ConsultarScore(consultar ConsultarScore) (int, error)
	// EstadoSolicitudPrestamo(estado EstadoSolicitudPrestamo) (EstadoSolicitudPrestamo, error)
	// DecisionFinal(decision DecisionFinal) (DecisionFinal, error)
	// ActualizarEstadoSolicitud(solicitudId string, estado string) error
	// ConsultarEstadoSolicitud(solicitudId string) (EstadoSolicitudPrestamo, error)
}

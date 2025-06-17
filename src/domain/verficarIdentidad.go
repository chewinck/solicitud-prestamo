package domain

import "github.com/getsentry/sentry-go"

type VerificarIdentidad struct {
}

func NewVerificarIdentidad() *VerificarIdentidad {
	return &VerificarIdentidad{}
}

func (v *VerificarIdentidad) CambiarEstado(solicitudPrestamo SolicitudPrestamo) {

	solicitudPrestamoDto, err := solicitudPrestamo.GetRepository().VerificarIdentidad(solicitudPrestamo.GetSolicitudPrestadoDto())
	if err != nil {
		sentry.CaptureException(err)
		return
	}
	solicitudPrestamo.SetsolicitudPrestadoDto(solicitudPrestamoDto)
	solicitudPrestamo.SetEstado(NewDecisionFinal())
	solicitudPrestamo.Avanzar()
}

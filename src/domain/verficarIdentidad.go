package domain

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

type VerificarIdentidad struct {
}

func NewVerificarIdentidad() *VerificarIdentidad {
	return &VerificarIdentidad{}
}

func (v *VerificarIdentidad) CambiarEstado(solicitudPrestamo *SolicitudPrestamo) {

	solicitudPrestamoDto, err := solicitudPrestamo.GetRepository().VerificarIdentidad(solicitudPrestamo.GetSolicitudPrestadoDto())
	if err != nil {
		sentry.CaptureException(err)
		return
	}
	fmt.Println("Ingresa al estado Verificar Identidad ")
	solicitudPrestamo.SetsolicitudPrestadoDto(solicitudPrestamoDto)
	solicitudPrestamo.SetEstado(NewDecisionFinal())
	solicitudPrestamo.Avanzar()
}

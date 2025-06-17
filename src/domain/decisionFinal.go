package domain

import (
	"fmt"
	"strings"
)

type DecisionFinal struct {
}

func NewDecisionFinal() *DecisionFinal {
	return &DecisionFinal{}
}

func (d *DecisionFinal) CambiarEstado(solicitudPrestamo *SolicitudPrestamo) {

	fmt.Print("solicitudPrestamo.GetScore() ", solicitudPrestamo.GetScore())
	if solicitudPrestamo.GetScore() >= 700 && strings.EqualFold(strings.TrimSpace(solicitudPrestamo.solicitudPrestadoDto.Estado), "Identidad Verificada Exitosamente") {
		solicitudPrestamo.GetRepository().ActualizarEstado(solicitudPrestamo.solicitudPrestadoDto.UUID, "Aprobada")
		solicitudPrestamo.solicitudPrestadoDto.Estado = "Aprobada"
		return
	}
	solicitudPrestamo.GetRepository().ActualizarEstado(solicitudPrestamo.solicitudPrestadoDto.UUID, "Rechazada")
	solicitudPrestamo.solicitudPrestadoDto.Estado = "Rechazada"
	solicitudPrestamo.SetEstado(NewDesembolso())
	solicitudPrestamo.Avanzar()
}

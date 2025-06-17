package domain

import "strings"

type DecisionFinal struct {
}

func NewDecisionFinal() *DecisionFinal {
	return &DecisionFinal{}
}

func (d *DecisionFinal) CambiarEstado(solicitudPrestamo SolicitudPrestamo) {
	if !(solicitudPrestamo.GetScore() >= 700 && strings.EqualFold(strings.TrimSpace(solicitudPrestamo.solicitudPrestadoDto.Estado), "Identidad Verificada Exitosamente")) {
		solicitudPrestamo.GetRepository().ActualizarEstado(solicitudPrestamo.solicitudPrestadoDto.UUID, "Rechazada")
		solicitudPrestamo.solicitudPrestadoDto.Estado = "Rechazada"
		return
	}
	solicitudPrestamo.GetRepository().ActualizarEstado(solicitudPrestamo.solicitudPrestadoDto.UUID, "Aprobada")
	solicitudPrestamo.solicitudPrestadoDto.Estado = "Aprobada"
	solicitudPrestamo.SetEstado(NewDesembolso())
	solicitudPrestamo.Avanzar()
}

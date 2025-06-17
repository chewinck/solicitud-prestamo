package domain

import (
	"fmt"
	"solicitudPrestamo/src/infraestructure/score"
)

type ConsultarScore struct {
}

func NewConsultarScore() *ConsultarScore {
	return &ConsultarScore{}
}

func (c *ConsultarScore) CambiarEstado(solicitudPrestamo *SolicitudPrestamo) {
	score := score.ConsultarScore(solicitudPrestamo.GetSolicitudPrestadoDto())
	solicitudPrestamo.GetRepository().GuardarScore(solicitudPrestamo.GetSolicitudPrestadoDto().DocumentoIdentidad, score, "Consultado score")
	solicitudPrestamo.SetScore(score)

	fmt.Println("entra estado scrore", solicitudPrestamo.GetScore())
}

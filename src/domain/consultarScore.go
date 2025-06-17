package domain

import "solicitudPrestamo/src/infraestructure/score"

type ConsultarScore struct {
}

func NewConsultarScore() *ConsultarScore {
	return &ConsultarScore{}
}

func (c *ConsultarScore) CambiarEstado(solicitudPrestamo SolicitudPrestamo) {
	score := score.ConsultarScore(solicitudPrestamo.GetSolicitudPrestadoDto())
	solicitudPrestamo.GetRepository().GuardarScore(solicitudPrestamo.GetSolicitudPrestadoDto().UUID, score)
	solicitudPrestamo.SetScore(score)
}

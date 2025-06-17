package domain

import (
	"solicitudPrestamo/src/view/dto"
)

type Desembolso struct {
}

func NewDesembolso() *Desembolso {
	return &Desembolso{}
}
func (d *Desembolso) CambiarEstado(solicitudPrestamo SolicitudPrestamo) {

	simulacionPrestamoDto := dto.SimulacionPrestamoDto{
		MontoSolicitado: solicitudPrestamo.solicitudPrestadoDto.MontoSolicitado,
		PlazoMeses:      36,
		TasaInteres:     10.5,
		TipoCredito:     "libre Inversión",
	}
	solicitudPrestamo.SetSimulacionPrestamo(simulacionPrestamoDto)
	return

}

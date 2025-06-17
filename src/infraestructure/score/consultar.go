package score

import (
	"solicitudPrestamo/src/infraestructure/util"
	"solicitudPrestamo/src/view/dto"
)

func ConsultarScore(solicitudPrestadoDto dto.SolicitudPrestamoDto) int {
	return util.GenerateRandomScore()

}

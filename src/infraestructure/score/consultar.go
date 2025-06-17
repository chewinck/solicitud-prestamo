package score

import (
	"solicitudPrestamo/src/infraestructure/util"
	"solicitudPrestamo/src/view/dto"
)

func ConsultarScore(solicitudPrestadoDto dto.SolicitudPrestamoDto) int {
	//lógica para consultar el score crediticio en un servicio externo con la infromación de solicitudPrestadoDto
	util.GenerateRandomScore()
	return 700
}

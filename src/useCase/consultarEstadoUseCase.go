package usecase

import (
	"solicitudPrestamo/src/domain"
	"solicitudPrestamo/src/view/dto"
)

type ConsultarEstadoUseCase struct {
	SolicitudPrestamoRepository domain.SolicitudPrestamoRepository
}

func NewConsultarEstadoUseCase(solicitudPrestamoRepository domain.SolicitudPrestamoRepository) *ConsultarEstadoUseCase {
	return &ConsultarEstadoUseCase{
		SolicitudPrestamoRepository: solicitudPrestamoRepository,
	}
}

func (usecase *ConsultarEstadoUseCase) Execute(uuid string) dto.ResponseHttpDto {
	resp := dto.ResponseHttpDto{}

	resp.Code = 200
	resp.Data = usecase.SolicitudPrestamoRepository.ConsultarEstado(uuid)
	return resp

}

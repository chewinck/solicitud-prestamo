package usecase

import (
	"solicitudPrestamo/src/domain"
	"solicitudPrestamo/src/view/dto"
)

type ConsultarScoreUseCase struct {
	SolicitudPrestamoRepository domain.SolicitudPrestamoRepository
}

func NewConsultarScoreUseCase(solicitudPrestamoRepository domain.SolicitudPrestamoRepository) *ConsultarScoreUseCase {
	return &ConsultarScoreUseCase{
		SolicitudPrestamoRepository: solicitudPrestamoRepository,
	}
}

func (useCase *ConsultarScoreUseCase) Execute(documento string) interface{} {
	solcitudPrestamoDto := dto.SolicitudPrestamoDto{
		DocumentoIdentidad: documento,
	}
	solicitudPrestamo := domain.NewSolicitudPrestamo()
	solicitudPrestamo.SetsolicitudPrestadoDto(solcitudPrestamoDto)
	solicitudPrestamo.SetEstado(domain.NewConsultarScore())
	solicitudPrestamo.SetRepository(useCase.SolicitudPrestamoRepository)
	solicitudPrestamo.Avanzar()

	return solicitudPrestamo.GetScore()
}

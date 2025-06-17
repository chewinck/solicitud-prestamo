package usecase

import "solicitudPrestamo/src/domain"

type ConsultarScoreUseCase struct {
	SolicitudPrestamoRepository domain.SolicitudPrestamoRepository
}

func NewConsultarScoreUseCase(solicitudPrestamoRepository domain.SolicitudPrestamoRepository) *ConsultarScoreUseCase {
	return &ConsultarScoreUseCase{
		SolicitudPrestamoRepository: solicitudPrestamoRepository,
	}
}

func (useCase *ConsultarScoreUseCase) Execute(uuid string) interface{} {
	solicitudPrestamo := domain.NewSolicitudPrestamo()
	solicitudPrestamo.SetEstado(domain.NewConsultarScore())
	solicitudPrestamo.SetRepository(useCase.SolicitudPrestamoRepository)
	solicitudPrestamo.Avanzar()

	return solicitudPrestamo.GetScore()
}

package usecase

import (
	"solicitudPrestamo/src/domain"
	"solicitudPrestamo/src/view/dto"
)

type IniciarSolicitudUseCase struct {
	SolicitudPrestamoRepository domain.SolicitudPrestamoRepository
}

func NewIniciarSolicitudUseCase(solicitudPrestamoRepository domain.SolicitudPrestamoRepository) *IniciarSolicitudUseCase {
	return &IniciarSolicitudUseCase{
		SolicitudPrestamoRepository: solicitudPrestamoRepository,
	}
}

func (useCase *IniciarSolicitudUseCase) Execute(solicitudPrestamoDto dto.SolicitudPrestamoDto) interface{} {

	solicitudPrestamo := domain.NewSolicitudPrestamo()
	solicitudPrestamo.SetEstado(domain.NewIniciarSolicitud())
	solicitudPrestamo.SetRepository(useCase.SolicitudPrestamoRepository)
	solicitudPrestamo.Avanzar()

	return solicitudPrestamo.GetScore()
}

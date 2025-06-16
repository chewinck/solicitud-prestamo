package usecase

import (
	"solicitudPrestamo/src/domain"
	"solicitudPrestamo/src/view/dto"

	"github.com/getsentry/sentry-go"
)

type IniciarSolicitudUseCase struct {
	SolicitudPrestamoRepository domain.SolicitudPrestamoRepository
}

func NewIniciarSolicitudUseCase(solicitudPrestamoRepository domain.SolicitudPrestamoRepository) *IniciarSolicitudUseCase {
	return &IniciarSolicitudUseCase{
		SolicitudPrestamoRepository: solicitudPrestamoRepository,
	}
}

func (useCase *IniciarSolicitudUseCase) Execute(solicitudPrestamoDto dto.SolicitudPrestamoDto) error {
	err := useCase.SolicitudPrestamoRepository.IniciarSolicitud(solicitudPrestamoDto)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	return nil
}

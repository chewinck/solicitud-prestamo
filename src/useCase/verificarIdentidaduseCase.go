package usecase

import (
	"fmt"
	"solicitudPrestamo/src/domain"
	"solicitudPrestamo/src/view/dto"
	"strings"
)

type verificarIdentidadUseCase struct {
	SolicitudPrestamoRepository domain.SolicitudPrestamoRepository
}

func NewVerificarIdentidadUseCase(solicitudPrestamoRepository domain.SolicitudPrestamoRepository) *verificarIdentidadUseCase {
	return &verificarIdentidadUseCase{
		SolicitudPrestamoRepository: solicitudPrestamoRepository,
	}
}
func (useCase *verificarIdentidadUseCase) Execute(solicitudPrestamoDto dto.SolicitudPrestamoDto) dto.ResponseHttpDto {
	solicitudPrestamo := domain.NewSolicitudPrestamo()
	solicitudPrestamo.SetsolicitudPrestadoDto(solicitudPrestamoDto)
	solicitudPrestamo.SetEstado(domain.NewVerificarIdentidad())
	solicitudPrestamo.SetRepository(useCase.SolicitudPrestamoRepository)
	solicitudPrestamo.Avanzar()
	

	fmt.Println("Ingresa al caso de uso de verificar")

	if !(strings.EqualFold(strings.TrimSpace(solicitudPrestamo.GetEstado()), "Aprobada")) {
		return dto.ResponseHttpDto{
			Code:    422,
			Message: "Identidad No verificada correctamente",
		}
	}

	return dto.ResponseHttpDto{
		Code:    200,
		Message: "Solicitud aprobada",
		Data:    solicitudPrestamo.GetSimulacionPrestamo(),
	}
}

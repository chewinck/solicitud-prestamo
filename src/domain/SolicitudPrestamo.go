package domain

import (
	"fmt"
	"solicitudPrestamo/src/view/dto"
)

type SolicitudPrestamo struct {
	estadoSolicitudPrestamo     EstadoSolicitudPrestamo
	solicitudPrestamoRepository SolicitudPrestamoRepository
	solicitudPrestadoDto        dto.SolicitudPrestamoDto
}

func NewSolicitudPrestamo() *SolicitudPrestamo {
	return &SolicitudPrestamo{}
}

func (s *SolicitudPrestamo) Avanzar() {
	if s.estadoSolicitudPrestamo == nil {
		fmt.Println("No se puede avanzar, el estado es nil")
		return
	}
	s.estadoSolicitudPrestamo.CambiarEstado(*s)
}

func (s *SolicitudPrestamo) SetEstado(estadoSolicitudPrestamo EstadoSolicitudPrestamo) {
	s.estadoSolicitudPrestamo = estadoSolicitudPrestamo

	fmt.Println("Estado cambiado a:", s.GetEstado())
}

func (s *SolicitudPrestamo) GetEstado() string {
	return s.GetSolicitudPrestadoDto().Estado
}

func (s *SolicitudPrestamo) SetRepository(SolicitudPrestamoRepository SolicitudPrestamoRepository) {
	s.solicitudPrestamoRepository = SolicitudPrestamoRepository
}

func (s *SolicitudPrestamo) GetRepository() SolicitudPrestamoRepository {
	return s.solicitudPrestamoRepository
}

func (s *SolicitudPrestamo) SetsolicitudPrestadoDto(SolicitudPrestamoDto dto.SolicitudPrestamoDto) {
	s.solicitudPrestadoDto = SolicitudPrestamoDto
}

func (s *SolicitudPrestamo) GetSolicitudPrestadoDto() dto.SolicitudPrestamoDto {
	return s.solicitudPrestadoDto
}

func (s *SolicitudPrestamo) SetScore(score int) {
	s.solicitudPrestadoDto.Score = score
}

func (s *SolicitudPrestamo) GetScore() int {
	return s.solicitudPrestadoDto.Score
}

func (s *SolicitudPrestamo) SetSimulacionPrestamo(simulacionPrestamoDto dto.SimulacionPrestamoDto) {
	s.solicitudPrestadoDto.SimulacionPrestamoDto = simulacionPrestamoDto
}

func (s *SolicitudPrestamo) GetSimulacionPrestamo() dto.SimulacionPrestamoDto {
	return s.solicitudPrestadoDto.SimulacionPrestamoDto
}
